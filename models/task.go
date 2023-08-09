package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID        primitive.ObjectID `json:"-" bson:"_id,omitempty"`               // ID в mongoDB
	Condition string             `json:"condition" bson:"condition"`           // Условие
	Answer    string             `json:"answer" bson:"answer"`                 // Ответ
	Tags      *TaskTags          `json:"tags,omitempty" bson:"tags,omitempty"` // Теги задачи

	Radio []string `json:"radio,omitempty" bson:"radio,omitempty"` // Варианты ответа

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
