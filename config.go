package moonshot

import (
	"errors"
	"strings"
)

type Config struct {
	Host   string
	APIKey string
	Debug  bool
}

const DefaultHost = "https://api.moonshot.cn"

func newConfigDefault() *Config {
	return &Config{
		Host: DefaultHost,
	}
}

func NewConfig(opts ...Option) *Config {
	cfg := newConfigDefault()
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

func (c *Config) PreCheck() error {
	if len(c.APIKey) == 0 {
		return errors.New("API key is required")
	}
	return nil
}

type Option func(*Config)

func SetHost(host string) Option {
	return func(c *Config) {
		c.Host = strings.TrimSuffix(host, "/")
	}
}

func SetAPIKey(key string) Option {
	return func(c *Config) {
		c.APIKey = key
	}
}

func SetDebugMod() Option {
	return func(c *Config) {
		c.Debug = true
	}
}
