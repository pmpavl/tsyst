//nolint:stylecheck,revive
package contextOS

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/pmpavl/tsyst/pkg/log"
)

func Background() context.Context {
	return wrapContext(context.Background())
}

func WrapContext(ctx context.Context) context.Context {
	return wrapContext(ctx)
}

func wrapContext(ctx context.Context) context.Context {
	logger := log.For("contextOS")
	signals := []os.Signal{syscall.SIGINT, syscall.SIGTERM}

	ctx, cancel := context.WithCancel(ctx)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, signals...)

	go func() {
		select {
		case <-ctx.Done():
		case sig := <-sigs:
			logger.Info().Msgf("got signal %s", sig)
			cancel()
		}
	}()

	return ctx
}
