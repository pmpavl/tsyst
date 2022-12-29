package request

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type AuthenticationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GetAuthenticationRequest(c *gin.Context) (*AuthenticationRequest, error) {
	var req AuthenticationRequest

	if err := c.BindJSON(&req); err != nil {
		return nil, errors.Wrap(err, "parse json body")
	}

	return &req, nil
}
