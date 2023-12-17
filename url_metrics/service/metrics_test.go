package service

import (
	"reflect"
	"testing"
	"url-shortener/domain"
	"url-shortener/url_metrics/repository"
)

func Test_GetTopDomains(t *testing.T) {
	metricStore := repository.NewInMemoryMetricStore()
	metricStore.IncreementDomainCountMetric("google.com")
	metricStore.IncreementDomainCountMetric("google.com")
	metricStore.IncreementDomainCountMetric("google.com")
	metricStore.IncreementDomainCountMetric("fb.com")
	metricStore.IncreementDomainCountMetric("fb.com")
	metricStore.IncreementDomainCountMetric("bing.com")
	metricSrvc := NewDomainMetricsService(metricStore)

	tests := []struct {
		name    string
		srvc    domain.DomainMetricsService
		arg     int
		want    map[string]int
		wantErr error
	}{
		{
			name:    "Limit 1",
			srvc:    metricSrvc,
			arg:     1,
			want:    map[string]int{"google.com": 3},
			wantErr: nil,
		},
		{
			name:    "Limit less than data points",
			srvc:    metricSrvc,
			arg:     2,
			want:    map[string]int{"google.com": 3, "fb.com": 2},
			wantErr: nil,
		},
		{
			name:    "Test Default Limit",
			srvc:    metricSrvc,
			arg:     0,
			want:    map[string]int{"google.com": 3, "fb.com": 2, "bing.com": 1},
			wantErr: nil,
		},
		{
			name:    "Limit Greater than data points",
			srvc:    metricSrvc,
			arg:     10,
			want:    map[string]int{"google.com": 3, "fb.com": 2, "bing.com": 1},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := tt.srvc.GetTopDomains(tt.arg)
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTopDomains() = %v, want %v", got, tt.want)
			}
		})
	}
}
