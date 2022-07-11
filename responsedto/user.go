package responsedto

import (
	"setu/dao"
)

type UserCreationResponse struct {
	UserID string `json:"userID,omitempty"`
	Name   string `json:"name,omitempty"`
	Email  string `json:"email,omitempty"`
}

func ConvertUserDtoToCreateResponseDto(user *dao.User) *UserCreationResponse {
	return &UserCreationResponse{
		UserID: user.ID,
		Name:   user.Name,
		Email:  user.Email,
	}
}

func ConvertDaoUsersToConnectionsResponse(users []dao.User) []UserCreationResponse {
	resp := []UserCreationResponse{}
	for _, v := range users {
		resp = append(resp, *ConvertUserDtoToCreateResponseDto(&v))
	}
	return resp
}

type UserValidattionResponse struct {
	SessionToken string `json:"sessionToken,omitempty"`
	Expiry       int64  `json:"expiry,omitempty"`
}

func ConvertUserSessionDtoToUservalidationResponseDto(usersession *dao.UserSession) UserValidattionResponse {
	return UserValidattionResponse{
		SessionToken: usersession.ID,
		Expiry:       usersession.Expiry,
	}
}
