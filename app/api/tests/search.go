package tests

import (
	"fmt"
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
	name := c.Query(queryName)

	page, err := s.parsePage(c)
	if err != nil {
		return nil, errors.Wrap(err, "parse page")
	}

	class, err := s.parseClass(c)
	if err != nil {
		return nil, errors.Wrap(err, "parse class")
	}

	countPages, err := s.dbTests.SearchCountPages(c.Request.Context(), name, class)
	if err != nil {
		return nil, errors.Wrap(err, "search count pages")
	}

	if page > countPages {
		return nil, fmt.Errorf("page should be less than count pages: %d", countPages)
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
		s.errorResponse(c, http.StatusInternalServerError, errors.Wrap(err, "db tests search"))

		return
	}

	s.okResponse(c, SearchResponse{Tests: models.TestsFormat(tests)})
}

type SearchCountPagesRequest struct {
	name  string
	class uint64
}

func (s *Service) getSearchCountPagesRequest(c *gin.Context) (*SearchCountPagesRequest, error) {
	name := c.Query(queryName)

	class, err := s.parseClass(c)
	if err != nil {
		return nil, errors.Wrap(err, "parse class")
	}

	return &SearchCountPagesRequest{
		name:  name,
		class: class,
	}, nil
}

type SearchCountPagesResponse struct {
	CountPages int64 `json:"countPages"`
}

func (s *Service) SearchCountPages(c *gin.Context) {
	req, err := s.getSearchCountPagesRequest(c)
	if err != nil {
		s.errorResponse(c, http.StatusBadRequest, errors.Wrap(err, "get search count pages request"))

		return
	}

	countPages, err := s.dbTests.SearchCountPages(c.Request.Context(), req.name, req.class)
	if err != nil {
		s.errorResponse(c, http.StatusInternalServerError, errors.Wrap(err, "db tests search count pages"))

		return
	}

	s.okResponse(c, SearchCountPagesResponse{CountPages: countPages})
}
