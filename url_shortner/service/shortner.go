package service

import (
	"fmt"
	"net/url"
	"url-shortner/domain"
)

type urlShortnerService struct {
	urlRepo     domain.URLShortnerRepository
	metricsRepo domain.DomainMetricsRepository
}

// returns new url shortner service
func NewURLShortnerService(urlRepo domain.URLShortnerRepository, metricsRepo domain.DomainMetricsRepository) domain.URLShortnerService {
	return &urlShortnerService{urlRepo: urlRepo, metricsRepo: metricsRepo}
}

func (s *urlShortnerService) ShortenURL(fullURL string) (string, error) {
	//validate input url
	if fullURL == "" {
		return "", fmt.Errorf("empty url")
	}

	//validate url format
	u, err := url.ParseRequestURI(fullURL)
	if !(err == nil && u.Scheme != "" && u.Host != "") {
		return "", fmt.Errorf("invalid url")
	}

	domain := GetDomainFromURL(u)
	if domain == "" {
		return "", fmt.Errorf("invalid url")
	}

	//shorten using hash
	shortURL := hashURL(fullURL)

	//store the mapping
	err = s.urlRepo.StoreShortURL(fullURL, shortURL)
	if err != nil {
		return "", err
	}

	//update metrics
	s.metricsRepo.IncreementDomainCountMetric(domain)

	return shortURL, nil
}

func (s *urlShortnerService) GetOriginalURL(shortURL string) (string, error) {
	if shortURL == "" {
		return "", fmt.Errorf("empty url")
	}

	fullURL, err := s.urlRepo.GetFullURL(shortURL)
	if err != nil {
		return "", err
	}

	return fullURL, nil
}
