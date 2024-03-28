package moonshot

import (
	"context"

	"github.com/northes/go-moonshot/enum"
	"github.com/northes/gox/httpx"
)

type tokenizersEstimateTokenCount struct {
	client *httpx.Client
}

func (c *Client) Tokenizers() *tokenizersEstimateTokenCount {
	return &tokenizersEstimateTokenCount{
		client: c.newHTTPClient(),
	}
}

type TokenizersEstimateTokenCountRequest struct {
	Model    enum.ChatCompletionsModelID `json:"model"`
	Messages []*ChatCompletionsMessage   `json:"messages"`
}

type TokenizersEstimateTokenCountResponse struct {
	CommonResponse
	Data *TokenizersEstimateTokenCountResponseData `json:"data"`
}

type TokenizersEstimateTokenCountResponseData struct {
	TotalTokens int64 `json:"total_tokens"`
}

func (t *tokenizersEstimateTokenCount) EstimateTokenCount(ctx context.Context, req *TokenizersEstimateTokenCountRequest) (*TokenizersEstimateTokenCountResponse, error) {
	const path = "/v1/tokenizers/estimate-token-count"
	estimateTokenCountResp := new(TokenizersEstimateTokenCountResponse)
	resp, err := t.client.AddPath(path).SetBody(req).Post()
	if err != nil {
		return nil, err
	}
	if !resp.StatusOK() {
		return nil, StatusCodeToError(resp.Raw().StatusCode)
	}
	err = resp.Unmarshal(estimateTokenCountResp)
	if err != nil {
		return nil, err
	}
	return estimateTokenCountResp, nil
}
