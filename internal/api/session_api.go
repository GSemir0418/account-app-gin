package api

type SessionRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type SessionResponse struct {
	Jwt    string `json:"jwt"`
	UserID uint   `json:"userId"`
}
