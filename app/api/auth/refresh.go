package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/pmpavl/tsyst/app/api/auth/request"
	"github.com/pmpavl/tsyst/app/api/auth/response"
)

var ErrRefreshTokenNotExist error = errors.New("Refresh Token Not Exist")

func (s *Service) Refresh(c *gin.Context) {
	req, err := request.GetRefreshRequest(c)
	if err != nil {
		s.errorResponse(c, http.StatusBadRequest, errors.Wrap(err, "get refresh request"))

		return
	}

	ok, err := s.dbUsers.RefreshTokenExist(c.Request.Context(), req.RefreshToken)
	if err != nil {
		s.errorResponse(c, http.StatusInternalServerError, errors.Wrap(err, "db users refresh token exist"))

		return
	}

	if !ok {
		s.errorResponse(c, http.StatusForbidden, ErrRefreshTokenNotExist)

		return
	}

	accessToken := uuid.New().String()
	if err := s.dbUsers.SetAccessTokenByRefreshToken(c.Request.Context(), req.RefreshToken, accessToken); err != nil {
		s.errorResponse(c, http.StatusInternalServerError, errors.Wrap(err, "db users set access token by refresh token"))

		return
	}

	s.okResponse(c, response.RefreshResponse{
		AccessToken:       accessToken,
		AccessTokenMaxAge: s.accessTokenMaxAge,
	})
}
