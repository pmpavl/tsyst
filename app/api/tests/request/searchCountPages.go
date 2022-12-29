package request

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type SearchCountPagesRequest struct {
	Name  string
	Class uint64
}

func GetSearchCountPagesRequest(c *gin.Context) (*SearchCountPagesRequest, error) {
	class, err := parseClass(c)
	if err != nil {
		return nil, errors.Wrap(err, "parse class")
	}

	return &SearchCountPagesRequest{
		Name:  c.Query(queryName),
		Class: class,
	}, nil
}
