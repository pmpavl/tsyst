package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/pmpavl/tsyst/app/api/auth/request"
	"github.com/pmpavl/tsyst/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Service) Registration(c *gin.Context) {
	req, err := request.GetRegistrationRequest(c)
	if err != nil {
		s.errorResponse(c, http.StatusBadRequest, errors.Wrap(err, "get registration request"))

		return //! 400
	}

	if _, err := s.dbUsers.ReadByEmail(c.Request.Context(), req.Email); err == nil {
		s.errorResponse(c, http.StatusConflict, ErrEmailAlreadyExist)

		return //! 409
	} else if !errors.Is(err, mongo.ErrNoDocuments) {
		s.errorResponse(c, http.StatusInternalServerError, errors.Wrap(err, "db users read by email"))

		return //! 500
	}

	user := models.NewUser(req.Email, req.Password, req.PasswordSalt)

	if _, err := s.dbUsers.Create(c.Request.Context(), user); err != nil {
		s.errorResponse(c, http.StatusInternalServerError, errors.Wrap(err, "db users create"))

		return //! 500
	}

	// TODO: Отправка пользователю письма с проверкой почты

	s.okResponse(c, nil) //! 200
}
