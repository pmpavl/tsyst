package request

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/pmpavl/tsyst/app/parser"
)

type CreateRequest struct {
	AccessToken string
	Path        string // Путь теста
}

func GetCreateRequest(c *gin.Context) (*CreateRequest, error) {
	accessToken, err := parser.ParseAccessToken(c)
	if err != nil {
		return nil, errors.Wrap(err, "parse access token")
	}

	path, err := parser.ParsePath(c)
	if err != nil {
		return nil, errors.Wrap(err, "parse path")
	}

	return &CreateRequest{
		AccessToken: accessToken,
		Path:        path,
	}, nil
}
