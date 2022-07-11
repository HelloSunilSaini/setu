package requestdto

type UserCreationRequest struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name,omitempty"`
}

type UserValidationRequest struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserConnectionRequest struct {
	Email string `json:"email,omitempty"`
}
