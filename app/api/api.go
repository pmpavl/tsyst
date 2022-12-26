package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pmpavl/tsyst/pkg/log"
	"github.com/rs/zerolog"
)

type Tests interface {
	Search(c *gin.Context)
	SearchCountPages(c *gin.Context)
}

type Api struct {
	log *zerolog.Logger

	tests Tests
}

func New(tests Tests) *Api {
	return &Api{
		log:   log.For("api"),
		tests: tests,
	}
}

func (a *Api) RegisterRoutes(router *gin.Engine) {
	router.NoRoute(a.NoRoute)
	router.GET("/ping", a.Ping)

	tests := router.Group("/tests")
	tests.GET("/search", a.tests.Search)
	tests.GET("/searchCountPages", a.tests.SearchCountPages)
}

func (a *Api) okResponse(c *gin.Context, response any) { a.response(c, http.StatusOK, response) }

type ErrorDefault struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (a *Api) errorResponse(c *gin.Context, code int, err error) {
	a.response(c, code, ErrorDefault{
		Code:    http.StatusText(code),
		Message: err.Error(),
	})
}

func (a *Api) response(c *gin.Context, code int, response any) { c.JSON(code, response) }
