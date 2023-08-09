package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pmpavl/tsyst/pkg/log"
	"github.com/rs/zerolog"
)

type Auth interface {
	Registration(c *gin.Context)
	Authentication(c *gin.Context)
	Refresh(c *gin.Context)
	Identification(c *gin.Context)
	PasswordSaltByEmail(c *gin.Context)
}

type Tests interface {
	Search(c *gin.Context)
	Landing(c *gin.Context)
}

type API struct {
	log *zerolog.Logger

	auth  Auth
	tests Tests
}

func New(auth Auth, tests Tests) *API {
	return &API{
		log:   log.For("api"),
		auth:  auth,
		tests: tests,
	}
}

func (a *API) RegisterRoutes(router *gin.Engine) {
	router.NoRoute(a.NoRoute)
	router.GET("/ping", a.Ping)

	auth := router.Group("/auth")
	auth.POST("/registration", a.auth.Registration)
	auth.POST("/authentication", a.auth.Authentication)
	auth.POST("/refresh", a.auth.Refresh)
	auth.GET("/identification", a.auth.Identification)
	auth.GET("/passwordSaltByEmail", a.auth.PasswordSaltByEmail)

	tests := router.Group("/tests")
	tests.GET("/search", a.tests.Search)
	tests.GET("/landing", a.tests.Landing)
}

func (a *API) okResponse(c *gin.Context, response any) { a.response(c, http.StatusOK, response) }

type ErrorDefault struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (a *API) errorResponse(c *gin.Context, code int, err error) {
	a.response(c, code, ErrorDefault{
		Code:    http.StatusText(code),
		Message: err.Error(),
	})
}

func (a *API) response(c *gin.Context, code int, response any) { c.JSON(code, response) }
