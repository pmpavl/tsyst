package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TestPassage struct {
	ID     primitive.ObjectID `json:"id"`     // ID прохождения
	Passed bool               `json:"passed"` // Пройден ли тест
	Start  time.Time          `json:"start"`
	End    time.Time          `json:"end"`
}

func NewTestPassage(id primitive.ObjectID, passed bool, start, end time.Time) *TestPassage {
	return &TestPassage{
		ID:     id,
		Passed: passed,
		Start:  start,
		End:    end,
	}
}
