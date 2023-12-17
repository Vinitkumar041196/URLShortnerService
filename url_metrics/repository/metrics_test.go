package repository

import (
	"fmt"
	"reflect"
	"testing"
	"url-shortner/domain"
)

func Test_IncreementDomainCountMetric(t *testing.T) {
	metricStore := NewInMemoryMetricStore()

	tests := []struct {
		name    string
		store   domain.DomainMetricsRepository
		arg     string
		wantErr error
	}{
		{
			name:    "StoreNotInitialized",
			store:   &inMemoryMetricStore{},
			arg:     "google.com",
			wantErr: fmt.Errorf("store not initialized"),
		},
		{
			name:    "ValidTestCase",
			store:   metricStore,
			arg:     "google.com",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := tt.store.IncreementDomainCountMetric(tt.arg)
			if tt.wantErr != nil {
				if gotErr == nil {
					t.Errorf("IncreementDomainCountMetric() error = %v, wantErr %v", gotErr, tt.wantErr)
				} else if tt.wantErr.Error() != gotErr.Error() {
					t.Errorf("IncreementDomainCountMetric() error = %v, wantErr %v", gotErr, tt.wantErr)
				}
			}

			if tt.wantErr == nil && gotErr != nil {
				t.Errorf("IncreementDomainCountMetric() error = %v, wantErr %v", gotErr, tt.wantErr)
			}
		})
	}
}

func Test_GetTopDomains(t *testing.T) {
	metricStore := NewInMemoryMetricStore()
	metricStore.IncreementDomainCountMetric("google.com")
	metricStore.IncreementDomainCountMetric("google.com")
	metricStore.IncreementDomainCountMetric("google.com")
	metricStore.IncreementDomainCountMetric("fb.com")
	metricStore.IncreementDomainCountMetric("fb.com")
	metricStore.IncreementDomainCountMetric("bing.com")

	tests := []struct {
		name    string
		store   domain.DomainMetricsRepository
		arg     int
		want    map[string]int
		wantErr error
	}{
		{
			name:    "StoreNotInitialized",
			store:   &inMemoryMetricStore{},
			arg:     3,
			want:    nil,
			wantErr: fmt.Errorf("store not initialized"),
		},
		{
			name:    "Limit 1",
			store:   metricStore,
			arg:     1,
			want:    map[string]int{"google.com": 3},
			wantErr: nil,
		},
		{
			name:    "Limit less than data points",
			store:   metricStore,
			arg:     2,
			want:    map[string]int{"google.com": 3, "fb.com": 2},
			wantErr: nil,
		},
		{
			name:    "Test Default Limit",
			store:   metricStore,
			arg:     0,
			want:    map[string]int{"google.com": 3, "fb.com": 2, "bing.com": 1},
			wantErr: nil,
		},
		{
			name:    "Limit Greater than data points",
			store:   metricStore,
			arg:     10,
			want:    map[string]int{"google.com": 3, "fb.com": 2, "bing.com": 1},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotData, gotErr := tt.store.GetTopDomains(tt.arg)
			if tt.wantErr != nil {
				if gotErr == nil {
					t.Errorf("GetTopDomains() error = %v, wantErr %v", gotErr, tt.wantErr)
				} else if tt.wantErr.Error() != gotErr.Error() {
					t.Errorf("GetTopDomains() error = %v, wantErr %v", gotErr, tt.wantErr)
				}
			}

			if tt.wantErr == nil && gotErr != nil {
				t.Errorf("GetTopDomains() error = %v, wantErr %v", gotErr, tt.wantErr)
			}

			if !reflect.DeepEqual(gotData, tt.want) {
				t.Errorf("GetTopDomains() = %v, want %v", gotData, tt.want)
			}
		})
	}
}
