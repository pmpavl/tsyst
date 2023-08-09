package tests

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/pmpavl/tsyst/app/api/tests/request"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Service) Landing(c *gin.Context) {
	req, err := request.GetLandingRequest(c)
	if err != nil {
		s.errorResponse(c, http.StatusBadRequest, errors.Wrap(err, "get landing request"))

		return //! 400
	}

	test, err := s.dbTests.ReadByPath(c.Request.Context(), req.Path)
	if errors.Is(err, mongo.ErrNoDocuments) {
		s.errorResponse(c, http.StatusNotFound, ErrNothingFound)

		return //! 404
	} else if err != nil {
		s.errorResponse(c, http.StatusInternalServerError, errors.Wrap(err, "db tests read by path"))

		return //! 500
	}

	s.okResponse(c, test) //! 200
}
