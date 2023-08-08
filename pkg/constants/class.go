package constants

import (
	"fmt"

	"github.com/pkg/errors"
)

var ErrUnexpectedClassNumber = errors.New("unexpected class number")

type ClassNumber int64 // Класс 0 или [5, 11]

const (
	ClassZero   ClassNumber = 0
	ClassFive   ClassNumber = 5
	ClassSix    ClassNumber = 6
	ClassSeven  ClassNumber = 7
	ClassEight  ClassNumber = 8
	ClassNine   ClassNumber = 9
	ClassTen    ClassNumber = 10
	ClassEleven ClassNumber = 11
)

func (n ClassNumber) Validate() error {
	switch n {
	case ClassZero:
		return nil
	case ClassFive:
		return nil
	case ClassSix:
		return nil
	case ClassSeven:
		return nil
	case ClassEight:
		return nil
	case ClassNine:
		return nil
	case ClassTen:
		return nil
	case ClassEleven:
		return nil
	default:
		return ErrUnexpectedClassNumber
	}
}

func (n ClassNumber) Readable() string {
	switch n { //nolint:exhaustive
	case ClassZero:
		return "Для любого класса"
	default:
		return fmt.Sprintf("Для %d класса", n)
	}
}

type ClassNumbers []ClassNumber // Список классов

func (ns ClassNumbers) Readable() string {
	switch len(ns) {
	case 0:
		return "Для любого класса"
	case 1:
		return ns[0].Readable()
	default:
		str := fmt.Sprintf("Для %d", ns[0])

		for i := 1; i < len(ns)-1; i++ {
			str += fmt.Sprintf(", %d", ns[i])
		}

		return fmt.Sprintf("%s и %d класса", str, ns[len(ns)-1])
	}
}
