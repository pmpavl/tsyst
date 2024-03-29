package models

import (
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/pmpavl/tsyst/pkg/constants"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TestTask struct {
	ID         *primitive.ObjectID      `json:"-" bson:"_id,omitempty"`       // ID задачи в mongoDB
	Complexity constants.ComplexityType `json:"complexity" bson:"complexity"` // Сложность
	Themes     []string                 `json:"themes" bson:"themes"`         // Темы задачи
	Points     constants.Points         `json:"points" bson:"points"`         // Баллов за задачу
}

func NewTestTask(
	id *primitive.ObjectID,
	complexity constants.ComplexityType,
	themes []string,
	points constants.Points,
) *TestTask {
	return &TestTask{
		ID:         id,
		Complexity: complexity,
		Themes:     themes,
		Points:     points,
	}
}

func (t TestTask) MarshalJSON() ([]byte, error) {
	if t.Themes == nil {
		t.Themes = make([]string, 0)
	}

	bytes, err := json.Marshal(t.marshal())
	if err != nil {
		return nil, errors.Wrap(err, "json marshal")
	}

	return bytes, nil
}

func (t TestTask) marshal() any {
	return &struct {
		Complexity string   `json:"complexity"`
		Themes     []string `json:"themes"`
		Points     string   `json:"points"`
	}{
		Complexity: t.Complexity.Readable(),
		Themes:     t.Themes,
		Points:     t.Points.Readable(),
	}
}
