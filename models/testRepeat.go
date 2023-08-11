package models

import (
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/pmpavl/tsyst/pkg/constants"
)

type TestRepeat struct {
	Type         constants.RepeatType `json:"type" bson:"type"`                                     // Повторяемость
	TimeToRepeat constants.Duration   `json:"timeToRepeat,omitempty" bson:"timeToRepeat,omitempty"` // Время до повторения
}

func NewTestRepeat(repeatType constants.RepeatType, timeToRepeat constants.Duration) *TestRepeat {
	return &TestRepeat{
		Type:         repeatType,
		TimeToRepeat: timeToRepeat,
	}
}

func (r TestRepeat) MarshalJSON() ([]byte, error) {
	bytes, err := json.Marshal(r.marshal())
	if err != nil {
		return nil, errors.Wrap(err, "json marshal")
	}

	return bytes, nil
}

func (r TestRepeat) marshal() any {
	return &struct {
		Type         string `json:"type"`
		TimeToRepeat string `json:"timeToRepeat"`
	}{
		Type:         r.Type.Readable(),
		TimeToRepeat: r.TimeToRepeat.Readable(),
	}
}
