package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/pmpavl/tsyst/app/api/auth/request"
	"github.com/pmpavl/tsyst/app/api/auth/response"
	"github.com/pmpavl/tsyst/models"
)

var (
	ErrEmailNotVerified error = errors.New("Email Not Verified")
	ErrWrongPassword    error = errors.New("Wrong Password")
)

func (s *Service) Authentication(c *gin.Context) {
	req, err := request.GetAuthenticationRequest(c)
	if err != nil {
		s.errorResponse(c, http.StatusBadRequest, errors.Wrap(err, "get authentication request"))

		return
	}

	user, err := s.dbUsers.SearchByEmail(c.Request.Context(), req.Email)
	if err != nil {
		s.errorResponse(c, http.StatusInternalServerError, errors.Wrap(err, "db users search by email"))

		return
	}

	if !user.EmailVerified {
		s.errorResponse(c, http.StatusForbidden, ErrEmailNotVerified)

		return
	}

	if req.Password != user.Password {
		s.errorResponse(c, http.StatusForbidden, ErrWrongPassword)

		return
	}

	accessToken := uuid.New().String()
	refreshToken := uuid.New().String()

	if err := s.dbUsers.SetTokensByEmail(c.Request.Context(), req.Email, &models.Tokens{
		Access:  accessToken,
		Refresh: refreshToken,
	}); err != nil {
		s.errorResponse(c, http.StatusInternalServerError, errors.Wrap(err, "db users set tokens by email"))

		return
	}

	s.okResponse(c, response.AuthenticationResponse{
		AccessToken:        accessToken,
		AccessTokenMaxAge:  s.accessTokenMaxAge,
		RefreshToken:       refreshToken,
		RefreshTokenMaxAge: s.refreshTokenMaxAge,
	})
}
