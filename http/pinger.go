package httppinger

import (
	"context"
	"fmt"
	"net/http"

	"github.com/amidgo/healthcheck"
)

type Pinger struct {
	client    *http.Client
	pingPath  string
	headerKey string
}

func New(client *http.Client, pingPath, headerKey string) *Pinger {
	return &Pinger{
		client:    client,
		pingPath:  pingPath,
		headerKey: headerKey,
	}
}

func (p *Pinger) Ping(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, p.pingPath, nil)
	if err != nil {
		return fmt.Errorf("make new request with context, %w", err)
	}

	p.addServices(ctx, req.Header)

	resp, err := p.client.Do(req)
	if err != nil {
		return fmt.Errorf("do request, %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return UnexpectedStatusCodeError{StatusCode: resp.StatusCode}
	}

	return nil
}

func (p *Pinger) addServices(ctx context.Context, header http.Header) {
	if p.headerKey == "" {
		return
	}

	for _, service := range healthcheck.ServicesFromContext(ctx) {
		header.Add(p.headerKey, service)
	}
}
