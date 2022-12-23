package db

import (
	"context"

	"github.com/pkg/errors"
	"github.com/pmpavl/tsyst/models"
	"github.com/pmpavl/tsyst/pkg/log"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	CollectionTestsName string = "tests"
	LimitPerPage        int64  = 10
)

type DBTests struct {
	log *zerolog.Logger

	coll *mongo.Collection
}

func NewDBTests(db *mongo.Database) *DBTests {
	return &DBTests{
		log:  log.For("db-tests"),
		coll: db.Collection(CollectionTestsName),
	}
}

func (db *DBTests) Search(ctx context.Context, page int64, name string, class uint64) ([]*models.Test, error) {
	var tests []*models.Test

	filter := bson.M{"name": bson.M{"$regex": name, "$options": "im"}}

	if class != 0 {
		filter["tags.classes"] = class
	}

	opts := options.Find().
		SetSort(bson.M{"name": 1}).
		SetSkip((page - 1) * LimitPerPage).
		SetLimit(LimitPerPage)

	cur, err := db.coll.Find(ctx, filter, opts)
	if err != nil {
		return nil, errors.Wrap(err, "search")
	}

	if err = cur.All(ctx, &tests); err != nil {
		return nil, errors.Wrap(err, "unmarshal search")
	}

	return tests, nil
}
