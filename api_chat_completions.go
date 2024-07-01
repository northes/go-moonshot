package moonshot

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/northes/go-moonshot/internal/httpx"
)

type IChat interface {
	Completions(ctx context.Context, req *ChatCompletionsRequest) (*ChatCompletionsResponse, error)
	CompletionsStream(ctx context.Context, req *ChatCompletionsRequest) (*ChatCompletionsStreamResponse, error)
}

type chat struct {
	client *Client
}

// Chat returns a new chat controller
func (c *Client) Chat() IChat {
	return &chat{
		client: c,
	}
}

type ChatCompletionsMessage struct {
	Role    ChatCompletionsMessageRole `json:"role"`
	Content string                     `json:"content"`
	Partial bool                       `json:"partial,omitempty"`
	Name    string                     `json:"name,omitempty"`
	// returns only in use tool response
	ToolCalls []*ChatCompletionsResponseToolCalls `json:"tool_calls,omitempty"`
	// use tool request need it
	ToolCallID string `json:"tool_call_id,omitempty"`
}

type ChatCompletionsRequest struct {
	Messages         []*ChatCompletionsMessage `json:"messages"`
	Model            ChatCompletionsModelID    `json:"model"`
	MaxTokens        int                       `json:"max_tokens"`
	Temperature      float64                   `json:"temperature"`
	TopP             float64                   `json:"top_p"`
	N                int                       `json:"n"`
	PresencePenalty  float64                   `json:"presence_penalty"`
	FrequencyPenalty float64                   `json:"frequency_penalty"`
	Stop             []string                  `json:"stop"`
	Stream           bool                      `json:"stream"`
	// When you use a tool, you need to define it
	Tools []*ChatCompletionsTool `json:"tools,omitempty"`
}

type ChatCompletionsResponse struct {
	ID      string                            `json:"id"`
	Object  string                            `json:"object"`
	Created int                               `json:"created"`
	Model   string                            `json:"model"`
	Choices []*ChatCompletionsResponseChoices `json:"choices"`
	// returns only in non-stream mode
	Usage *ChatCompletionsResponseUsage `json:"usage,omitempty"`
}

type ChatCompletionsResponseChoices struct {
	Index int `json:"index"`

	// return with no stream
	Message *ChatCompletionsMessage `json:"message,omitempty"`
	// return With stream
	Delta *ChatCompletionsMessage `json:"delta,omitempty"`

	FinishReason ChatCompletionsFinishReason `json:"finish_reason"`

	// returns only in stream mode
	Usage *ChatCompletionsResponseUsage `json:"usage,omitempty"`
}

type ChatCompletionsResponseUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// Completions sends a request to the chat completions endpoint
func (c *chat) Completions(ctx context.Context, req *ChatCompletionsRequest) (*ChatCompletionsResponse, error) {
	const path = "/v1/chat/completions"
	req.Stream = false
	chatCompletionsResp := new(ChatCompletionsResponse)
	resp, err := c.client.HTTPClient().SetPath(path).SetBody(req).Post(ctx)
	if err != nil {
		return chatCompletionsResp, err
	}
	if !resp.StatusOK() {
		return nil, ResponseToError(resp)
	}
	err = resp.Unmarshal(chatCompletionsResp)
	if err != nil {
		return nil, err
	}
	return chatCompletionsResp, nil
}

func (c *ChatCompletionsResponse) GetMessage() (*ChatCompletionsMessage, error) {
	if len(c.Choices) == 0 {
		return nil, fmt.Errorf("empty choices")
	}
	for _, choice := range c.Choices {
		if choice.Message != nil {
			return choice.Message, nil
		}
	}

	return nil, fmt.Errorf("no such choice")
}

type ChatCompletionsStreamResponse struct {
	resp *httpx.Response
}

type ChatCompletionsStreamResponseReceive struct {
	ChatCompletionsResponse
	isFinished bool
	err        error
}

// CompletionsStream sends a request to the chat completions endpoint with stream
func (c *chat) CompletionsStream(ctx context.Context, req *ChatCompletionsRequest) (*ChatCompletionsStreamResponse, error) {
	const path = "/v1/chat/completions"

	req.Stream = true

	resp, err := c.client.HTTPClient().SetPath(path).SetBody(req).Post(ctx)
	if err != nil {
		return nil, err
	}
	if !resp.StatusOK() {
		return nil, ResponseToError(resp)
	}

	streamResp := new(ChatCompletionsStreamResponse)
	streamResp.resp = resp
	return streamResp, nil
}

// Receive returns a channel to receive messages from the stream
func (c *ChatCompletionsStreamResponse) Receive() <-chan *ChatCompletionsStreamResponseReceive {
	receiveCh := make(chan *ChatCompletionsStreamResponseReceive, 1)
	reader := bufio.NewReader(c.resp.Raw().Body)

	if c.resp == nil || c.resp.Raw() == nil {
		c.sendWithError(receiveCh, fmt.Errorf("nil response"))
		return receiveCh
	}

	go func() {
		defer func() {
			close(receiveCh)
			_ = c.resp.Raw().Body.Close()
		}()
		for {
			line, err := reader.ReadBytes('\n')
			rr := ChatCompletionsStreamResponseReceive{}
			//slog.Debug("next line", string(line))
			if err != nil {
				if err == io.EOF {
					c.sendWithFinish(receiveCh)
					break
				}
				c.sendWithError(receiveCh, fmt.Errorf("error reading response body line: %w", err))
				break
			}

			prefix := []byte("data: ")

			if !bytes.HasPrefix(line, prefix) {
				//slog.Debug("no hava prefix,continue", slog.String("line", string(line)))
				continue
			}

			line = bytes.TrimPrefix(bytes.TrimSpace(line), prefix)

			if string(line) == "[DONE]" {
				c.sendWithFinish(receiveCh)
				break
			}

			err = json.Unmarshal(line, &rr)
			if err != nil {
				c.sendWithError(receiveCh, fmt.Errorf("error unmarshalling response body line: %w", err))
				break
			}

			c.sendWithMsg(receiveCh, &rr)
		}
	}()

	return receiveCh
}

func (c *ChatCompletionsStreamResponse) sendWithMsg(ch chan<- *ChatCompletionsStreamResponseReceive, msg *ChatCompletionsStreamResponseReceive) {
	ch <- msg
}

func (c *ChatCompletionsStreamResponse) sendWithError(ch chan<- *ChatCompletionsStreamResponseReceive, err error) {
	ch <- &ChatCompletionsStreamResponseReceive{
		err: err,
	}
}

func (c *ChatCompletionsStreamResponse) sendWithFinish(ch chan<- *ChatCompletionsStreamResponseReceive) {
	ch <- &ChatCompletionsStreamResponseReceive{
		isFinished: true,
	}
}

// GetMessage returns the message from the stream
func (c *ChatCompletionsStreamResponseReceive) GetMessage() (*ChatCompletionsMessage, error) {
	if c.err != nil {
		return nil, c.err
	}
	if c.isFinished {
		return nil, io.EOF
	}
	if len(c.Choices) == 0 {
		return nil, fmt.Errorf("empty choices")
	}
	for _, choice := range c.Choices {
		if choice.FinishReason == FinishReasonStop {
			return nil, io.EOF
		}
		if choice.Message != nil {
			return choice.Message, nil
		}
		if choice.Delta != nil {
			return choice.Delta, nil
		}
	}

	return nil, fmt.Errorf("no such choice")
}
