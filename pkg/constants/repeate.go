package constants

import (
	"strings"

	"github.com/pkg/errors"
)

var ErrUnexpectedRepeatType = errors.New("unexpected repeat type")

type RepeatType string // Повторяемость

const (
	Disposable RepeatType = "disposable" // Тест можно пройти только один раз
	Repeatable RepeatType = "repeatable" // Тест можно пройти повторно
)

func (t RepeatType) String() string { return string(t) }

func (t RepeatType) Equal(tt RepeatType) bool {
	return strings.Compare(t.String(), tt.String()) == 0
}

func (t RepeatType) Validate() error {
	switch t {
	case Disposable:
		return nil
	case Repeatable:
		return nil
	default:
		return ErrUnexpectedRepeatType
	}
}

func (t RepeatType) Readable() string {
	switch t { //nolint:exhaustive
	case Disposable:
		return "Тест можно пройти только один раз"
	case Repeatable:
		return "Тест можно пройти повторно"
	default:
		return "Тест можно пройти только один раз"
	}
}
