package models

import (
	"math/rand"
	"time"

	"github.com/pmpavl/tsyst/pkg/constants"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID        primitive.ObjectID `json:"-" bson:"_id,omitempty"`               // ID в mongoDB
	Condition string             `json:"condition" bson:"condition"`           // Условие, LaTeX
	Answer    string             `json:"answer" bson:"answer"`                 // Ответ
	Tags      *TaskTags          `json:"tags,omitempty" bson:"tags,omitempty"` // Теги задачи

	Radio []string `json:"radio,omitempty" bson:"radio,omitempty"` // Варианты ответа, LaTeX

	CreatedAt time.Time `json:"-" bson:"createdAt"`
	UpdatedAt time.Time `json:"-" bson:"updatedAt"`
	DeletedAt time.Time `json:"-" bson:"deletedAt,omitempty"`
}

func NewTask(
	condition, answer string,
	tags *TaskTags,
	radio []string,
) *Task {
	now := time.Now()

	return &Task{
		Condition: condition,
		Answer:    answer,
		Tags:      tags,
		Radio:     radio,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// Если задача с выбором ответа, то перемешиваем варианты ответа.
func (t *Task) ShuffleRadio() *Task {
	if t.Tags.Type == constants.TaskRadio && t.Radio != nil {
		rand.Shuffle(len(t.Radio), func(i, j int) {
			t.Radio[i], t.Radio[j] = t.Radio[j], t.Radio[i]
		})
	}

	return t
}
