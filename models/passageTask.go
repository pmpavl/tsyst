package models

import (
	"encoding/json"
	"strings"

	"github.com/pkg/errors"
	"github.com/pmpavl/tsyst/pkg/constants"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PassageTask struct {
	Task `json:",inline" bson:",inline"` // Задача

	Points     constants.Points   `json:"points" bson:"points"`                             // Баллов за задачу
	Correct    bool               `json:"correct" bson:"correct"`                           // Правильность ответа
	AnswerUser string             `json:"answerUser,omitempty" bson:"answerUser,omitempty"` // Ответ пользователя
	TimeSpent  constants.Duration `json:"timeSpent,omitempty" bson:"timeSpent,omitempty"`   // Времени потрачено
}

func NewPassageTask(task *Task, points constants.Points) *PassageTask {
	return &PassageTask{
		Task:    *task,
		Points:  points,
		Correct: false,
	}
}

func (t *PassageTask) IsAnswered() bool { return t.AnswerUser != "" }
func (t *PassageTask) IsCorrect() bool  { return strings.EqualFold(t.Answer, t.AnswerUser) }

func (t PassageTask) MarshalBSON() ([]byte, error) {
	bytes, err := bson.Marshal(t.marshalBSON())
	if err != nil {
		return nil, errors.Wrap(err, "bson marshal")
	}

	return bytes, nil
}

func (t PassageTask) marshalBSON() any {
	return &struct {
		ID         primitive.ObjectID `bson:"_id,omitempty"`
		Condition  string             `bson:"condition"`
		Answer     string             `bson:"answer"`
		Tags       *TaskTags          `bson:"tags,omitempty"`
		Radio      []string           `bson:"radio,omitempty"`
		Points     constants.Points   `bson:"points"`
		Correct    bool               `bson:"correct"`
		AnswerUser string             `bson:"answerUser,omitempty"`
		TimeSpent  constants.Duration `bson:"timeSpent,omitempty"`
	}{
		ID:         t.ID,
		Condition:  t.Condition,
		Answer:     t.Answer,
		Tags:       t.Tags,
		Radio:      t.Radio,
		Points:     t.Points,
		Correct:    t.Correct,
		AnswerUser: t.AnswerUser,
		TimeSpent:  t.TimeSpent,
	}
}

func (t PassageTask) MarshalJSON() ([]byte, error) {
	bytes, err := json.Marshal(t.marshalJSON())
	if err != nil {
		return nil, errors.Wrap(err, "json marshal")
	}

	return bytes, nil
}

func (t PassageTask) marshalJSON() any {
	return &struct {
		Condition  string    `json:"condition"`
		Answer     string    `json:"answer"`
		Tags       *TaskTags `json:"tags,omitempty"`
		Radio      []string  `json:"radio,omitempty"`
		Points     string    `json:"points"`
		Correct    bool      `json:"correct"`
		AnswerUser string    `json:"answerUser,omitempty"`
		TimeSpent  string    `json:"timeSpent,omitempty"`
	}{
		Condition:  t.Condition,
		Answer:     t.Answer,
		Tags:       t.Tags,
		Radio:      t.Radio,
		Points:     t.Points.Readable(),
		Correct:    t.Correct,
		AnswerUser: t.AnswerUser,
		TimeSpent:  t.TimeSpent.Readable(),
	}
}
