package tests

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pmpavl/tsyst/models"
	"github.com/pmpavl/tsyst/pkg/constants"
	"github.com/pmpavl/tsyst/pkg/log"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ErrNothingFound = errors.New("nothing found")

type DBUsers interface {
	ReadByAccessToken(ctx context.Context, accessToken string) (*models.User, error)
}

type DBTests interface {
	ReadByPath(ctx context.Context, path string) (*models.Test, error)
	Search(
		ctx context.Context, name string, class constants.ClassNumber, complexity constants.ComplexityType, page int64,
	) ([]*models.Test, error)
	SearchCountPages(
		ctx context.Context, name string, class constants.ClassNumber, complexity constants.ComplexityType,
	) (int64, error)
}

type DBPassages interface {
	SearchUserPassages(ctx context.Context, userID, testID primitive.ObjectID) ([]*models.Passage, error)
	SearchLastUserPassage(ctx context.Context, userID, testID primitive.ObjectID) (*models.Passage, error)
}

type Service struct {
	log *zerolog.Logger

	dbUsers    DBUsers
	dbTests    DBTests
	dbPassages DBPassages
}

func New(dbUsers DBUsers, dbTests DBTests, dbPassages DBPassages) *Service {
	return &Service{
		log:        log.For("service-tests"),
		dbUsers:    dbUsers,
		dbTests:    dbTests,
		dbPassages: dbPassages,
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
