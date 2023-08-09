package models

import (
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/pmpavl/tsyst/pkg/constants"
)

type TaskTags struct {
	Type       constants.TaskType       `json:"type" bson:"type"`             // Тип
	Complexity constants.ComplexityType `json:"complexity" bson:"complexity"` // Сложность
	Themes     []string                 `json:"themes" bson:"themes"`         // Темы задачи
}

func NewTaskTags(taskType constants.TaskType, complexity constants.ComplexityType, themes []string) *TaskTags {
	return &TaskTags{
		Type:       taskType,
		Complexity: complexity,
		Themes:     themes,
	}
}

func (t TaskTags) MarshalJSON() ([]byte, error) {
	if t.Themes == nil {
		t.Themes = make([]string, 0)
	}

	bytes, err := json.Marshal(t.marshal())
	if err != nil {
		return nil, errors.Wrap(err, "json marshal")
	}

	return bytes, nil
}

func (t TaskTags) marshal() any {
	return &struct {
		Type       string   `json:"type"`
		Complexity string   `json:"complexity"`
		Themes     []string `json:"themes"`
	}{
		Type:       t.Type.Readable(),
		Complexity: t.Complexity.Readable(),
		Themes:     t.Themes,
	}
}
