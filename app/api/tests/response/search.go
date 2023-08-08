package response

import "github.com/pmpavl/tsyst/models"

type SearchResponse struct {
	CountPages int64              `json:"countPages"`
	Cards      []*models.TestCard `json:"cards"`
}
