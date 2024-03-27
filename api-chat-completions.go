package moonshot

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/northes/go-moonshot/enum"
)

var (
	_ IRequest = (*ChatCompletionsRequest)(nil)
)

type ChatCompletions struct{}

type ChatCompletionsRequest struct {
	Messages         []*ChatCompletionsMessage `json:"messages"`
	Model            enum.ChatCompletionsModelID
	MaxTokens        int64
	Temperature      float64
	TopP             float64
	N                int64
	PresencePenalty  float64
	FrequencyPenalty float64
	Stop             []string
	Stream           bool
}

type ChatCompletionsMessage struct {
	Role    enum.ChatCompletionsMessageRole `json:"role"`
	Content string                          `json:"content"`
}

type ChatCompletionsResponseChoices struct {
	Index int `json:"index"`

	Message *ChatCompletionsMessage `json:"message,omitempty"`
	Delta   *ChatCompletionsMessage `json:"delta,omitempty"`

	FinishReason enum.ChatCompletionsFinishReason `json:"finish_reason"`
}

type ChatCompletionsResponseUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type ChatCompletionsResponse struct {
	Id      string                            `json:"id"`
	Object  string                            `json:"object"`
	Created int                               `json:"created"`
	Model   string                            `json:"model"`
	Choices []*ChatCompletionsResponseChoices `json:"choices"`
	Usage   *ChatCompletionsResponseUsage     `json:"usage"`
}

func (c *ChatCompletionsRequest) Path() string {
	return "/v1/chat/completions"
}

func (c *Client) ChatCompletions(ctx context.Context, req *ChatCompletionsRequest) (*ChatCompletionsResponse, error) {
	req.Stream = false

	resp, err := c.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	if err = StatusCodeToError(resp.StatusCode); err != nil {
		return nil, fmt.Errorf("bad response from moonshot: %d", resp.StatusCode)
	}

	var chatResp ChatCompletionsResponse
	err = json.Unmarshal(body, &chatResp)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response body: %w", err)
	}
	return &chatResp, nil
}

func (c *Client) ChatCompletionsStream(ctx context.Context, req *ChatCompletionsRequest, respCh chan<- *ChatCompletionsResponse, done chan<- struct{}) error {

	if respCh == nil || done == nil {
		return errors.New("chat completions streaming requests must have a non-nil channel")
	}

	req.Stream = true

	resp, err := c.Do(ctx, req)
	if err != nil {
		return fmt.Errorf("error do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad response from moonshot: %d", resp.StatusCode)
	}

	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadBytes('\n')
		//fmt.Printf("next line: %v\n", string(line))
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("error reading response body line: %w", err)
		}

		prefix := []byte("data: ")

		if !bytes.HasPrefix(line, prefix) {
			//fmt.Println("no hava prefix,continue")
			continue
		}

		line = bytes.TrimPrefix(bytes.TrimSpace(line), prefix)

		if string(line) == "[DONE]" {
			break
		}

		//fmt.Printf("trim prefix line: %v %v %v\n", string(line), "[DONE]", string(line) == "[DONE]")

		rr := ChatCompletionsResponse{}
		err = json.Unmarshal(line, &rr)
		if err != nil {
			return fmt.Errorf("error unmarshalling response body line: %w", err)
		}
		respCh <- &rr
	}

	done <- struct{}{}

	return nil
}

func (i *ChatCompletionsResponseChoices) IsFinishStop() bool {
	return i.FinishReason == enum.FinishReasonStop
}

func (i *ChatCompletionsResponseChoices) IsFinishLength() bool {
	return i.FinishReason == enum.FinishReasonLength
}

func (c *ChatCompletionsResponse) CanGetContent() bool {
	for _, choice := range c.Choices {
		if choice.FinishReason == enum.FinishReasonStop {
			return false
		}
		if choice.Message != nil {
			if len(choice.Message.Content) == 0 {
				return false
			}
		}
		if choice.Delta != nil {
			if len(choice.Delta.Content) == 0 {
				return false
			}
		}
	}
	return true
}
