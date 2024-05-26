package api

import "account-app-gin/internal/database"

type CreateTagRequest struct {
	UserID uint   `json:"userId"`
	Sign   string `json:"sign"`
	Kind   string `json:"kind"`
	Name   string `json:"name"`
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
	TagID   uint   `json:"tagId"`
	Name    string `json:"name"`
	Sign    string `json:"sign"`
	Kind    string `json:"kind"`
	Summary int    `json:"summary"`
}
type GetTagSummaryWithMonthResponse struct {
	Resources []TagSummary `json:"resources"`
}
