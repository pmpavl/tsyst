package request

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type SearchRequest struct {
	Page  int64
	Name  string
	Class uint64
}

func GetSearchRequest(c *gin.Context) (*SearchRequest, error) {
	page, err := parsePage(c)
	if err != nil {
		return nil, errors.Wrap(err, "parse page")
	}

	class, err := parseClass(c)
	if err != nil {
		return nil, errors.Wrap(err, "parse class")
	}

	return &SearchRequest{
		Page:  page,
		Name:  c.Query(queryName),
		Class: class,
	}, nil
}
