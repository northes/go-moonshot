package moonshot

import (
	"net/http"
)

type Client struct {
	cfg             *Config
	requestToURLMap map[string]*Httpx
}

func NewClient(cfg *Config) (*Client, error) {
	if err := cfg.PreCheck(); err != nil {
		return nil, err
	}

	c := &Client{
		cfg:             cfg,
		requestToURLMap: make(map[string]*Httpx),
	}

	if err := c.initAPI(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) initAPI() error {
	err := c.RegisterAPI(new(ChatCompletionsRequest), http.MethodPost)
	if err != nil {
		return err
	}
	err = c.RegisterAPI(new(ListModelsRequest), http.MethodGet)
	if err != nil {
		return err
	}
	err = c.RegisterAPI(new(TokenizersEstimateTokenCountRequest), http.MethodPost)
	if err != nil {
		return err
	}
	return nil
}
