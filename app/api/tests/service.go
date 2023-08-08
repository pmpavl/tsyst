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
)

var ErrNothingFound = errors.New("nothing found")

type DBTests interface {
	Search(
		ctx context.Context, name string, class constants.ClassNumber, complexity constants.ComplexityType, page int64,
	) ([]*models.Test, error)
	SearchCountPages(
		ctx context.Context, name string, class constants.ClassNumber, complexity constants.ComplexityType,
	) (int64, error)
}

type Service struct {
	log *zerolog.Logger

	dbTests DBTests
}

func New(dbTests DBTests) *Service {
	return &Service{
		log:     log.For("service-tests"),
		dbTests: dbTests,
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
