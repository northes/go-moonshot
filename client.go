package moonshot

import (
	"github.com/northes/gox/httpx"
	"github.com/northes/gox/httpx/httpxutils"
)

type Client struct {
	cfg *Config
}

func NewClient(cfg *Config) (*Client, error) {
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
	).AddHeader(httpxutils.AuthorizationHeaderKey, httpxutils.ToBearToken(c.cfg.APIKey))
}
