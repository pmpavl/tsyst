package tests

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pmpavl/tsyst/models"
	"github.com/pmpavl/tsyst/pkg/log"
	"github.com/rs/zerolog"
)

const (
	queryPage  string = "page"
	queryName  string = "name"
	queryClass string = "class"
)

type DBTests interface {
	Search(ctx context.Context, page int64, name string, class uint64) ([]*models.Test, error)
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
