package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/pmpavl/tsyst/app/api/auth/request"
	"github.com/pmpavl/tsyst/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Service) Authentication(c *gin.Context) {
	req, err := request.GetAuthenticationRequest(c)
	if err != nil {
		s.errorResponse(c, http.StatusBadRequest, errors.Wrap(err, "get authentication request"))

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

	if !user.EmailVerified {
		s.errorResponse(c, http.StatusForbidden, ErrEmailNotVerified)

		return //! 403
	} else if req.Password != user.Password {
		s.errorResponse(c, http.StatusForbidden, ErrWrongPassword)

		return //! 403
	}

	// Новые токены пользователя
	user.Tokens = models.NewUserTokens()

	if err := s.dbUsers.UpdateTokensByID(c.Request.Context(), user.ID, user.Tokens); err != nil {
		s.errorResponse(c, http.StatusInternalServerError, errors.Wrap(err, "db users update tokens by id"))

		return //! 500
	}

	s.okResponse(c, user.Tokens) //! 200
}
