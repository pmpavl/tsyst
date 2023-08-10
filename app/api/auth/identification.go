package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/pmpavl/tsyst/app/api/auth/request"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Service) Identification(c *gin.Context) {
	req, err := request.GetIdentificationRequest(c)
	if err != nil {
		s.errorResponse(c, http.StatusBadRequest, errors.Wrap(err, "get identification request"))

		return //! 400
	}

	if _, err := s.dbUsers.ReadByAccessToken(c.Request.Context(), req.AccessToken); errors.Is(err, mongo.ErrNoDocuments) {
		s.errorResponse(c, http.StatusNotFound, ErrAccessTokenNotExist)

		return //! 404
	} else if err != nil {
		s.errorResponse(c, http.StatusInternalServerError, errors.Wrap(err, "db users read by access token"))

		return //! 500
	}

	s.okResponse(c, nil) //! 200
}
