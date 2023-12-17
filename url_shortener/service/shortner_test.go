package service

import (
	"fmt"
	"testing"
	metricRepo "url-shortener/url_metrics/repository"
	urlRepo "url-shortener/url_shortener/repository"
)

func Test_ShortenURL(t *testing.T) {

	urlSrvc := NewURLShortenerService(urlRepo.NewInMemoryURLStore(), metricRepo.NewInMemoryMetricStore())

	tests := []struct {
		name    string
		arg     string
		want    string
		wantErr error
	}{
		{
			name:    "Valid URL",
			arg:     "https://google.com",
			want:    "cv6VxVduxjTiFIFKMqt4VWjtp2o=",
			wantErr: nil,
		},
		{
			name:    "Empty URL",
			arg:     "",
			want:    "",
			wantErr: fmt.Errorf("empty url"),
		},
		{
			name:    "Invalid Domain",
			arg:     "https://google//.com",
			want:    "",
			wantErr: fmt.Errorf("invalid url"),
		},
		{
			name:    "Invalid URL",
			arg:     "google.com",
			want:    "",
			wantErr: fmt.Errorf("invalid url"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := urlSrvc.ShortenURL(tt.arg)
			if tt.wantErr != nil {
				if gotErr == nil {
					t.Errorf("ShortenURL() error = %v, wantErr %v", gotErr, tt.wantErr)
				} else if tt.wantErr.Error() != gotErr.Error() {
					t.Errorf("ShortenURL() error = %v, wantErr %v", gotErr, tt.wantErr)
				}
			}

			if tt.wantErr == nil && gotErr != nil {
				t.Errorf("ShortenURL() error = %v, wantErr %v", gotErr, tt.wantErr)
			}

			if got != tt.want {
				t.Errorf("ShortenURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_GetOriginalURL(t *testing.T) {
	urlRepo := urlRepo.NewInMemoryURLStore()
	urlSrvc := NewURLShortenerService(urlRepo, metricRepo.NewInMemoryMetricStore())
	urlSrvc.ShortenURL("https://google.com")

	tests := []struct {
		name    string
		arg     string
		want    string
		wantErr error
	}{
		{
			name:    "Valid URL",
			arg:     "cv6VxVduxjTiFIFKMqt4VWjtp2o=",
			want:    "https://google.com",
			wantErr: nil,
		},
		{
			name:    "Empty URL",
			arg:     "",
			want:    "",
			wantErr: fmt.Errorf("empty url"),
		},
		{
			name:    "URL Not Found",
			arg:     "http://fb.com",
			want:    "",
			wantErr: fmt.Errorf("url not found"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := urlSrvc.GetOriginalURL(tt.arg)
			if tt.wantErr != nil {
				if gotErr == nil {
					t.Errorf("GetOriginalURL() error = %v, wantErr %v", gotErr, tt.wantErr)
				} else if tt.wantErr.Error() != gotErr.Error() {
					t.Errorf("GetOriginalURL() error = %v, wantErr %v", gotErr, tt.wantErr)
				}
			}

			if tt.wantErr == nil && gotErr != nil {
				t.Errorf("GetOriginalURL() error = %v, wantErr %v", gotErr, tt.wantErr)
			}

			if got != tt.want {
				t.Errorf("GetOriginalURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
