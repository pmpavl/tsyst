package request

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type RegistrationRequest struct {
	Email    string                  `json:"email"`
	Password string                  `json:"password"`
	Salt     string                  `json:"salt"`
	Meta     RegistrationMetaRequest `json:"meta"`
}

type RegistrationMetaRequest struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	School    uint64 `json:"school"`
	Class     uint64 `json:"class"`
}

func GetRegistrationRequest(c *gin.Context) (*RegistrationRequest, error) {
	var req RegistrationRequest

	if err := c.BindJSON(&req); err != nil {
		return nil, errors.Wrap(err, "parse json body")
	}

	req.Meta.FirstName = strings.Title(req.Meta.FirstName)
	req.Meta.LastName = strings.Title(req.Meta.LastName)

	return &req, nil
}
