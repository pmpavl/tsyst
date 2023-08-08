//nolint:gochecknoglobals,gochecknoinits
package log

import (
	"log"
	"os"
	"time"

	"github.com/rs/zerolog"
)

var Logger *zerolog.Logger

func init() {
	zerolog.TimeFieldFormat = time.RFC3339

	setGlobalLevel(LevelDefault)
	setGlobalFormat(FormatDefault)
	log.SetFlags(0)
	log.SetOutput(Logger)

	Logger.Info().Msgf("logger initialized with %s level and %s format output", LevelDefault.String(), FormatDefault)
}

func For(source string) *zerolog.Logger {
	logger := Logger.With().Str("source", source).Logger()

	return &logger
}

func SetGlobalLevel(level Level) {
	setGlobalLevel(level)

	Logger.Info().Msgf("logger level change to %s", level.String())
}

func setGlobalLevel(level Level) {
	zerolog.SetGlobalLevel(level)
}

func SetGlobalFormat(format Format) {
	setGlobalFormat(format)

	Logger.Info().Msgf("logger format output change to %s", format)
}

func setGlobalFormat(format Format) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger() // FormatJSON

	if format == FormatConsole { // FormatConsole
		logger = zerolog.New(zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: time.Kitchen,
		}).With().Timestamp().Logger()
	}

	Logger = &logger
}
