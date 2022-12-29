package request

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

const (
	queryEmail        string = "email"
	queryAccessToken  string = "access"
	queryRefreshToken string = "refresh"
)

func parseEmail(c *gin.Context) (string, error) {
	email := c.Query(queryEmail)

	if email == "" {
		return email, errors.New("email is required")
	}

	return email, nil
}

func parseAccessToken(c *gin.Context) (string, error) {
	accessToken := c.Query(queryAccessToken)
	if accessToken == "" {
		return accessToken, errors.New("access token is required")
	}

	if _, err := uuid.Parse(accessToken); err != nil {
		return "", errors.Wrap(err, "uuid parse")
	}

	return accessToken, nil
}

func parseRefreshToken(c *gin.Context) (string, error) {
	refreshToken := c.Query(queryRefreshToken)
	if refreshToken == "" {
		return refreshToken, errors.New("refresh token is required")
	}

	if _, err := uuid.Parse(refreshToken); err != nil {
		return "", errors.Wrap(err, "uuid parse")
	}

	return refreshToken, nil
}
