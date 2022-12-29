package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/pmpavl/tsyst/app/api/auth/request"
	"github.com/pmpavl/tsyst/app/api/auth/response"
	"github.com/pmpavl/tsyst/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Service) Registration(c *gin.Context) {
	req, err := request.GetRegistrationRequest(c)
	if err != nil {
		s.errorResponse(c, http.StatusBadRequest, errors.Wrap(err, "get registration request"))

		return
	}

	if err := s.dbUsers.Create(c.Request.Context(), models.User{
		ID:            primitive.NewObjectID(),
		Email:         req.Email,
		EmailVerified: false,
		Password:      req.Password,
		Salt:          req.Salt,
		Meta: &models.Meta{
			FirstName: req.Meta.FirstName,
			LastName:  req.Meta.LastName,
			School:    req.Meta.School,
			Class:     req.Meta.Class,
		},
	}); err != nil {
		s.errorResponse(c, http.StatusInternalServerError, errors.Wrap(err, "db users create"))

		return
	}

	// TODO: Send email verification
	// Отправка пользователю письма с проверкой почты

	s.okResponse(c, response.RegistrationResponse{Status: "ok"})
}
