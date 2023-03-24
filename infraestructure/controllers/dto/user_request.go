package dto

type UserCreateRequest struct {
	ID       string `json:"ID,omitempty"`
	Name     string `json:"Name,omitempty"`
	Email    string `json:"Email,omitempty"`
	Password string `json:"Password,omitempty"`
}
