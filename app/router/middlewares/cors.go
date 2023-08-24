package middlewares

import (
	"time"

	gincors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const maxAge time.Duration = 12 * time.Hour

func CORS() gin.HandlerFunc {
	return gincors.New(gincors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:80", "https://localhost:443"},
		AllowMethods:     []string{"GET", "POST", "PATCH"},
		AllowHeaders:     []string{"Content-Length", "Content-Type", "X-Access-Token", "X-Refresh-Token"},
		AllowCredentials: true,
		AllowWildcard:    true,
		MaxAge:           maxAge,
	})
}
