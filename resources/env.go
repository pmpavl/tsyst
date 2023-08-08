package resources

import (
	"context"

	"github.com/pkg/errors"
	"github.com/pmpavl/tsyst/pkg/log"
	"github.com/sethvargo/go-envconfig"
)

type Env struct {
	ServiceName     string `env:"SERVICE_NAME, default=tsyst"`
	ServiceHTTPPort int    `env:"SERVICE_HTTP_PORT, default=7784"`
	LogLevel        string `env:"LOG_LEVEL, default=debug"`
	LogFormat       string `env:"LOG_FORMAT, default=console"`
	GinMode         string `env:"GIN_MODE, default=release"`
	MongoHost       string `env:"MONGO_HOST, required"`
}

func (r *Resources) getEnv(ctx context.Context) error {
	var env Env

	if err := envconfig.Process(ctx, &env); err != nil {
		return errors.Wrap(err, "envconfig process")
	}

	r.Env = &env

	log.Logger.Info().Msg("init env success")

	return nil
}
