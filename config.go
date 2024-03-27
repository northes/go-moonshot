package moonshot

import (
	"errors"
	"strings"
)

type Config struct {
	Host   string
	APIKey string
}

const DefaultHost = "https://api.moonshot.cn"

var ConfigDefault = &Config{
	Host: DefaultHost,
}

func newConfigDefault() *Config {
	return ConfigDefault
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
