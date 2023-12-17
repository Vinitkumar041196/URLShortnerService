package service

import "testing"

func Test_hashURL(t *testing.T) {

	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "empty url",
			arg:  "",
			want: "",
		},
		{
			name: "hash google.com",
			arg:  "https://google.com",
			want: "cv6VxVduxjTiFIFKMqt4VWjtp2o=",
		},
		{
			name: "test consistent output for google.com",
			arg:  "https://google.com",
			want: "cv6VxVduxjTiFIFKMqt4VWjtp2o=",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hashURL(tt.arg); got != tt.want {
				t.Errorf("hashURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
