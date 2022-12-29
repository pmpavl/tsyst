package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/pmpavl/tsyst/app/api/auth/request"
	"github.com/pmpavl/tsyst/app/api/auth/response"
)

func (s *Service) Identification(c *gin.Context) {
	req, err := request.GetIdentificationRequest(c)
	if err != nil {
		s.errorResponse(c, http.StatusBadRequest, errors.Wrap(err, "get identification request"))

		return
	}

	ok, err := s.dbUsers.AccessTokenExist(c.Request.Context(), req.AccessToken)
	if err != nil {
		s.errorResponse(c, http.StatusInternalServerError, errors.Wrap(err, "db users access token exist"))

		return
	}

	s.okResponse(c, response.IdentificationResponse{Exist: ok})
}
