package middlewares

import (
	"time"

	gincors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const maxAge time.Duration = 12 * time.Hour

func CORS() gin.HandlerFunc {
	return gincors.New(gincors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		AllowWildcard:    true,
		MaxAge:           maxAge,
	})
}
