package auth

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pmpavl/tsyst/models"
	"github.com/pmpavl/tsyst/pkg/log"
	"github.com/rs/zerolog"
)

type DBUsers interface {
	Create(ctx context.Context, user models.User) error
	SetTokensByEmail(ctx context.Context, email string, tokens *models.Tokens) error
	SetAccessTokenByEmail(ctx context.Context, email, accessToken string) error
	SetAccessTokenByRefreshToken(ctx context.Context, refreshToken, accessToken string) error
	SearchByEmail(ctx context.Context, email string) (*models.User, error)
	SearchByRefreshToken(ctx context.Context, refreshToken string) (*models.User, error)
	EmailExist(ctx context.Context, email string) (bool, error)
	EmailSalt(ctx context.Context, email string) (string, error)
	AccessTokenExist(ctx context.Context, accessToken string) (bool, error)
	RefreshTokenExist(ctx context.Context, refreshToken string) (bool, error)
}

type Service struct {
	log *zerolog.Logger

	dbUsers DBUsers

	accessTokenMaxAge  uint64
	refreshTokenMaxAge uint64
}

func New(dbUsers DBUsers, accessTokenMaxAge, refreshTokenMaxAge uint64) *Service {
	return &Service{
		log:                log.For("service-auth"),
		dbUsers:            dbUsers,
		accessTokenMaxAge:  accessTokenMaxAge,
		refreshTokenMaxAge: refreshTokenMaxAge,
	}
}

func (s *Service) okResponse(c *gin.Context, response any) { s.response(c, http.StatusOK, response) }

type ErrorDefault struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (s *Service) errorResponse(c *gin.Context, code int, err error) {
	s.response(c, code, ErrorDefault{
		Code:    http.StatusText(code),
		Message: err.Error(),
	})
}

func (s *Service) response(c *gin.Context, code int, response any) { c.JSON(code, response) }
