package db

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/pmpavl/tsyst/models"
	"github.com/pmpavl/tsyst/pkg/constants"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Users struct {
	coll *mongo.Collection
}

func NewUsers(client *mongo.Client) *Users {
	db := client.Database(constants.DatabaseCore.String())

	return &Users{coll: db.Collection(constants.CollectionUsers.String())}
}

func (db *Users) Create(ctx context.Context, user *models.User) (primitive.ObjectID, error) {
	res, err := db.coll.InsertOne(ctx, user)

	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, constants.ErrResultNotObjectID
	}

	return id, errors.Wrap(err, "insert one")
}

func (db *Users) ReadByID(ctx context.Context, id primitive.ObjectID) (*models.User, error) {
	return db.read(ctx, bson.M{"_id": id})
}

func (db *Users) ReadByEmail(ctx context.Context, email string) (*models.User, error) {
	return db.read(ctx, bson.M{"email": email})
}

func (db *Users) ReadByAccessToken(ctx context.Context, accessToken string) (*models.User, error) {
	return db.read(ctx, bson.M{"tokens.accessToken": accessToken})
}

func (db *Users) ReadByRefreshToken(ctx context.Context, refreshToken string) (*models.User, error) {
	return db.read(ctx, bson.M{"tokens.refreshToken": refreshToken})
}

func (db *Users) read(ctx context.Context, filter primitive.M) (*models.User, error) {
	var user models.User

	if err := db.coll.FindOne(ctx, filter).Decode(&user); errors.Is(err, mongo.ErrNoDocuments) {
		return nil, mongo.ErrNoDocuments
	} else if err != nil {
		return nil, errors.Wrap(err, "find one")
	}

	return &user, nil
}

func (db *Users) UpdateTokensByID(ctx context.Context, id primitive.ObjectID, tokens *models.UserTokens) error {
	return db.update(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": bson.M{
			"updatedAt": time.Now(),
			"tokens":    tokens,
		}},
	)
}

func (db *Users) UpdateTokensByEmail(ctx context.Context, email string, tokens *models.UserTokens) error {
	return db.update(
		ctx,
		bson.M{"email": email},
		bson.M{"$set": bson.M{
			"updatedAt": time.Now(),
			"tokens":    tokens,
		}},
	)
}

func (db *Users) UpdateAccessTokenByID(ctx context.Context, id primitive.ObjectID, accessToken string) error {
	return db.update(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": bson.M{
			"updatedAt":          time.Now(),
			"tokens.accessToken": accessToken,
		}},
	)
}

func (db *Users) UpdateAccessTokenByEmail(ctx context.Context, email string, accessToken string) error {
	return db.update(
		ctx,
		bson.M{"email": email},
		bson.M{"$set": bson.M{
			"updatedAt":          time.Now(),
			"tokens.accessToken": accessToken,
		}},
	)
}

func (db *Users) update(ctx context.Context, filter, update primitive.M) error {
	_, err := db.coll.UpdateOne(ctx, filter, update)

	return errors.Wrap(err, "update one")
}
