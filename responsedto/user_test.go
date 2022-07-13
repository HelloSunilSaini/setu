package responsedto

import (
	"reflect"
	"setu/dao"
	"testing"
)

func TestConvertUserDtoToCreateResponseDto(t *testing.T) {
	type args struct {
		user *dao.User
	}
	tests := []struct {
		name string
		args args
		want *UserCreationResponse
	}{
		{
			name: "success",
			args: args{
				user: &dao.User{
					ID:    "kjbkjbk",
					Name:  "sunil saini",
					Email: "hellosunilsaini@gmail.com",
				},
			},
			want: &UserCreationResponse{
				UserID: "kjbkjbk",
				Name:   "sunil saini",
				Email:  "hellosunilsaini@gmail.com",
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertUserDtoToCreateResponseDto(tt.args.user); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertUserDtoToCreateResponseDto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertDaoUsersToConnectionsResponse(t *testing.T) {
	type args struct {
		users []dao.User
	}
	tests := []struct {
		name string
		args args
		want []UserCreationResponse
	}{
		{
			name: "success",
			args: args{
				users: []dao.User{
					{
						ID:    "kjbkjbk",
						Name:  "sunil saini",
						Email: "hellosunilsaini@gmail.com",
					},
					{
						ID:    "kjbkjbk 1",
						Name:  "sunil saini 1",
						Email: "hellosunilsaini1@gmail.com",
					},
				},
			},
			want: []UserCreationResponse{
				{
					UserID: "kjbkjbk",
					Name:   "sunil saini",
					Email:  "hellosunilsaini@gmail.com",
				},
				{
					UserID: "kjbkjbk 1",
					Name:   "sunil saini 1",
					Email:  "hellosunilsaini1@gmail.com",
				},
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertDaoUsersToConnectionsResponse(tt.args.users); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertDaoUsersToConnectionsResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertUserSessionDtoToUservalidationResponseDto(t *testing.T) {
	type args struct {
		usersession *dao.UserSession
	}
	tests := []struct {
		name string
		args args
		want UserValidattionResponse
	}{
		{
			name: "success",
			args: args{
				usersession: &dao.UserSession{
					ID:        "gchgcgh",
					UserID:    "jhfjhvjvj",
					Expiry:    64567,
					CreatedOn: 64864,
				},
			},
			want: UserValidattionResponse{
				SessionToken: "gchgcgh",
				Expiry:       64567,
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertUserSessionDtoToUservalidationResponseDto(tt.args.usersession); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertUserSessionDtoToUservalidationResponseDto() = %v, want %v", got, tt.want)
			}
		})
	}
}
