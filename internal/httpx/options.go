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
