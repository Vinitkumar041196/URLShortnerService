package service

import (
	"net/url"
	"testing"
)

func TestGetDomainFromURL(t *testing.T) {
	u1, _ := url.Parse("https://google.com")
	u2, _ := url.Parse("https://googlecom/hello")
	tests := []struct {
		name string
		arg  *url.URL
		want string
	}{
		{
			name: "nil URL",
			arg:  nil,
			want: "",
		},
		{
			name: "valid url",
			arg:  u1,
			want: "google.com",
		},
		{
			name: "url with invali domain",
			arg:  u2,
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDomainFromURL(tt.arg); got != tt.want {
				t.Errorf("GetDomainFromURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
