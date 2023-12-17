package service

import "url-shortner/domain"

type domainMetricsService struct {
	repo domain.DomainMetricsRepository
}

//returns new domain metric service
func NewDomainMetricsService(r domain.DomainMetricsRepository) domain.DomainMetricsService {
	return &domainMetricsService{repo: r}
}

//increements domain shorten count
func (s *domainMetricsService) IncreementDomainCountMetric(domain string) error {
	return s.repo.IncreementDomainCountMetric(domain)
}

//return top shortened domain list 
func (s *domainMetricsService) GetTopDomains(limit int) (map[string]int, error) {
	return s.repo.GetTopDomains(limit)
}
