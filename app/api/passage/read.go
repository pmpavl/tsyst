package passage

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/pmpavl/tsyst/app/api/passage/request"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Service) Read(c *gin.Context) {
	req, err := request.GetReadRequest(c)
	if err != nil {
		s.errorResponse(c, http.StatusBadRequest, errors.Wrap(err, "get read request"))

		return //! 400
	}

	user, err := s.dbUsers.ReadByAccessToken(c.Request.Context(), req.AccessToken)
	if errors.Is(err, mongo.ErrNoDocuments) {
		s.errorResponse(c, http.StatusForbidden, ErrAccessTokenNotExist)

		return //! 403
	} else if err != nil {
		s.errorResponse(c, http.StatusInternalServerError, errors.Wrap(err, "db users read by access token"))

		return //! 500
	}

	passage, err := s.dbPassages.Read(c.Request.Context(), req.ID)
	if errors.Is(err, mongo.ErrNoDocuments) {
		s.errorResponse(c, http.StatusBadRequest, ErrPassageIDNotExist)

		return //! 400
	} else if err != nil {
		s.errorResponse(c, http.StatusInternalServerError, errors.Wrap(err, "db passages read"))

		return //! 500
	}

	if user.ID != passage.UserID {
		s.errorResponse(c, http.StatusForbidden, ErrIncorrectPassageUser)

		return //! 403
	}

	s.okResponse(c, passage) //! 200
}
