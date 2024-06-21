package api

import "account-app-gin/internal/database"

type CreateTagRequest struct {
	Sign string `json:"sign"`
	Kind string `json:"kind"`
	Name string `json:"name"`
}

type UpdateTagRequest struct {
	Sign *string `json:"sign"`
	Kind *string `json:"kind"`
	Name *string `json:"name"`
}

type GetAllTagResponse struct {
	Resources []database.Tag `json:"resources"`
}

type TagSummary struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Sign    string `json:"sign"`
	Kind    string `json:"kind"`
	Summary int    `json:"summary"`
}
type GetTagSummaryWithMonthResponse struct {
	Resources []TagSummary `json:"resources"`
}
