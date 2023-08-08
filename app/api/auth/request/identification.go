package request

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/pmpavl/tsyst/app/parser"
)

type IdentificationRequest struct {
	AccessToken string
}

func GetIdentificationRequest(c *gin.Context) (*IdentificationRequest, error) {
	accessToken, err := parser.ParseAccessToken(c)
	if err != nil {
		return nil, errors.Wrap(err, "parse access token")
	}

	return &IdentificationRequest{AccessToken: accessToken}, nil
}
