package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func Logger(log *zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			logger    *zerolog.Event
			startTime time.Time = time.Now()
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

		logger.
			Str("method", c.Request.Method).
			Str("uri", c.Request.URL.Path).
			Int("size", c.Writer.Size()).
			Float64("latency_ms", float64(time.Since(startTime).Nanoseconds())/float64(time.Millisecond)).
			Int("status", c.Writer.Status()).
			Msg("request")
	}
}
