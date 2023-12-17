package domain

type DomainMetricsService interface {
	IncreementDomainCountMetric(domain string) error
	GetTopDomains(limit int) (map[string]int, error)
}

type DomainMetricsRepository interface {
	IncreementDomainCountMetric(domain string) error
	GetTopDomains(limit int) (map[string]int, error)
}
