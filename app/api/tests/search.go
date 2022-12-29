package tests

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/pmpavl/tsyst/app/api/tests/request"
	"github.com/pmpavl/tsyst/app/api/tests/response"
	"github.com/pmpavl/tsyst/models"
)

func (s *Service) Search(c *gin.Context) {
	req, err := request.GetSearchRequest(c)
	if err != nil {
		s.errorResponse(c, http.StatusBadRequest, errors.Wrap(err, "get search request"))

		return
	}

	countPages, err := s.dbTests.SearchCountPages(c.Request.Context(), req.Name, req.Class)
	if err != nil {
		s.errorResponse(c, http.StatusInternalServerError, errors.Wrap(err, "db tests search count pages"))

		return
	}

	if req.Page > countPages {
		s.errorResponse(c, http.StatusBadRequest, fmt.Errorf("page should be less than count pages: %d", countPages))

		return
	}

	tests, err := s.dbTests.Search(c.Request.Context(), req.Page, req.Name, req.Class)
	if err != nil {
		s.errorResponse(c, http.StatusInternalServerError, errors.Wrap(err, "db tests search"))

		return
	}

	s.okResponse(c, response.SearchResponse{Tests: models.TestsFormat(tests)})
}
