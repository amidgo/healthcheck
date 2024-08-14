package healthcheck_test

import (
	"context"
	"testing"

	"github.com/amidgo/healthcheck"
	"github.com/stretchr/testify/require"
)

func Test_Wrap(t *testing.T) {
	ctx := context.Background()

	counter := 0

	var first healthcheck.MiddlewareFunc = func(ctx context.Context, p healthcheck.Pinger) error {
		require.Equal(t, counter, 0)
		counter++

		return p.Ping(ctx)
	}

	var second healthcheck.MiddlewareFunc = func(ctx context.Context, p healthcheck.Pinger) error {
		require.Equal(t, counter, 1)
		counter++

		return p.Ping(ctx)
	}

	var third healthcheck.MiddlewareFunc = func(ctx context.Context, p healthcheck.Pinger) error {
		require.Equal(t, counter, 2)
		counter++

		return p.Ping(ctx)
	}

	pinger := healthcheck.Wrap(healthcheck.EmptyPinger{}, first, second, third)

	err := pinger.Ping(ctx)
	require.NoError(t, err)

	require.Equal(t, counter, 3)
}
