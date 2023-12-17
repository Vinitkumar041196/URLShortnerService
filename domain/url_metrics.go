package domain

type DomainMetricsService interface {
	IncreementDomainShortenCount(domain string) error
	GetTopDomains(limit int) (map[string]int, error)
}

type DomainMetricsRepository interface {
	IncreementDomainShortenCount(domain string) error
	GetTopDomains(limit int) (map[string]int, error)
}
