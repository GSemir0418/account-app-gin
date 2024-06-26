package api

import (
	"account-app-gin/internal/database"
	"time"
)

type GetPagedResponse struct {
	Resources []database.Item `json:"resources"`
	Pager     Pager           `json:"pager"`
}

type CreateItemRequest struct {
	TagID uint `json:"tagId"`

	Amount     int       `json:"amount"`
	Kind       string    `json:"kind"`
	HappenedAt time.Time `json:"happenedAt"`
}
