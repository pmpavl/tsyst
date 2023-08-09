package constants

import (
	"strings"

	"github.com/pkg/errors"
)

var ErrUnexpectedTaskType = errors.New("unexpected task type")

type TaskType string // Тип задачи

const (
	TaskText  TaskType = "text"  // Текстовая задача
	TaskRadio TaskType = "radio" // Задача с выбором ответа
)

func (t TaskType) String() string { return string(t) }

func (t TaskType) Equal(tt TaskType) bool {
	return strings.Compare(t.String(), tt.String()) == 0
}

func (t TaskType) Validate() error {
	switch t {
	case TaskText:
		return nil
	case TaskRadio:
		return nil
	default:
		return ErrUnexpectedTaskType
	}
}

func (t TaskType) Readable() string {
	switch t {
	case TaskText:
		return "Текстовая задача"
	case TaskRadio:
		return "Задача с выбором ответа"
	default:
		return "Неопределенный тип задачи"
	}
}
