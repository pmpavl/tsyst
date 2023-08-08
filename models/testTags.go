package models

import (
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/pmpavl/tsyst/pkg/constants"
)

type TestTags struct {
	Classes    constants.ClassNumbers   `json:"classes" bson:"classes,omitempty"` // Подходящие классы для теста
	Complexity constants.ComplexityType `json:"complexity" bson:"complexity"`     // Сложность
}

func NewTestTags(classes constants.ClassNumbers, complexity constants.ComplexityType) *TestTags {
	return &TestTags{
		Classes:    classes,
		Complexity: complexity,
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
		Classes    string `json:"classes"`
		Complexity string `json:"complexity"`
	}{
		Classes:    t.Classes.Readable(),
		Complexity: t.Complexity.Readable(),
	}
}
