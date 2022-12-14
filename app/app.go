package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pmpavl/tsyst/app/api"
	"github.com/pmpavl/tsyst/app/api/auth"
	"github.com/pmpavl/tsyst/app/api/tests"
	"github.com/pmpavl/tsyst/app/router"
	"github.com/pmpavl/tsyst/db"
	"github.com/pmpavl/tsyst/resources"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
)

type App struct {
	log *zerolog.Logger
}

func New(log *zerolog.Logger) *App {
	return &App{log: log}
}

func (a *App) Start(ctx context.Context) error {
	res := resources.New(ctx)

	dbTests := db.NewDBTests(res.Mongo)
	dbUsers := db.NewDBUsers(res.Mongo)

	tests := tests.New(dbTests)
	auth := auth.New(dbUsers, res.Env.AccessTokenMaxAge, res.Env.RefreshTokenMaxAge)

	api := api.New(auth, tests)
	rtr := router.New(api)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", res.Env.ServiceHTTPPort),
		Handler: rtr.Handler(),
	}

	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		a.log.Info().Msgf("start server at %s addr", srv.Addr)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.log.Err(err).Msg("listen and serve")

			return err
		}

		return nil
	})

	group.Go(func() error {
		<-ctx.Done()

		return srv.Shutdown(ctx)
	})

	return group.Wait()
}
