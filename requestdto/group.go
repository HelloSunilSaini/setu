package requestdto

type GroupCreationRequest struct {
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Users       []string `json:"users,omitempty"`
}
