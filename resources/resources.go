package resources

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/pmpavl/tsyst/pkg/log"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/sync/errgroup"
)

var ErrUnexpectedGinMode = errors.New("unexpected gin mode")

type Resources struct {
	Env   *Env
	Mongo *mongo.Client
}

func New(ctx context.Context) *Resources {
	r := &Resources{}

	if err := r.initDotEnv(); err != nil {
		log.Logger.Fatal().Err(err).Msg("init dotenv")
	}

	if err := r.getEnv(ctx); err != nil {
		log.Logger.Fatal().Err(err).Msg("init env")
	}

	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		return errors.Wrap(r.initMongo(ctx), "init mongo")
	})

	if err := group.Wait(); err != nil {
		log.Logger.Fatal().Err(err).Msg("init resources")
	}

	return r.setLogger().setGinMode()
}

func (r *Resources) setLogger() *Resources {
	level, err := log.ParseLevel(r.Env.LogLevel)
	if err != nil {
		log.Logger.Warn().Err(err).Msg("parse log level")
	}

	if err == nil && level != log.LevelDefault {
		log.SetGlobalLevel(level)
	}

	format, err := log.ParseFormat(r.Env.LogFormat)
	if err != nil {
		log.Logger.Warn().Err(err).Msg("parse log format")
	}

	if err == nil && format != log.FormatDefault {
		log.SetGlobalFormat(format)
	}

	return r
}

func (r *Resources) setGinMode() *Resources {
	switch r.Env.GinMode {
	case gin.DebugMode:
		gin.SetMode(gin.DebugMode)
	case gin.ReleaseMode:
		gin.SetMode(gin.ReleaseMode)
	case gin.TestMode:
		gin.SetMode(gin.TestMode)
	default: // missmatch gin mode
		log.Logger.Warn().
			Err(ErrUnexpectedGinMode).
			Msg("set gin mode")

		return r
	}

	log.Logger.Info().Msgf("gin mode set to %s", gin.Mode())

	return r
}
