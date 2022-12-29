package request

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type IdentificationRequest struct {
	AccessToken string `json:"accessToken"`
}

func GetIdentificationRequest(c *gin.Context) (*IdentificationRequest, error) {
	accessToken, err := parseAccessToken(c)
	if err != nil {
		return nil, errors.Wrap(err, "parse access token")
	}

	return &IdentificationRequest{
		AccessToken: accessToken,
	}, nil
}
