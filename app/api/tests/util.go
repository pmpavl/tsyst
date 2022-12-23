package tests

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	zero string = "0"
	one  string = "1"

	maxClass uint64 = 11
)

// Page in [1, inf].
func (s *Service) parsePage(c *gin.Context) (int64, error) {
	pageStr := c.Query(queryPage)

	if pageStr == "" {
		pageStr = one
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

// Class in [0, 11]. If zero no search by class.
func (s *Service) parseClass(c *gin.Context) (uint64, error) {
	classStr := c.Query(queryClass)

	if classStr == "" {
		classStr = zero
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
