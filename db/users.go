package db

import (
	"context"

	"github.com/pkg/errors"
	"github.com/pmpavl/tsyst/models"
	"github.com/pmpavl/tsyst/pkg/log"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const CollectionUsersName string = "users"

type DBUsers struct {
	log *zerolog.Logger

	coll *mongo.Collection
}

func NewDBUsers(db *mongo.Database) *DBUsers {
	return &DBUsers{
		log:  log.For("db-users"),
		coll: db.Collection(CollectionUsersName),
	}
}

func (db *DBUsers) Create(ctx context.Context, user models.User) error {
	_, err := db.coll.InsertOne(ctx, user)

	return errors.Wrap(err, "create user")
}

func (db *DBUsers) SetTokensByEmail(ctx context.Context, email string, tokens *models.Tokens) error {
	_, err := db.coll.UpdateOne(ctx, bson.M{"email": email}, bson.M{"$set": bson.M{"tokens": &tokens}})

	return errors.Wrap(err, "update one")
}

func (db *DBUsers) SetAccessTokenByEmail(ctx context.Context, email, accessToken string) error {
	_, err := db.coll.UpdateOne(ctx, bson.M{"email": email}, bson.M{"$set": bson.M{"tokens.access": accessToken}})

	return errors.Wrap(err, "update one")
}

func (db *DBUsers) SetAccessTokenByRefreshToken(ctx context.Context, refreshToken, accessToken string) error {
	_, err := db.coll.UpdateOne(
		ctx,
		bson.M{"tokens.refresh": refreshToken},
		bson.M{"$set": bson.M{"tokens.access": accessToken}},
	)

	return errors.Wrap(err, "update one")
}

func (db *DBUsers) SearchByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	err := db.coll.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, errors.Wrap(err, "find one")
	}

	return &user, nil
}

func (db *DBUsers) SearchByRefreshToken(ctx context.Context, refreshToken string) (*models.User, error) {
	var user models.User

	err := db.coll.FindOne(ctx, bson.M{"tokens.refresh": refreshToken}).Decode(&user)
	if err != nil {
		return nil, errors.Wrap(err, "find one")
	}

	return &user, nil
}

func (db *DBUsers) EmailExist(ctx context.Context, email string) (bool, error) {
	result := db.coll.FindOne(ctx, bson.M{"email": email})
	if result.Err() == mongo.ErrNoDocuments {
		return false, nil
	} else if result.Err() != nil {
		return false, result.Err()
	}

	return true, nil
}

func (db *DBUsers) EmailSalt(ctx context.Context, email string) (string, error) {
	var user models.User

	err := db.coll.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return "", errors.Wrap(err, "find one")
	}

	return user.Salt, nil
}

func (db *DBUsers) AccessTokenExist(ctx context.Context, accessToken string) (bool, error) {
	result := db.coll.FindOne(ctx, bson.M{"tokens.access": accessToken})
	if result.Err() == mongo.ErrNoDocuments {
		return false, nil
	} else if result.Err() != nil {
		return false, result.Err()
	}

	return true, nil
}

func (db *DBUsers) RefreshTokenExist(ctx context.Context, refreshToken string) (bool, error) {
	result := db.coll.FindOne(ctx, bson.M{"tokens.refresh": refreshToken})
	if result.Err() == mongo.ErrNoDocuments {
		return false, nil
	} else if result.Err() != nil {
		return false, result.Err()
	}

	return true, nil
}
