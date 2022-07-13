package utils

import "testing"

func TestIsStringInArray(t *testing.T) {
	type args struct {
		str string
		arr []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				str: "hello",
				arr: []string{"hell", "hello"},
			},
			want: true,
		},
		{
			name: "failure",
			args: args{
				str: "hekkk",
				arr: []string{"hell", "hello"},
			},
			want: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsStringInArray(tt.args.str, tt.args.arr); got != tt.want {
				t.Errorf("IsStringInArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
