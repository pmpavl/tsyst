package api

import (
	"github.com/gin-gonic/gin"
)

type PingResponse struct {
	Message string `json:"message"`
}

func (a *Api) Ping(c *gin.Context) {
	a.okResponse(c, PingResponse{Message: "pong"})
}
