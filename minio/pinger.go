package miniopinger

import (
	"context"

	"github.com/minio/minio-go/v7"
)

type Pinger struct {
	client *minio.Client
}

func New(client *minio.Client) *Pinger {
	return &Pinger{
		client: client,
	}
}

func (p *Pinger) Ping(ctx context.Context) error {
	_, err := p.client.ListBuckets(ctx)
	if minio.IsNetworkOrHostDown(err, false) {
		return err
	}

	return nil
}

func (p *Pinger) Name() string {
	return "minio"
}
