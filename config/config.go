package config

import (
	"context"

	"github.com/cto-kit/config/cloud"
)

type Config interface {
	Config(ctx context.Context, key string, opts ...Option) (string, error)
}

type client interface {
	Get(context.Context, string) (string, error)
}

type config struct {
	client client
}

func New() Config {
	cloudSession := cloud.New()
	return &config{
		client: cloudSession,
	}
}

// optional chainable pattern
func (c *config) Config(ctx context.Context, key string, opts ...Option) (string, error) {
	opt := &option{}
	for _, o := range opts {
		o(opt)
	}
	value, err := c.client.Get(ctx, key)
	if err != nil {
		if value == "" && opt.defaultValue != "" {
			return opt.defaultValue, nil
		}
		return "", ErrNotFound
	}
	return value, nil
}
