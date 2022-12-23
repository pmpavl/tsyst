package tests

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/pmpavl/tsyst/models"
)

type SearchRequest struct {
	page  int64
	name  string
	class uint64
}

func (s *Service) getSearchRequest(c *gin.Context) (*SearchRequest, error) {
	page, err := s.parsePage(c)
	if err != nil {
		return nil, errors.Wrap(err, "parse page")
	}

	class, err := s.parseClass(c)
	if err != nil {
		return nil, errors.Wrap(err, "parse class")
	}

	return &SearchRequest{
		page:  page,
		name:  c.Query(queryName),
		class: class,
	}, nil
}

type SearchResponse struct {
	Tests []*models.TestFormat `json:"tests"`
}

func (s *Service) Search(c *gin.Context) {
	req, err := s.getSearchRequest(c)
	if err != nil {
		s.errorResponse(c, http.StatusBadRequest, errors.Wrap(err, "get search request"))

		return
	}

	tests, err := s.dbTests.Search(c.Request.Context(), req.page, req.name, req.class)
	if err != nil {
		s.errorResponse(c, http.StatusInternalServerError, errors.Wrap(err, "db tests"))

		return
	}

	s.okResponse(c, SearchResponse{Tests: models.TestsFormat(tests)})
}
