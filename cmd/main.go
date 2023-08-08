package main

import (
	"context"
	"errors"
	"net/http"

	"github.com/pmpavl/tsyst/app"
	"github.com/pmpavl/tsyst/pkg/contextOS"
	"github.com/pmpavl/tsyst/pkg/log"
	_ "go.uber.org/automaxprocs"
)

func main() {
	log := log.For("tsyst")
	ctx := contextOS.Background()
	app := app.New(log)

	if err := app.Start(ctx); !errors.Is(err, nil) &&
		!errors.Is(err, context.Canceled) &&
		!errors.Is(err, http.ErrServerClosed) {
		log.Error().Err(err).Msg("app")
	}

	log.Info().Msg("shutdown service")
}
