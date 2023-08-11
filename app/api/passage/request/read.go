package request

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/pmpavl/tsyst/app/parser"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReadRequest struct {
	AccessToken string
	ID          primitive.ObjectID
}

func GetReadRequest(c *gin.Context) (*ReadRequest, error) {
	accessToken, err := parser.ParseAccessToken(c)
	if err != nil {
		return nil, errors.Wrap(err, "parse access token")
	}

	id, err := parser.ParseID(c)
	if err != nil {
		return nil, errors.Wrap(err, "parse id")
	}

	return &ReadRequest{
		AccessToken: accessToken,
		ID:          id,
	}, nil
}
