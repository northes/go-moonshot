package moonshot

import (
	"github.com/northes/go-moonshot/internal/httpx"
	"github.com/northes/go-moonshot/internal/httpx/tools"
)

type Client struct {
	cfg *Config
}

// NewClient creates a new client
func NewClient(key string) (*Client, error) {
	cfg := NewConfig(
		WithAPIKey(key),
	)

	if err := cfg.preCheck(); err != nil {
		return nil, err
	}

	return NewClientWithConfig(cfg)
}

// NewClientWithConfig creates a new client with a custom configuration
func NewClientWithConfig(cfg *Config) (*Client, error) {
	if cfg == nil {
		cfg = newConfigDefault()
	}

	if err := cfg.preCheck(); err != nil {
		return nil, err
	}

	c := &Client{
		cfg: cfg,
	}

	return c, nil
}

// HTTPClient returns a new http client
func (c *Client) HTTPClient() *httpx.Client {
	return httpx.NewClient(c.cfg.Host).AddHeader(tools.AuthorizationHeaderKey, tools.ToBearToken(c.cfg.APIKey))
}
