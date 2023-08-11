package models

import (
	"time"

	"github.com/pmpavl/tsyst/pkg/constants"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Passage struct {
	ID     primitive.ObjectID `json:"-" bson:"_id,omitempty"`    // ID в mongoDB
	UserID primitive.ObjectID `json:"-" bson:"userID,omitempty"` // ID пользователя
	TestID primitive.ObjectID `json:"-" bson:"testID,omitempty"` // ID теста

	Points constants.Points `json:"points" bson:"points"` // Баллов для прохождения
	Score  constants.Points `json:"score" bson:"score"`   // Набрано баллов

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
		Points:    test.Tags.Points,
		Score:     constants.PointsZero,
		Tasks:     tasks,
		End:       test.Tags.TimePassing.End(),
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (p *Passage) IsLastTask(num int) bool { return len(p.Tasks)-1 == num }

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
