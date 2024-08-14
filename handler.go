package healthcheck

import (
	"context"
	"slices"
)

type servicesContextKey struct{}

func contextWithServices(ctx context.Context, services []string) context.Context {
	return context.WithValue(ctx, servicesContextKey{}, services)
}

func ServicesFromContext(ctx context.Context) []string {
	services, ok := ctx.Value(servicesContextKey{}).([]string)
	if !ok {
		return nil
	}

	return services
}

//go:generate mockgen -source handler.go -destination mocks/handler.go -package healthcheckmocks

type Handler interface {
	Handle(ctx context.Context, services ...string) error
}

type handler struct {
	name   string
	pinger Pinger
}

func NewHandler(name string, pinger Pinger) Handler {
	return &handler{
		name:   name,
		pinger: pinger,
	}
}

func (h *handler) Handle(ctx context.Context, services ...string) error {
	if slices.Contains(services, h.name) {
		return nil
	}

	services = append(services, h.name)

	ctx = contextWithServices(ctx, services)

	return h.pinger.Ping(ctx)
}
