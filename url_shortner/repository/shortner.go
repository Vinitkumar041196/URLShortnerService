package repository

import (
	"fmt"
	"sync"
	"url-shortner/domain"
)

type inMemoryURLStore struct {
	store map[string]string
	lock  sync.Mutex
}

func NewInMemoryURLStore() domain.URLShortnerRepository {
	return &inMemoryURLStore{
		store: make(map[string]string),
	}
}

// Stores the mapping for short and full URL
func (s *inMemoryURLStore) StoreShortURL(url, shortURL string) error {
	//check if store initialized
	if s.store == nil {
		return fmt.Errorf("store not initialized")
	}

	//acquire a lock on map to avoid simultaneous read write
	s.lock.Lock()
	//releasing lock on return
	defer s.lock.Unlock()

	if _, ok := s.store[shortURL]; !ok {
		//if not found add url to store
		s.store[shortURL] = url
	}
	return nil
}

// Get the full url from store
func (s *inMemoryURLStore) GetFullURL(shortURL string) (string, error) {
	//check if store initialized
	if s.store == nil {
		return "", fmt.Errorf("store not initialized")
	}

	//acquire a lock on map to avoid simultaneous read write
	s.lock.Lock()
	//releasing lock on return
	defer s.lock.Unlock()

	url, ok := s.store[shortURL]
	if !ok { //url not found in map
		return "", fmt.Errorf("url not found")
	}
	return url, nil
}
