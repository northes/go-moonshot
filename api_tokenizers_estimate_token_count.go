package moonshot

import (
	"context"
)

type ITokenizers interface {
	EstimateTokenCount(ctx context.Context, req *TokenizersEstimateTokenCountRequest) (resp *TokenizersEstimateTokenCountResponse, err error)
}

type tokenizersEstimateTokenCount struct {
	client *Client
}

func (c *Client) Tokenizers() ITokenizers {
	return &tokenizersEstimateTokenCount{
		client: c,
	}
}

type TokenizersEstimateTokenCountRequest struct {
	Model    ChatCompletionsModelID    `json:"model"`
	Messages []*ChatCompletionsMessage `json:"messages"`
}

type TokenizersEstimateTokenCountResponse struct {
	CommonResponse
	Data *TokenizersEstimateTokenCountResponseData `json:"data"`
}

type TokenizersEstimateTokenCountResponseData struct {
	TotalTokens int `json:"total_tokens"`
}

func (t *tokenizersEstimateTokenCount) EstimateTokenCount(ctx context.Context, req *TokenizersEstimateTokenCountRequest) (*TokenizersEstimateTokenCountResponse, error) {
	const path = "/v1/tokenizers/estimate-token-count"
	estimateTokenCountResp := new(TokenizersEstimateTokenCountResponse)
	resp, err := t.client.HTTPClient().SetPath(path).SetBody(req).Post(ctx)
	if err != nil {
		return nil, err
	}
	if !resp.StatusOK() {
		return nil, ResponseToError(resp)
	}
	err = resp.Unmarshal(estimateTokenCountResp)
	if err != nil {
		return nil, err
	}
	return estimateTokenCountResp, nil
}
