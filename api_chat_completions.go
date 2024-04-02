package moonshot

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/northes/gox/httpx"
)

type chat struct {
	client *Client
}

func (c *Client) Chat() *chat {
	return &chat{
		client: c,
	}
}

type ChatCompletionsMessage struct {
	Role    ChatCompletionsMessageRole `json:"role"`
	Content string                     `json:"content"`
}

type ChatCompletionsRequest struct {
	Messages         []*ChatCompletionsMessage `json:"messages"`
	Model            ChatCompletionsModelID    `json:"model"`
	MaxTokens        int64                     `json:"max_tokens"`
	Temperature      float64                   `json:"temperature"`
	TopP             float64                   `json:"top_p"`
	N                int64                     `json:"n"`
	PresencePenalty  float64                   `json:"presence_penalty"`
	FrequencyPenalty float64                   `json:"frequency_penalty"`
	Stop             []string                  `json:"stop"`
	Stream           bool                      `json:"stream"`
}

type ChatCompletionsResponse struct {
	Id      string                            `json:"id"`
	Object  string                            `json:"object"`
	Created int                               `json:"created"`
	Model   string                            `json:"model"`
	Choices []*ChatCompletionsResponseChoices `json:"choices"`
	Usage   *ChatCompletionsResponseUsage     `json:"usage"`
}

type ChatCompletionsResponseChoices struct {
	Index int `json:"index"`

	// return with no stream
	Message *ChatCompletionsMessage `json:"message,omitempty"`
	// return With stream
	Delta *ChatCompletionsMessage `json:"delta,omitempty"`

	FinishReason ChatCompletionsFinishReason `json:"finish_reason"`
}

type ChatCompletionsResponseUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// Completions return the conversation at one time
func (c *chat) Completions(ctx context.Context, req *ChatCompletionsRequest) (*ChatCompletionsResponse, error) {
	const path = "/v1/chat/completions"
	req.Stream = false
	chatCompletionsResp := new(ChatCompletionsResponse)
	resp, err := c.client.HTTPClient().AddPath(path).SetBody(req).Post()
	if err != nil {
		return chatCompletionsResp, err
	}
	if !resp.StatusOK() {
		return nil, StatusCodeToError(resp.Raw().StatusCode)
	}
	err = resp.Unmarshal(chatCompletionsResp)
	if err != nil {
		return nil, err
	}
	return chatCompletionsResp, nil
}

type ChatCompletionsStreamResponse struct {
	resp       *httpx.Response
	isFinished bool
}

type ChatCompletionsStreamResponseReceive struct {
	ChatCompletionsResponse
	isFinished bool
	err        error
}

// CompletionsStream streaming back conversation content
func (c *chat) CompletionsStream(ctx context.Context, req *ChatCompletionsRequest) (*ChatCompletionsStreamResponse, error) {
	const path = "/v1/chat/completions"

	req.Stream = true

	resp, err := c.client.HTTPClient().AddPath(path).SetBody(req).Post()
	if err != nil {
		return nil, err
	}
	if !resp.StatusOK() {
		return nil, StatusCodeToError(resp.Raw().StatusCode)
	}

	streamResp := new(ChatCompletionsStreamResponse)
	streamResp.resp = resp
	return streamResp, nil
}

func (c *ChatCompletionsStreamResponse) Next() <-chan *ChatCompletionsStreamResponseReceive {
	revCh := make(chan *ChatCompletionsStreamResponseReceive, 1)
	reader := bufio.NewReader(c.resp.Raw().Body)

	go func() {
		defer func() {
			close(revCh)
			_ = c.resp.Raw().Body.Close()
		}()
		for {
			line, err := reader.ReadBytes('\n')
			rr := ChatCompletionsStreamResponseReceive{}
			//slog.Debug("next line", string(line))
			if err != nil {
				if err == io.EOF {
					c.sendWithFinish(revCh)
					break
				}
				c.sendWithError(revCh, fmt.Errorf("error reading response body line: %w", err))
				break
			}

			prefix := []byte("data: ")

			if !bytes.HasPrefix(line, prefix) {
				//slog.Debug("no hava prefix,continue", slog.String("line", string(line)))
				continue
			}

			line = bytes.TrimPrefix(bytes.TrimSpace(line), prefix)

			if string(line) == "[DONE]" {
				c.sendWithFinish(revCh)
				break
			}

			err = json.Unmarshal(line, &rr)
			if err != nil {
				c.sendWithError(revCh, fmt.Errorf("error unmarshalling response body line: %w", err))
				break
			}

			c.sendWithMsg(revCh, &rr)
		}
	}()

	return revCh
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
