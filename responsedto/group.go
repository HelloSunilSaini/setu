package responsedto

import "setu/dao"

type SingleGroupResponse struct {
	GroupDetails dao.Group              `json:"groupDetails,omitempty"`
	Members      []UserCreationResponse `json:"members,omitempty"`
}
