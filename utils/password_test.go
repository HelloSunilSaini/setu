package utils

import "testing"

func TestGetPasswordHash(t *testing.T) {
	type args struct {
		password string
		key      string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "password 1 success",
			args: args{
				password: "sunil@123",
				key:      "hello",
			},
			want: "b96cfe149da6ee632f295f249f513c2e",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPasswordHash(tt.args.password, tt.args.key); got != tt.want {
				t.Errorf("GetPasswordHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
