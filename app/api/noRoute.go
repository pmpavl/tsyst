package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var ErrNoRoute = errors.New("this route not allowed")

func (a *API) NoRoute(c *gin.Context) { a.errorResponse(c, http.StatusNotFound, ErrNoRoute) }
