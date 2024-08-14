package redispinger

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Pinger struct {
	client *redis.Client
}

func New(client *redis.Client) *Pinger {
	return &Pinger{
		client: client,
	}
}

func (p *Pinger) Ping(ctx context.Context) error {
	return p.client.Ping(ctx).Err()
}

func (p *Pinger) Name() string {
	return "redis"
}
