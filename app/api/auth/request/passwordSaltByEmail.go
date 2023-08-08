package request

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/pmpavl/tsyst/app/parser"
)

type PasswordSaltByEmailRequest struct {
	Email string
}

func GetPasswordSaltByEmailRequest(c *gin.Context) (*PasswordSaltByEmailRequest, error) {
	email, err := parser.ParseEmail(c)
	if err != nil {
		return nil, errors.Wrap(err, "parse email")
	}

	return &PasswordSaltByEmailRequest{Email: email}, nil
}
