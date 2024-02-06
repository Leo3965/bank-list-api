package structs

type LoginRequest struct {
	Identification int    `json:"identification"`
	Password       string `json:"password"`
}
