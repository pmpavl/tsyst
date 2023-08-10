package auth

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pmpavl/tsyst/models"
	"github.com/pmpavl/tsyst/pkg/log"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrAccessTokenNotExist   = errors.New("access token not exist")  // Access токен не существует
	ErrRefreshTokenNotExist = errors.New("refresh token not exist") // Refresh токен не существует
	ErrEmailNotExist        = errors.New("email not exist")         // Почта не существует
	ErrEmailAlreadyExist    = errors.New("email already exist")     // Почта уже используется
	ErrEmailNotVerified     = errors.New("email not verified")      // Почта не верефицирована
	ErrWrongPassword        = errors.New("wrong password")          // Неверный пароль
)

type DBUsers interface {
	Create(ctx context.Context, user *models.User) (primitive.ObjectID, error)
	ReadByEmail(ctx context.Context, email string) (*models.User, error)
	ReadByAccessToken(ctx context.Context, accessToken string) (*models.User, error)
	ReadByRefreshToken(ctx context.Context, refreshToken string) (*models.User, error)
	UpdateTokensByID(ctx context.Context, id primitive.ObjectID, tokens *models.UserTokens) error
	UpdateAccessTokenByID(ctx context.Context, id primitive.ObjectID, accessToken string) error
}

type Service struct {
	log *zerolog.Logger

	dbUsers DBUsers
}

func New(dbUsers DBUsers) *Service {
	return &Service{
		log:     log.For("service-auth"),
		dbUsers: dbUsers,
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
