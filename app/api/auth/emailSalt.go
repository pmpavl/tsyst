package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/pmpavl/tsyst/app/api/auth/request"
	"github.com/pmpavl/tsyst/app/api/auth/response"
)

func (s *Service) EmailSalt(c *gin.Context) {
	req, err := request.GetEmailSaltRequest(c)
	if err != nil {
		s.errorResponse(c, http.StatusBadRequest, errors.Wrap(err, "get email salt request"))

		return
	}

	salt, err := s.dbUsers.EmailSalt(c.Request.Context(), req.Email)
	if err != nil {
		s.errorResponse(c, http.StatusInternalServerError, errors.Wrap(err, "db users email salt"))

		return
	}

	s.okResponse(c, response.EmailSaltResponse{Salt: salt})
}
