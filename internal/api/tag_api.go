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
