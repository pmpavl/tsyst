package request

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/pmpavl/tsyst/app/parser"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateRequest struct {
	AccessToken string
	ID          primitive.ObjectID

	Num        uint64 `json:"num"`        // Номер задачи в массиве
	AnswerUser string `json:"answerUser"` // Ответ пользователя
}

func GetUpdateRequest(c *gin.Context) (*UpdateRequest, error) {
	var req UpdateRequest

	if err := c.BindJSON(&req); err != nil {
		return nil, errors.Wrap(err, "parse json body")
	}

	accessToken, err := parser.ParseAccessToken(c)
	if err != nil {
		return nil, errors.Wrap(err, "parse access token")
	}

	id, err := parser.ParseID(c)
	if err != nil {
		return nil, errors.Wrap(err, "parse id")
	}

	req.AccessToken = accessToken
	req.ID = id

	return &req, nil
}
