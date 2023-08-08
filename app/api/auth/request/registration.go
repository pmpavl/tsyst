package request

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type RegistrationRequest struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	PasswordSalt string `json:"passwordSalt"`
}

func GetRegistrationRequest(c *gin.Context) (*RegistrationRequest, error) {
	var req RegistrationRequest

	if err := c.BindJSON(&req); err != nil {
		return nil, errors.Wrap(err, "parse json body")
	}

	return &req, nil
}
