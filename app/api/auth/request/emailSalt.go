package request

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type EmailSaltRequest struct {
	Email string `json:"email"`
}

func GetEmailSaltRequest(c *gin.Context) (*EmailSaltRequest, error) {
	email, err := parseEmail(c)
	if err != nil {
		return nil, errors.Wrap(err, "parse email")
	}

	return &EmailSaltRequest{
		Email: email,
	}, nil
}
