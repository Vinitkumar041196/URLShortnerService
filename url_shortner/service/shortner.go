package service

import (
	"fmt"
	"net/url"
	"url-shortner/domain"
)

type urlShortnerService struct {
	repo domain.URLShortnerRepository
}

//returns new url shortner service
func NewURLShortnerService(r domain.URLShortnerRepository) domain.URLShortnerService {
	return &urlShortnerService{repo: r}
}

func (s *urlShortnerService) ShortenURL(fullURL string) (string, error) {
	//validate input url
	if fullURL == "" {
		return "", fmt.Errorf("empty url")
	}

	//validate url format
	_, err := url.ParseRequestURI(fullURL)
	if err != nil {
		return "", fmt.Errorf("invalid url")
	}

	//shorten using hash
	shortURL := hashURL(fullURL)

	//store the mapping
	err = s.repo.StoreShortURL(fullURL, shortURL)
	if err != nil {
		return "", err
	}

	return shortURL, nil
}

func (s *urlShortnerService) GetOriginalURL(shortURL string) (string, error) {
	if shortURL == "" {
		return "", fmt.Errorf("empty url")
	}

	fullURL, err := s.repo.GetFullURL(shortURL)
	if err != nil {
		return "", err
	}

	return fullURL, nil
}
