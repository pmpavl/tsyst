package request

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/pmpavl/tsyst/app/parser"
)

type RefreshRequest struct {
	RefreshToken string
}

func GetRefreshRequest(c *gin.Context) (*RefreshRequest, error) {
	refreshToken, err := parser.ParseRefreshToken(c)
	if err != nil {
		return nil, errors.Wrap(err, "parse refresh token")
	}

	return &RefreshRequest{RefreshToken: refreshToken}, nil
}
