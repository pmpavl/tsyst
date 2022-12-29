package request

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type RefreshRequest struct {
	RefreshToken string `json:"refreshToken"`
}

func GetRefreshRequest(c *gin.Context) (*RefreshRequest, error) {
	var req RefreshRequest

	if err := c.BindJSON(&req); err != nil {
		return nil, errors.Wrap(err, "parse json body")
	}

	return &req, nil
}
