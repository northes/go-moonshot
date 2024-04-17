package moonshot

import (
	"errors"

	"github.com/northes/go-moonshot/internal/httpx"
	"github.com/northes/go-moonshot/internal/httpx/tools"
)

type Client struct {
	cfg *Config
}

func NewClient(key string) (*Client, error) {
	if len(key) == 0 {
		return nil, errors.New("key is required")
	}

	cfg := NewConfig(
		WithAPIKey(key),
	)

	return NewClientWithConfig(cfg)
}

func NewClientWithConfig(cfg *Config) (*Client, error) {
	if cfg == nil {
		cfg = newConfigDefault()
	}
	if err := cfg.PreCheck(); err != nil {
		return nil, err
	}
	c := &Client{
		cfg: cfg,
	}
	return c, nil
}

func (c *Client) HTTPClient() *httpx.Client {
	return httpx.NewClient(c.cfg.Host,
		httpx.WithDebug(c.cfg.Debug),
	).AddHeader(tools.AuthorizationHeaderKey, tools.ToBearToken(c.cfg.APIKey))
}
