package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/pmpavl/tsyst/app/api/auth/request"
	"github.com/pmpavl/tsyst/app/api/auth/response"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Service) PasswordSaltByEmail(c *gin.Context) {
	req, err := request.GetPasswordSaltByEmailRequest(c)
	if err != nil {
		s.errorResponse(c, http.StatusBadRequest, errors.Wrap(err, "get password salt by email request"))

		return //! 400
	}

	user, err := s.dbUsers.ReadByEmail(c.Request.Context(), req.Email)
	if errors.Is(err, mongo.ErrNoDocuments) {
		s.errorResponse(c, http.StatusNotFound, ErrEmailNotExist)

		return //! 404
	} else if err != nil {
		s.errorResponse(c, http.StatusInternalServerError, errors.Wrap(err, "db users read by email"))

		return //! 500
	}

	s.okResponse(c, response.PasswordSaltByEmailResponse{PasswordSalt: user.PasswordSalt}) //! 200
}
