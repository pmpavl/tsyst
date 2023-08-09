package request

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/pmpavl/tsyst/app/parser"
)

type LandingRequest struct {
	AccessToken string
	Path        string
}

func GetLandingRequest(c *gin.Context) (*LandingRequest, error) {
	accessToken, _ := parser.ParseAccessToken(c)

	path, err := parser.ParsePath(c)
	if err != nil {
		return nil, errors.Wrap(err, "parse path")
	}

	return &LandingRequest{
		AccessToken: accessToken,
		Path:        path,
	}, nil
}
