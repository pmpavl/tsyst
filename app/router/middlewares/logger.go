package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pmpavl/tsyst/app/parser"
	"github.com/rs/zerolog"
)

func Logger(log *zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			logger    *zerolog.Event
			startTime = time.Now()
		)

		// Before request
		c.Next()
		// After request

		switch c.Writer.Status() {
		case http.StatusOK:
			logger = log.Debug()
		default:
			logger = log.Error()
		}

		if accessToken := c.Request.Header.Get(parser.HeaderAccessToken); accessToken != "" {
			logger.Str(parser.HeaderAccessToken, accessToken)
		}

		if query := c.Request.URL.RawQuery; query != "" {
			logger.Str("query", query)
		}

		logger.
			Str("method", c.Request.Method).
			Str("uri", c.Request.URL.Path).
			Int("size", c.Writer.Size()).
			Float64("latency_ms", float64(time.Since(startTime).Nanoseconds())/float64(time.Millisecond)).
			Int("status", c.Writer.Status()).
			Msg("request")
	}
}
