package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var ErrNoRoute error = errors.New("This route not allowed")

func (a *Api) NoRoute(c *gin.Context) { a.errorResponse(c, http.StatusBadRequest, ErrNoRoute) }
