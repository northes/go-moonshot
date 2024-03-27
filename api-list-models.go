package moonshot

import (
	"context"
)

var (
	_ IRequest = (*ListModelsRequest)(nil)
)

type ListModelsRequest struct {
}

type ListModelResponseData struct {
	Created    int                             `json:"created"`
	Id         string                          `json:"id"`
	Object     string                          `json:"object"`
	OwnedBy    string                          `json:"owned_by"`
	Permission []*ListModelsResponsePermission `json:"permission"`
	Root       string                          `json:"root"`
	Parent     string                          `json:"parent"`
}

type ListModelsResponsePermission struct {
	Created            int    `json:"created"`
	Id                 string `json:"id"`
	Object             string `json:"object"`
	AllowCreateEngine  bool   `json:"allow_create_engine"`
	AllowSampling      bool   `json:"allow_sampling"`
	AllowLogprobs      bool   `json:"allow_logprobs"`
	AllowSearchIndices bool   `json:"allow_search_indices"`
	AllowView          bool   `json:"allow_view"`
	AllowFineTuning    bool   `json:"allow_fine_tuning"`
	Organization       string `json:"organization"`
	Group              string `json:"group"`
	IsBlocking         bool   `json:"is_blocking"`
}

type ListModelsResponse struct {
	Object string                   `json:"object"`
	Data   []*ListModelResponseData `json:"data"`
}

type APIListModel struct{}

func (l *ListModelsRequest) Path() string {
	return "/v1/models"
}

func (c *Client) ListModels(ctx context.Context) (*ListModelsResponse, error) {
	resp := new(ListModelsResponse)

	_, err := c.Do(ctx, new(ListModelsRequest), resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
