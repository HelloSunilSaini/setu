package requestdto

type SettleUpRequest struct {
	GroupID string `json:"groupID,omitempty"`
	UserID  string `json:"userID,omitempty"`
	Comment string `json:"comment,omitempty"`
}
