package healthcheck

import (
	"context"
	"fmt"
	"strings"

	"golang.org/x/sync/errgroup"
)

//go:generate mockgen -source pinger.go -destination mocks/pinger.go -package healthcheckmocks

type Pinger interface {
	Ping(ctx context.Context) error
	Name() string
}

type inlinePinger struct {
	pf   func(ctx context.Context) error
	name string
}

func InlinePinger(name string, pf func(ctx context.Context) error) Pinger {
	return inlinePinger{
		name: name,
		pf:   pf,
	}
}

func (i inlinePinger) Ping(ctx context.Context) error {
	return i.pf(ctx)
}

func (i inlinePinger) Name() string {
	return i.name
}

func Join(pingers ...Pinger) Pinger {
	switch len(pingers) {
	case 0:
		return EmptyPinger{}
	case 1:
		return pingers[0]
	default:
		return &joinPinger{
			pingers: pingers,
		}
	}
}

type EmptyPinger struct{}

func (EmptyPinger) Ping(ctx context.Context) error {
	return nil
}

func (EmptyPinger) Name() string {
	return "empty"
}

type joinPinger struct {
	pingers []Pinger
}

func (j joinPinger) Ping(ctx context.Context) error {
	errgr, ctx := errgroup.WithContext(ctx)

	for _, pinger := range j.pingers {
		errgr.Go(ping(ctx, pinger))
	}

	return errgr.Wait()
}

func ping(ctx context.Context, pinger Pinger) func() error {
	return func() error {
		return pinger.Ping(ctx)
	}
}

func (j *joinPinger) Name() string {
	names := make([]string, 0, len(j.pingers))

	for _, pinger := range j.pingers {
		names = append(names, pinger.Name())
	}

	return fmt.Sprintf("join of [%s]", strings.Join(names, ","))
}
