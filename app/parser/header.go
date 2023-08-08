package parser

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

const (
	HeaderAccessToken  string = "X-Access-Token"
	HeaderRefreshToken string = "X-Refresh-Token"
)

var (
	ErrUnexpectedAccessToken  = errors.New("unexpected access token")
	ErrUnexpectedRefreshToken = errors.New("unexpected refresh token")
)

// Парсинг Access-Token из хедера запроса. Токен должен быть не пустым uuid.
func ParseAccessToken(c *gin.Context) (string, error) {
	accessToken := c.Request.Header.Get(HeaderAccessToken)
	if accessToken == "" {
		return "", ErrUnexpectedAccessToken
	}

	if _, err := uuid.Parse(accessToken); err != nil {
		return "", errors.Wrap(err, "uuid parse")
	}

	c.Set(HeaderAccessToken, accessToken)

	return accessToken, nil
}

// Парсинг Refresh-Token из хедера запроса. Токен должен быть не пустым uuid.
func ParseRefreshToken(c *gin.Context) (string, error) {
	refreshToken := c.Request.Header.Get(HeaderRefreshToken)
	if refreshToken == "" {
		return "", ErrUnexpectedRefreshToken
	}

	if _, err := uuid.Parse(refreshToken); err != nil {
		return "", errors.Wrap(err, "uuid parse")
	}

	c.Set(HeaderRefreshToken, refreshToken)

	return refreshToken, nil
}
