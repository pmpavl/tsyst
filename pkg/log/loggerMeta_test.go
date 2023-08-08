//nolint:funlen
package log_test

import (
	"testing"

	"github.com/pmpavl/tsyst/pkg/log"
	"github.com/stretchr/testify/assert"
)

type answParseLevel struct {
	level  string
	errStr string
}

func ttTestParseLevel() []struct {
	name  string
	level string
	answ  answParseLevel
} {
	return []struct {
		name  string
		level string
		answ  answParseLevel
	}{
		{
			name:  "trace level",
			level: "trace",
			answ: answParseLevel{
				level:  "trace",
				errStr: "",
			},
		},
		{
			name:  "debug level",
			level: "debug",
			answ: answParseLevel{
				level:  "debug",
				errStr: "",
			},
		},
		{
			name:  "info level",
			level: "info",
			answ: answParseLevel{
				level:  "info",
				errStr: "",
			},
		},
		{
			name:  "warn level",
			level: "warn",
			answ: answParseLevel{
				level:  "warn",
				errStr: "",
			},
		},
		{
			name:  "error level",
			level: "error",
			answ: answParseLevel{
				level:  "error",
				errStr: "",
			},
		},
		{
			name:  "fatal level",
			level: "fatal",
			answ: answParseLevel{
				level:  "fatal",
				errStr: "",
			},
		},
		{
			name:  "panic level",
			level: "panic",
			answ: answParseLevel{
				level:  "panic",
				errStr: "",
			},
		},
		{
			name:  "nolevel level", // empty level
			level: "",
			answ: answParseLevel{
				level:  "",
				errStr: "",
			},
		},
		{
			name:  "disabled level",
			level: "disabled",
			answ: answParseLevel{
				level:  "disabled",
				errStr: "",
			},
		},
		{
			name:  "wrong level",
			level: "wrong",
			answ: answParseLevel{
				level:  "debug",
				errStr: "unexpected level",
			},
		},
	}
}

func TestParseLevel(t *testing.T) {
	t.Parallel()

	tt := ttTestParseLevel()
	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			level, err := log.ParseLevel(tc.level)

			var errString string
			if err != nil {
				errString = err.Error()
			}

			answ := answParseLevel{level.String(), errString}

			assert.EqualValues(t, answ, tc.answ)
		})
	}
}

type answParseFormat struct {
	format string
	errStr string
}

func ttTestParseFormat() []struct {
	name   string
	format string
	answ   answParseFormat
} {
	return []struct {
		name   string
		format string
		answ   answParseFormat
	}{
		{
			name:   "console format",
			format: "console",
			answ: answParseFormat{
				format: "console",
				errStr: "",
			},
		},
		{
			name:   "json format",
			format: "json",
			answ: answParseFormat{
				format: "json",
				errStr: "",
			},
		},
		{
			name:   "empty format",
			format: "",
			answ: answParseFormat{
				format: "console",
				errStr: "unexpected format",
			},
		},
		{
			name:   "wrong format",
			format: "wrong",
			answ: answParseFormat{
				format: "console",
				errStr: "unexpected format",
			},
		},
	}
}

func TestParseFormat(t *testing.T) {
	t.Parallel()

	tt := ttTestParseFormat()
	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			format, err := log.ParseFormat(tc.format)

			var errString string
			if err != nil {
				errString = err.Error()
			}

			answ := answParseFormat{format, errString}

			assert.EqualValues(t, answ, tc.answ)
		})
	}
}
