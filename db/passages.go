package db

import (
	"context"

	"github.com/pkg/errors"
	"github.com/pmpavl/tsyst/models"
	"github.com/pmpavl/tsyst/pkg/constants"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Passages struct {
	coll *mongo.Collection
}

func NewPassages(client *mongo.Client) *Passages {
	db := client.Database(constants.DatabasePassages.String())

	return &Passages{coll: db.Collection(constants.CollectionPassages.String())}
}

func (db *Passages) Create(ctx context.Context, passage *models.Passage) (primitive.ObjectID, error) {
	res, err := db.coll.InsertOne(ctx, passage)

	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, constants.ErrResultNotObjectID
	}

	return id, errors.Wrap(err, "insert one")
}

func (db *Passages) Read(ctx context.Context, id primitive.ObjectID) (*models.Passage, error) {
	return db.read(ctx, bson.M{"_id": id})
}

func (db *Passages) read(ctx context.Context, filter primitive.M) (*models.Passage, error) {
	var passage models.Passage

	if err := db.coll.FindOne(ctx, filter).Decode(&passage); errors.Is(err, mongo.ErrNoDocuments) {
		return nil, mongo.ErrNoDocuments
	} else if err != nil {
		return nil, errors.Wrap(err, "find one")
	}

	return &passage, nil
}

func (db *Passages) SearchUserPassages(
	ctx context.Context,
	userID, testID primitive.ObjectID,
) ([]*models.Passage, error) {
	var passages []*models.Passage

	cur, err := db.coll.Find(ctx, bson.M{"userID": userID, "testID": testID}, options.Find().SetSort(bson.M{"end": -1}))
	if err != nil {
		return nil, errors.Wrap(err, "find")
	} else if err := cur.All(ctx, &passages); err != nil {
		return nil, errors.Wrap(err, "unmarshal all")
	}

	if len(passages) == 0 {
		return nil, mongo.ErrNoDocuments
	}

	return passages, nil
}

// Поиск последнего прохождения пользователем теста.
func (db *Passages) SearchLastUserPassage(
	ctx context.Context,
	userID, testID primitive.ObjectID,
) (*models.Passage, error) {
	passages, err := db.SearchUserPassages(ctx, userID, testID)
	if err != nil {
		return nil, err
	}

	return passages[0], err
}

func (db *Passages) Update(ctx context.Context, passage *models.Passage) error {
	return db.update(ctx, bson.M{"_id": passage.ID}, bson.M{"$set": passage})
}

func (db *Passages) update(ctx context.Context, filter, update primitive.M) error {
	_, err := db.coll.UpdateOne(ctx, filter, update)

	return errors.Wrap(err, "update one")
}
