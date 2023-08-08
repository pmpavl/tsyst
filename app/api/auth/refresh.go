package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/pmpavl/tsyst/app/api/auth/request"
	"github.com/pmpavl/tsyst/app/api/auth/response"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Service) Refresh(c *gin.Context) {
	req, err := request.GetRefreshRequest(c)
	if err != nil {
		s.errorResponse(c, http.StatusBadRequest, errors.Wrap(err, "get refresh request"))

		return //! 400
	}

	user, err := s.dbUsers.ReadByRefreshToken(c.Request.Context(), req.RefreshToken)
	if errors.Is(err, mongo.ErrNoDocuments) {
		s.errorResponse(c, http.StatusNotFound, ErrRefreshTokenNotExist)

		return //! 404
	} else if err != nil {
		s.errorResponse(c, http.StatusInternalServerError, errors.Wrap(err, "db users read by refresh token"))

		return //! 500
	}

	// Обновляем Access-Token пользователя
	user.Tokens.Refresh()

	if err := s.dbUsers.UpdateAccessTokenByID(c.Request.Context(), user.ID, user.Tokens.AccessToken); err != nil {
		s.errorResponse(c, http.StatusInternalServerError, errors.Wrap(err, "db users update access token by id"))

		return //! 500
	}

	s.okResponse(c, response.RefreshResponse{AccessToken: user.Tokens.AccessToken}) //! 200
}
