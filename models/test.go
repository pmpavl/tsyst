package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Test struct {
	ID          primitive.ObjectID `json:"-" bson:"_id,omitempty"`               // ID в mongoDB
	Path        string             `json:"path" bson:"path"`                     // Путь
	Name        string             `json:"name" bson:"name"`                     // Название
	Description string             `json:"description" bson:"description"`       // Описание
	Tags        *TestTags          `json:"tags,omitempty" bson:"tags,omitempty"` // Теги теста

	Tasks []*TestTask `json:"tasks,omitempty" bson:"tasks,omitempty"` // Задачи теста

	CreatedAt time.Time `json:"-" bson:"createdAt"`
	UpdatedAt time.Time `json:"-" bson:"updatedAt"`
	DeletedAt time.Time `json:"-" bson:"deletedAt,omitempty"`
}

func NewTest(
	path, name, description string,
	tags *TestTags,
	tasks []*TestTask,
) *Test {
	now := time.Now()

	return &Test{
		Path:        path,
		Name:        name,
		Description: description,
		Tags:        tags,
		Tasks:       tasks,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func (t *Test) Card() *TestCard {
	return &TestCard{
		Path:        t.Path,
		Name:        t.Name,
		Description: t.Description,
		Tags:        t.Tags,
	}
}
