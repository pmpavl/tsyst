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

const LimitPerPage int64 = 20 // Количество задач на страницу

type Tests struct {
	coll *mongo.Collection
}

func NewTests(client *mongo.Client) *Tests {
	db := client.Database(constants.DatabaseSource.String())

	return &Tests{coll: db.Collection(constants.CollectionTests.String())}
}

func (db *Tests) Create(ctx context.Context, test *models.Test) (primitive.ObjectID, error) {
	res, err := db.coll.InsertOne(ctx, test)

	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, constants.ErrResultNotObjectID
	}

	return id, errors.Wrap(err, "insert one")
}

func (db *Tests) ReadByID(ctx context.Context, id primitive.ObjectID) (*models.Test, error) {
	return db.read(ctx, bson.M{"_id": id})
}

func (db *Tests) ReadByPath(ctx context.Context, path string) (*models.Test, error) {
	return db.read(ctx, bson.M{"path": path})
}

func (db *Tests) read(ctx context.Context, filter primitive.M) (*models.Test, error) {
	var test models.Test

	if err := db.coll.FindOne(ctx, filter).Decode(&test); errors.Is(err, mongo.ErrNoDocuments) {
		return nil, mongo.ErrNoDocuments
	} else if err != nil {
		return nil, errors.Wrap(err, "find one")
	}

	return &test, nil
}

// Поиск тестов по имени, классу, сложности и номеру страницы.
func (db *Tests) Search(
	ctx context.Context,
	name string,
	class constants.ClassNumber,
	complexity constants.ComplexityType,
	page int64,
) ([]*models.Test, error) {
	var tests []*models.Test

	opts := options.Find().
		SetSort(bson.M{"name": 1}).         // Сортировка по имени ABC
		SetSkip((page - 1) * LimitPerPage). // Пропускаем первые page - 1 страниц
		SetLimit(LimitPerPage)              // Количество задач на страницу

	cur, err := db.coll.Find(ctx, db.searchFilter(name, class, complexity), opts)
	if err != nil {
		return nil, errors.Wrap(err, "find")
	} else if err := cur.All(ctx, &tests); err != nil {
		return nil, errors.Wrap(err, "unmarshal all")
	}

	return tests, nil
}

// Количество страниц при данном имени, классе и сложности.
func (db *Tests) SearchCountPages(
	ctx context.Context,
	name string,
	class constants.ClassNumber,
	complexity constants.ComplexityType,
) (int64, error) {
	count, err := db.coll.CountDocuments(ctx, db.searchFilter(name, class, complexity))
	if err != nil {
		return 0, errors.Wrap(err, "count documents")
	}

	if count == 0 {
		return 0, nil
	}

	if count%LimitPerPage > 0 {
		return (count / LimitPerPage) + 1, nil
	}

	return (count / LimitPerPage), nil
}

// Поисковой фильтр по имени, классу и сложности.
func (db *Tests) searchFilter(
	name string,
	class constants.ClassNumber,
	complexity constants.ComplexityType,
) primitive.M {
	filter := bson.M{"name": bson.M{"$regex": name, "$options": "im"}}

	if class != constants.ClassZero {
		filter["tags.classes"] = class
	}

	if complexity != constants.ComplexityUndefined {
		filter["tags.complexity"] = complexity
	}

	return filter
}
