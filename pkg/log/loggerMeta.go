package log

import (
	"fmt"

	"github.com/rs/zerolog"
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
	Level, err := zerolog.ParseLevel(levelStr)
	if err != nil {
		return LevelDefault, fmt.Errorf("no such level: %s", levelStr)
	}

	return Level, nil
}

type Format = string

const (
	FormatJson    Format = "json"
	FormatConsole Format = "console"

	FormatDefault = FormatConsole
)

func ParseFormat(formatStr string) (Format, error) {
	switch formatStr {
	case FormatJson:
		return FormatJson, nil
	case FormatConsole:
		return FormatConsole, nil
	}

	return FormatDefault, fmt.Errorf("no such format: %s", formatStr)
}
