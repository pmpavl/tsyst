package db

import (
	"context"
	"math/rand"

	"github.com/pkg/errors"
	"github.com/pmpavl/tsyst/models"
	"github.com/pmpavl/tsyst/pkg/constants"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Tasks struct {
	coll *mongo.Collection
}

func NewTasks(client *mongo.Client) *Tasks {
	db := client.Database(constants.DatabaseSource.String())

	return &Tasks{coll: db.Collection(constants.CollectionTasks.String())}
}

func (db *Tasks) Create(ctx context.Context, task *models.Task) (primitive.ObjectID, error) {
	res, err := db.coll.InsertOne(ctx, task)

	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, constants.ErrResultNotObjectID
	}

	return id, errors.Wrap(err, "insert one")
}

func (db *Tasks) Read(ctx context.Context, id primitive.ObjectID) (*models.Task, error) {
	return db.read(ctx, bson.M{"_id": id})
}

func (db *Tasks) read(ctx context.Context, filter primitive.M) (*models.Task, error) {
	var task models.Task

	if err := db.coll.FindOne(ctx, filter).Decode(&task); errors.Is(err, mongo.ErrNoDocuments) {
		return nil, mongo.ErrNoDocuments
	} else if err != nil {
		return nil, errors.Wrap(err, "find one")
	}

	return &task, nil
}

// Поиск задач по сложности и темам без уже использованных IDs.
// Перемешивает найденные задачи, для унификации созданных тестов.
func (db *Tasks) Search(
	ctx context.Context,
	complexity constants.ComplexityType,
	themes []string,
	usedIDs []primitive.ObjectID,
) (*models.Task, error) {
	var tasks []*models.Task

	cur, err := db.coll.Find(ctx, db.searchFilter(complexity, themes, usedIDs))
	if err != nil {
		return nil, errors.Wrap(err, "find")
	} else if err := cur.All(ctx, &tasks); err != nil {
		return nil, errors.Wrap(err, "unmarshal all")
	}

	// Если ничего не найдено, возвращаем ошибку
	if len(tasks) == 0 {
		return nil, mongo.ErrNoDocuments
	}

	// Перемешиваем найденные задачи
	rand.Shuffle(len(tasks), func(i, j int) {
		tasks[i], tasks[j] = tasks[j], tasks[i]
	})

	// Берем первую из найденных задач
	return tasks[0].ShuffleRadio(), nil
}

// Поисковой фильтр по сложности и темам без уже использованных IDs.
func (db *Tasks) searchFilter(
	complexity constants.ComplexityType,
	themes []string,
	usedIDs []primitive.ObjectID,
) primitive.M {
	return bson.M{
		"_id":             bson.M{"$nin": usedIDs},
		"tags.complexity": complexity,
		"tags.themes":     bson.M{"$all": themes},
	}
}
