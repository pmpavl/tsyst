package constants

import (
	"strings"

	"github.com/pkg/errors"
)

var ErrUnexpectedComplexityType = errors.New("unexpected complexity type")

type ComplexityType string // Условная сложность

const (
	ComplexityUndefined ComplexityType = "undefined"
	ComplexityEasy      ComplexityType = "easy"     // Легко
	ComplexityNormal    ComplexityType = "normal"   // Нормально
	ComplexityHard      ComplexityType = "hard"     // Сложно
	ComplexityVeryHard  ComplexityType = "veryHard" // Очень сложно
)

func (t ComplexityType) String() string { return string(t) }

func (t ComplexityType) Equal(tt ComplexityType) bool {
	return strings.Compare(t.String(), tt.String()) == 0
}

func (t ComplexityType) Validate() error {
	switch t {
	case ComplexityUndefined:
		return nil
	case ComplexityEasy:
		return nil
	case ComplexityNormal:
		return nil
	case ComplexityHard:
		return nil
	case ComplexityVeryHard:
		return nil
	default:
		return ErrUnexpectedComplexityType
	}
}

func (t ComplexityType) Readable() string {
	switch t { //nolint:exhaustive
	case ComplexityEasy:
		return "Легко"
	case ComplexityNormal:
		return "Нормально"
	case ComplexityHard:
		return "Сложно"
	case ComplexityVeryHard:
		return "Очень сложно"
	default:
		return "Неопределено"
	}
}
