package response

import "github.com/pmpavl/tsyst/models"

type SearchResponse struct {
	Tests []*models.TestFormat `json:"tests"`
}
