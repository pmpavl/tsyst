package request

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type EmailExistRequest struct {
	Email string `json:"email"`
}

func GetEmailExistRequest(c *gin.Context) (*EmailExistRequest, error) {
	email, err := parseEmail(c)
	if err != nil {
		return nil, errors.Wrap(err, "parse email")
	}

	return &EmailExistRequest{
		Email: email,
	}, nil
}
