package api

import "account-app-gin/internal/database"

type GetPagedResponse struct {
	Resources []database.Item `json:"resources"`
	Pager     Pager
}
