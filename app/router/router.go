package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pmpavl/tsyst/app/router/middlewares"
	"github.com/pmpavl/tsyst/pkg/log"
	"github.com/rs/zerolog"
)

type API interface{ RegisterRoutes(router *gin.Engine) }

type Router struct {
	log *zerolog.Logger

	api API
}

func New(api API) *Router {
	return &Router{
		log: log.For("router"),
		api: api,
	}
}

func (r *Router) Handler() *gin.Engine {
	router := gin.New()
	router.Use(r.middlewares()...)

	r.api.RegisterRoutes(router)

	r.routesLog(router.Routes())

	return router
}

// Массив middleware которые нужно применить к router.
func (r *Router) middlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middlewares.CORS(),
		middlewares.Logger(r.log),
	}
}

// Логирование раутов сервиса.
func (r *Router) routesLog(routes gin.RoutesInfo) {
	routesLog := "init routes success"
	for _, route := range routes {
		routesLog += fmt.Sprintf("\n\t%s: %s", route.Method, route.Path)
	}

	log.Logger.Info().Msg(routesLog)
}
