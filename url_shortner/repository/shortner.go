package repository

import (
	"fmt"
	"sync"
)

type inMemoryURLStore struct {
	store map[string]string
	lock  sync.Mutex
}

var Store *inMemoryURLStore

func NewInMemoryURLStore() *inMemoryURLStore {
	return &inMemoryURLStore{
		store: make(map[string]string),
	}
}

// Stores the mapping for short and full URL
func (s *inMemoryURLStore) StoreShortURL(url, shortURL string) error {
	if s.store == nil {
		return fmt.Errorf("store not initialized")
	}

	s.lock.Lock()
	defer s.lock.Unlock()
	if _, ok := s.store[shortURL]; !ok {
		s.store[shortURL] = url
	}
	return nil
}

// Get the full url from store
func (s *inMemoryURLStore) GetFullURL(shortURL string) (string, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	url, ok := s.store[shortURL]
	if !ok {
		return "", fmt.Errorf("url not found")
	}
	return url, nil
}
