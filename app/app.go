package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"github.com/pmpavl/tsyst/app/api"
	"github.com/pmpavl/tsyst/app/api/auth"
	"github.com/pmpavl/tsyst/app/api/passage"
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
	var (
		res = resources.New(ctx)

		dbUsers    = db.NewUsers(res.Mongo)
		dbTests    = db.NewTests(res.Mongo)
		dbTasks    = db.NewTasks(res.Mongo)
		dbPassages = db.NewPassages(res.Mongo)

		auth    = auth.New(dbUsers)
		tests   = tests.New(dbTests)
		passage = passage.New(dbUsers, dbTests, dbTasks, dbPassages)

		api = api.New(auth, tests, passage)
		rtr = router.New(api)

		srv = &http.Server{
			Addr:    fmt.Sprintf(":%d", res.Env.ServiceHTTPPort),
			Handler: rtr.Handler(),
		}
	)

	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		a.log.Info().Msgf("start server at %s addr", srv.Addr)

		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.log.Err(err).Msg("listen and serve")

			return errors.Wrap(err, "listen and serve")
		}

		return nil
	})

	group.Go(func() error {
		<-ctx.Done()

		return errors.Wrap(srv.Shutdown(ctx), "shutdown")
	})

	return errors.Wrap(group.Wait(), "group")
}
