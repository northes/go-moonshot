package moonshot

import (
	"context"
	"fmt"

	"github.com/northes/go-moonshot/enum"
)

var (
	_ IRequest = (*TokenizersEstimateTokenCountRequest)(nil)
)

type TokenizersEstimateTokenCount struct{}

type TokenizersEstimateTokenCountRequest struct {
	Model    enum.ChatCompletionsModelID `json:"model"`
	Messages []*ChatCompletionsMessage   `json:"messages"`
}

func (t *TokenizersEstimateTokenCountRequest) Path() string {
	return "/v1/tokenizers/estimate-token-count"
}

type TokenizersEstimateTokenCountResponseData struct {
	TotalTokens int64 `json:"total_tokens"`
}

type TokenizersEstimateTokenCountResponse struct {
	Data *TokenizersEstimateTokenCountResponseData `json:"data"`
}

func (c *Client) TokenizersEstimateTokenCount(ctx context.Context, req *TokenizersEstimateTokenCountRequest) (*TokenizersEstimateTokenCountResponse, error) {
	resp := new(TokenizersEstimateTokenCountResponse)

	_, err := c.Do(ctx, req, resp)
	if err != nil {
		return nil, fmt.Errorf("TokenizersEstimateTokenCount doRequest error: %v", err)
	}

	return resp, nil
}
