package structs

type LoginResponse struct {
	ID    int    `json:"id"`
	Token string `json:"token"`
}
