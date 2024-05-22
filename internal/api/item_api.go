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
	UserID     uint      `json:"userId"`
	TagIDs     []uint    `json:"tagIds"`
	Amount     int       `json:"amount"`
	Kind       string    `json:"kind"`
	HappenedAt time.Time `json:"happenedAt"`
}
