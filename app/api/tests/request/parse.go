package request

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	queryPage  string = "page"
	queryName  string = "name"
	queryClass string = "class"
	maxClass   uint64 = 11
)

// Parse query page must be in [1, inf), if empty set first page.
func parsePage(c *gin.Context) (int64, error) {
	pageStr := c.Query(queryPage)

	if pageStr == "" {
		pageStr = "1"
	}

	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		return 0, err
	}

	if page <= 0 {
		return 0, fmt.Errorf("page must be in [1, inf), have: %d", page)
	}

	return page, nil
}

// Parse query class must be in [0, 11], if empty set zero class.
func parseClass(c *gin.Context) (uint64, error) {
	classStr := c.Query(queryClass)

	if classStr == "" {
		classStr = "0"
	}

	class, err := strconv.ParseUint(classStr, 10, 64)
	if err != nil {
		return 0, err
	}

	if class > maxClass {
		return 0, fmt.Errorf("class must be in [0, 11], have: %d", class)
	}

	return class, nil
}
