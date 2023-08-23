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
		s.errorResponse(c, http.StatusNotFound, ErrNothingFound)

		return //! 400
	}

	tests, err := s.dbTests.Search(c.Request.Context(), req.Name, req.Class, req.Complexity, req.Page)
	if err != nil {
		s.errorResponse(c, http.StatusInternalServerError, errors.Wrap(err, "db tests search"))

		return //! 500
	}

	user, _ := s.dbUsers.ReadByAccessToken(c.Request.Context(), req.AccessToken)

	cards := make([]*models.TestCard, 0, len(tests))

	for _, test := range tests {
		card := test.Card()

		if user != nil {
			passage, _ := s.dbPassages.SearchLastUserPassage(c.Request.Context(), user.ID, test.ID)
			if passage != nil {
				card.AddLastPassage(passage.Test())
			}
		}

		cards = append(cards, card)
	}

	s.okResponse(c, response.SearchResponse{CountPages: countPages, Cards: cards}) //! 200
}
