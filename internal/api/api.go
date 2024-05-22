package api

type Pager struct {
	Total    int64 `json:"total"`
	Page     int64 `json:"page"`
	PageSize int64 `json:"page_size"`
}

type Error struct {
	Error string `json:"error"`
}

type Resource struct {
	Resource *map[string]interface{} `json:"resource"`
}
