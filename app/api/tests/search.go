package tests

import (
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

		return //! 400
	}

	countPages, err := s.dbTests.SearchCountPages(c.Request.Context(), req.Name, req.Class, req.Complexity)
	if err != nil {
		s.errorResponse(c, http.StatusInternalServerError, errors.Wrap(err, "db tests search count pages"))

		return //! 500
	} else if req.Page > countPages {
		s.errorResponse(c, http.StatusBadRequest, ErrNothingFound)

		return //! 400
	}

	tests, err := s.dbTests.Search(c.Request.Context(), req.Name, req.Class, req.Complexity, req.Page)
	if err != nil {
		s.errorResponse(c, http.StatusInternalServerError, errors.Wrap(err, "db tests search"))

		return //! 500
	}

	cards := make([]*models.TestCard, 0, len(tests))
	for _, t := range tests {
		cards = append(cards, t.Card())
	}

	s.okResponse(c, response.SearchResponse{CountPages: countPages, Cards: cards}) //! 200
}
