package models

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
	"github.com/pmpavl/tsyst/pkg/constants"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Passage struct {
	ID     primitive.ObjectID `json:"-" bson:"_id,omitempty"`    // ID в mongoDB
	UserID primitive.ObjectID `json:"-" bson:"userID,omitempty"` // ID пользователя
	TestID primitive.ObjectID `json:"-" bson:"testID,omitempty"` // ID теста

	Path   string           `json:"path" bson:"path"`     // Путь до теста
	Name   string           `json:"name" bson:"name"`     // Название теста
	Points constants.Points `json:"points" bson:"points"` // Баллов для прохождения
	Score  constants.Points `json:"score" bson:"score"`   // Набрано баллов
	Passed bool             `json:"passed" bson:"passed"` // Пройден ли тест

	Tasks []*PassageTask `json:"tasks" bson:"tasks"` // Задачи прохождения

	End time.Time `json:"end" bson:"end"` // Время, когда закончится прохождение

	CreatedAt time.Time `json:"-" bson:"createdAt"`
	UpdatedAt time.Time `json:"-" bson:"updatedAt"`
	DeletedAt time.Time `json:"-" bson:"deletedAt,omitempty"`
}

func NewPassage(userID primitive.ObjectID, test *Test, tasks []*PassageTask) *Passage {
	now := time.Now()

	return &Passage{
		UserID:    userID,
		TestID:    test.ID,
		Path:      test.Path,
		Name:      test.Name,
		Points:    test.Tags.Points,
		Score:     constants.PointsZero,
		Passed:    false,
		Tasks:     tasks,
		End:       test.Tags.TimePassing.End(),
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (p *Passage) Test() *TestPassage {
	return &TestPassage{
		ID:     p.ID,
		Passed: p.Passed,
		Start:  p.CreatedAt,
		End:    p.End,
	}
}

func (p *Passage) IsLastTask(num int) bool { return len(p.Tasks)-1 == num }

func (p *Passage) IsPassed() bool { return p.Points <= p.Score }

func (p *Passage) ActualTaskNum() int {
	for num, task := range p.Tasks {
		if !task.IsAnswered() {
			return num
		}
	}

	return -1
}

func (p *Passage) TimeSpent() constants.Duration {
	timeSpent := time.Now().Sub(p.CreatedAt)

	for _, task := range p.Tasks {
		timeSpent -= task.TimeSpent.Time()
	}

	return constants.Duration(timeSpent)
}

func (p Passage) MarshalJSON() ([]byte, error) {
	bytes, err := json.Marshal(p.marshal())
	if err != nil {
		return nil, errors.Wrap(err, "json marshal")
	}

	return bytes, nil
}

func (p Passage) marshal() any {
	return &struct {
		Path   string         `json:"path"`
		Name   string         `json:"name"`
		Points string         `json:"points"`
		Score  string         `json:"score"`
		Passed bool           `json:"passed"`
		Tasks  []*PassageTask `json:"tasks"`
		Start  time.Time      `json:"start"`
		End    time.Time      `json:"end"`
	}{
		Path:   p.Path,
		Name:   p.Name,
		Points: p.Points.Readable(),
		Score:  p.Score.Readable(),
		Passed: p.Passed,
		Tasks:  p.Tasks,
		Start:  p.CreatedAt.Local(),
		End:    p.End.Local(),
	}
}
