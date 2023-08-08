package log

import (
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

var (
	ErrUnexpectedLevel  = errors.New("unexpected level")
	ErrUnexpectedFormat = errors.New("unexpected format")
)

type Level = zerolog.Level

const (
	LevelTrace Level = iota - 1
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
	LevelNoLevel
	LevelDisabled

	LevelDefault = LevelDebug
)

func ParseLevel(levelStr string) (Level, error) {
	level, err := zerolog.ParseLevel(levelStr)
	if err != nil {
		return LevelDefault, ErrUnexpectedLevel
	}

	return level, nil
}

type Format = string

const (
	FormatJSON    Format = "json"
	FormatConsole Format = "console"

	FormatDefault = FormatConsole
)

func ParseFormat(formatStr string) (Format, error) {
	switch formatStr {
	case FormatJSON:
		return FormatJSON, nil
	case FormatConsole:
		return FormatConsole, nil
	}

	return FormatDefault, ErrUnexpectedFormat
}
