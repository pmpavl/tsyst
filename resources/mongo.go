package resources

import (
	"context"

	"github.com/pkg/errors"
	"github.com/pmpavl/tsyst/pkg/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Resources) initMongo(ctx context.Context) error {
	clientOptions := options.Client().
		SetAppName(r.Env.ServiceName).
		ApplyURI(r.Env.MongoHost)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return errors.Wrap(err, "connect mongo")
	}

	if err := client.Ping(ctx, nil); err != nil {
		return errors.Wrap(err, "ping mongo")
	}

	r.Mongo = client

	log.Logger.Info().Msg("init mongo success")

	return nil
}
