package httpx

import (
	"time"
)

type Option func(client *Client)

func WithTimeout(d time.Duration) Option {
	return func(c *Client) {
		c.timeout = d
	}
}

func WithDebug(b bool) Option {
	return func(c *Client) {
		c.debug = b
	}
}

func WithLogger(logger Logger) Option {
	return func(c *Client) {
		c.logger = logger
	}
}
