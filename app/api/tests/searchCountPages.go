package tests

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/pmpavl/tsyst/app/api/tests/request"
	"github.com/pmpavl/tsyst/app/api/tests/response"
)

func (s *Service) SearchCountPages(c *gin.Context) {
	req, err := request.GetSearchCountPagesRequest(c)
	if err != nil {
		s.errorResponse(c, http.StatusBadRequest, errors.Wrap(err, "get search count pages request"))

		return
	}

	countPages, err := s.dbTests.SearchCountPages(c.Request.Context(), req.Name, req.Class)
	if err != nil {
		s.errorResponse(c, http.StatusInternalServerError, errors.Wrap(err, "db tests search count pages"))

		return
	}

	s.okResponse(c, response.SearchCountPagesResponse{CountPages: countPages})
}
