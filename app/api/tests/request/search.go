package request

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/pmpavl/tsyst/app/parser"
	"github.com/pmpavl/tsyst/pkg/constants"
)

type SearchRequest struct {
	AccessToken string
	Name        string
	Class       constants.ClassNumber
	Complexity  constants.ComplexityType
	Page        int64
}

func GetSearchRequest(c *gin.Context) (*SearchRequest, error) {
	accessToken, _ := parser.ParseAccessToken(c)

	class, err := parser.ParseClass(c)
	if err != nil {
		return nil, errors.Wrap(err, "parse class")
	}

	complexity, err := parser.ParseComplexity(c)
	if err != nil {
		return nil, errors.Wrap(err, "parse complexity")
	}

	page, err := parser.ParsePage(c)
	if err != nil {
		return nil, errors.Wrap(err, "parse page")
	}

	return &SearchRequest{
		AccessToken: accessToken,
		Name:        parser.ParseName(c),
		Class:       class,
		Complexity:  complexity,
		Page:        page,
	}, nil
}
