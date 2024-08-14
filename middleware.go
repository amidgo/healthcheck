package healthcheck

import "context"

type Middleware interface{ Pinger(p Pinger) Pinger }

type MiddlewareFunc func(ctx context.Context, p Pinger) error

func (f MiddlewareFunc) Pinger(p Pinger) Pinger {
	var pn Pinger = InlinePinger(
		p.Name(),
		func(ctx context.Context) error {
			return f(ctx, p)
		},
	)

	return pn
}

func Wrap(p Pinger, middlewares ...Middleware) Pinger {
	for i := len(middlewares) - 1; i >= 0; i-- {
		p = middlewares[i].Pinger(p)
	}

	return p
}

func IgnoreServiceMiddleware(serviceName string) Middleware {
	return ignoreServiceMiddleware{
		serviceName: serviceName,
	}
}

type ignoreServiceMiddleware struct {
	serviceName string
}

func (i ignoreServiceMiddleware) Pinger(p Pinger) Pinger {
	var pn Pinger = InlinePinger(
		p.Name(),
		func(ctx context.Context) error {
			services := ServicesFromContext(ctx)
			ctx = contextWithServices(ctx, append(services, i.serviceName))

			return p.Ping(ctx)
		},
	)

	return pn
}
