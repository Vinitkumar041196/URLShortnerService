package repository

import (
	"fmt"
	"testing"
	"url-shortener/domain"
)

func Test_StoreShortURL(t *testing.T) {
	repo := NewInMemoryURLStore()
	repo.StoreShortURL("https://google.com", "cv6VxVduxjTiFIFKMqt4VWjtp2o=")

	type args struct {
		url      string
		shortURL string
	}
	tests := []struct {
		name    string
		repo    domain.URLShortenerRepository
		args    args
		wantErr error
	}{
		{
			name:    "Store not initialized",
			repo:    &inMemoryURLStore{},
			args:    args{url: "https://google.com", shortURL: "cv6VxVduxjTiFIFKMqt4VWjtp2o="},
			wantErr: fmt.Errorf("store not initialized"),
		},
		{
			name:    "ValidTest",
			repo:    repo,
			args:    args{url: "https://google.com", shortURL: "cv6VxVduxjTiFIFKMqt4VWjtp2o="},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := tt.repo.StoreShortURL(tt.args.url, tt.args.shortURL)
			if tt.wantErr != nil {
				if gotErr == nil {
					t.Errorf("StoreShortURL() error = %v, wantErr %v", gotErr, tt.wantErr)
				} else if tt.wantErr.Error() != gotErr.Error() {
					t.Errorf("StoreShortURL() error = %v, wantErr %v", gotErr, tt.wantErr)
				}
			}

			if tt.wantErr == nil && gotErr != nil {
				t.Errorf("StoreShortURL() error = %v, wantErr %v", gotErr, tt.wantErr)
			}
		})
	}
}

func Test_GetFullURL(t *testing.T) {
	repo := NewInMemoryURLStore()
	repo.StoreShortURL("https://google.com", "cv6VxVduxjTiFIFKMqt4VWjtp2o=")

	tests := []struct {
		name    string
		repo    domain.URLShortenerRepository
		arg     string
		want    string
		wantErr error
	}{
		{
			name:    "Store not initialized",
			repo:    &inMemoryURLStore{},
			arg:     "cv6VxVduxjTiFIFKMqt4VWjtp2o=",
			wantErr: fmt.Errorf("store not initialized"),
			want:    "",
		},
		{
			name:    "ValidTest",
			repo:    repo,
			arg:     "cv6VxVduxjTiFIFKMqt4VWjtp2o=",
			wantErr: nil,
			want:    "https://google.com",
		},
		{
			name:    "URL Not found",
			repo:    repo,
			arg:     "cjldNufejYAULxDqn-U7taBUizU=",
			wantErr: fmt.Errorf("url not found"),
			want:    "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := tt.repo.GetFullURL(tt.arg)
			if tt.wantErr != nil {
				if gotErr == nil {
					t.Errorf("GetFullURL() error = %v, wantErr %v", gotErr, tt.wantErr)
				} else if tt.wantErr.Error() != gotErr.Error() {
					t.Errorf("GetFullURL() error = %v, wantErr %v", gotErr, tt.wantErr)
				}
			}

			if tt.wantErr == nil && gotErr != nil {
				t.Errorf("GetFullURL() error = %v, wantErr %v", gotErr, tt.wantErr)
			}

			if got != tt.want {
				t.Errorf("GetFullURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
