//nolint:gomnd
package models

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Test struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Tags        *Tags              `json:"tags" bson:"tags"`
}

func (t *Test) Format() *TestFormat {
	return &TestFormat{
		ID:          t.ID.Hex(),
		Name:        t.Name,
		Description: t.Description,
		Tags:        t.Tags.Format(),
	}
}

type Tags struct {
	Minute  uint64   `json:"minute" bson:"minute"`
	Classes []uint64 `json:"classes" bson:"classes"`
}

func (t *Tags) Format() *TagsFormat {
	return &TagsFormat{
		Minute:  minuteFormat(t.Minute),
		Classes: classesFormat(t.Classes),
	}
}

type TestFormat struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Tags        *TagsFormat `json:"tags"`
}

type TagsFormat struct {
	Minute  string `json:"minute"`
	Classes string `json:"classes"`
}

func TestsFormat(tests []*Test) []*TestFormat {
	testsFormat := make([]*TestFormat, 0, len(tests))
	for _, test := range tests {
		testsFormat = append(testsFormat, test.Format())
	}

	return testsFormat
}

func minuteFormat(minute uint64) string {
	lastDigit := minute % 10

	if lastDigit > 1 && lastDigit < 5 {
		return fmt.Sprintf("%d минуты", minute)
	} else if lastDigit == 1 {
		return fmt.Sprintf("%d минута", minute)
	}

	return fmt.Sprintf("%d минут", minute)
}

func classesFormat(classes []uint64) string {
	length := len(classes)

	switch length {
	case 0:
		return ""
	case 1:
		return fmt.Sprintf("для %d класса", classes[0])
	}

	enumeration := ""

	for i, class := range classes {
		switch i {
		case 0:
			enumeration += fmt.Sprintf("%d", class)
		case length - 1:
			enumeration += fmt.Sprintf(" и %d", class)
		default:
			enumeration += fmt.Sprintf(" , %d", class)
		}
	}

	return fmt.Sprintf("для %s классов", enumeration)
}
