package models

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"github.com/pmpavl/tsyst/pkg/constants"
)

type TestTags struct {
	Complexity constants.ComplexityType `json:"complexity" bson:"complexity"`     // Сложность
	Classes    constants.ClassNumbers   `json:"classes" bson:"classes,omitempty"` // Подходящие классы для теста
	Points     constants.Points         `json:"points" bson:"points"`             // Баллов для прохождения
}

func NewTestTags(
	complexity constants.ComplexityType,
	classes constants.ClassNumbers,
	points constants.Points,
) *TestTags {
	return &TestTags{
		Complexity: complexity,
		Classes:    classes,
		Points:     points,
	}
}

func (t TestTags) MarshalJSON() ([]byte, error) {
	if t.Classes == nil {
		t.Classes = make(constants.ClassNumbers, 0)
	}

	bytes, err := json.Marshal(t.marshal())
	if err != nil {
		return nil, errors.Wrap(err, "json marshal")
	}

	return bytes, nil
}

func (t TestTags) marshal() any {
	return &struct {
		Complexity string `json:"complexity"`
		Classes    string `json:"classes"`
		Points     string `json:"points"`
	}{
		Complexity: t.Complexity.Readable(),
		Classes:    t.Classes.Readable(),
		Points:     fmt.Sprintf("Нужно набрать %s", t.Points.Readable()),
	}
}
