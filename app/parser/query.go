package parser

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/pmpavl/tsyst/pkg/constants"
)

const (
	QueryEmail      string = "email"
	QueryName       string = "name"
	QueryClass      string = "class"
	QueryComplexity string = "complexity"
	QueryPage       string = "page"
	QueryPath       string = "path"
)

var (
	ErrUnexpectedEmail = errors.New("unexpected email")
	ErrUnexpectedPage  = errors.New("unexpected page")
	ErrUnexpectedPath  = errors.New("unexpected path")
)

// Парсинг почты из query запроса. Почта не должна быть пустая.
func ParseEmail(c *gin.Context) (string, error) {
	email := c.Query(QueryEmail)
	if email == "" {
		return "", ErrUnexpectedEmail
	}

	return email, nil
}

// Парсинг имени из query запроса. Может быть пустым.
func ParseName(c *gin.Context) string {
	return c.Query(QueryName)
}

// Парсинг класса из query запроса. Если пустой, то 0.
func ParseClass(c *gin.Context) (constants.ClassNumber, error) {
	classStr := c.Query(QueryClass)

	if classStr == "" {
		classStr = "0"
	}

	classInt, err := strconv.ParseInt(classStr, 10, 64)
	if err != nil {
		return constants.ClassZero, errors.Wrap(err, "strconv parse int")
	}

	class := constants.ClassNumber(classInt)
	if err := class.Validate(); err != nil {
		return constants.ClassZero, errors.Wrap(err, "validate class number")
	}

	return class, nil
}

// Парсинг сложности из query запроса. Если пустая, то 0.
func ParseComplexity(c *gin.Context) (constants.ComplexityType, error) {
	complexityStr := c.Query(QueryComplexity)

	if complexityStr == "" {
		complexityStr = "undefined"
	}

	complexity := constants.ComplexityType(complexityStr)
	if err := complexity.Validate(); err != nil {
		return constants.ComplexityUndefined, errors.Wrap(err, "validate complexity type")
	}

	return complexity, nil
}

// Парсинг номера страницы из query запроса. Если пустая, то 1.
func ParsePage(c *gin.Context) (int64, error) {
	pageStr := c.Query(QueryPage)

	if pageStr == "" {
		pageStr = "1"
	}

	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		return 0, errors.Wrap(err, "strconv parse int")
	} else if page <= 0 {
		return 0, ErrUnexpectedPage
	}

	return page, nil
}

// Парсинг пути из запроса. Путь не должен быть пустой.
func ParsePath(c *gin.Context) (string, error) {
	path := c.Query(QueryPath)
	if path == "" {
		return "", ErrUnexpectedPath
	}

	return path, nil
}
