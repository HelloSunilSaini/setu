package requestdto

type GroupCreationRequest struct {
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Users       []string `json:"users,omitempty"`
}

type GroupUsersRequest struct {
	GroupID string   `json:"groupID,omitempty"`
	Action  string   `json:"action,omitempty"`
	Users   []string `json:"users,omitempty"`
}
