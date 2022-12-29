package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/pmpavl/tsyst/app/api/auth/request"
	"github.com/pmpavl/tsyst/app/api/auth/response"
)

func (s *Service) EmailExist(c *gin.Context) {
	req, err := request.GetEmailExistRequest(c)
	if err != nil {
		s.errorResponse(c, http.StatusBadRequest, errors.Wrap(err, "get email exist request"))

		return
	}

	ok, err := s.dbUsers.EmailExist(c.Request.Context(), req.Email)
	if err != nil {
		s.errorResponse(c, http.StatusInternalServerError, errors.Wrap(err, "db users email exist"))

		return
	}

	s.okResponse(c, response.EmailExistResponse{Exist: ok})
}
