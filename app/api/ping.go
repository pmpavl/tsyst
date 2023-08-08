package api

import "github.com/gin-gonic/gin"

func (a *API) Ping(c *gin.Context) { a.okResponse(c, "pong") }
