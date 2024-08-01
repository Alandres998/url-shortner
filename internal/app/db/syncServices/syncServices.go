package syncservices

import (
	"errors"
	"sync"

	"github.com/Alandres998/url-shortner/internal/app/db/storage"
)

type URLMap struct {
	s sync.RWMutex
	m map[string]string
}

func NewMemoryStorage() storage.Storage {
	return &URLMap{
		m: make(map[string]string),
	}
}

func (Store *URLMap) Set(key string, value string) error {
	Store.s.Lock()
	defer Store.s.Unlock()
	Store.m[key] = value
	return nil
}

func (Store *URLMap) Get(key string) (string, error) {
	Store.s.RLock()
	defer Store.s.RUnlock()
	value, exists := Store.m[key]
	if !exists {
		return "", errors.New("key not found")
	}
	return value, nil
}

func (store *URLMap) GetbyOriginURL(originalURL string) (storage.URLData, error) {
	store.s.RLock()
	defer store.s.RUnlock()

	for key, data := range store.m {
		if data == originalURL {
			return storage.URLData{ShortURL: key, OriginalURL: data}, nil
		}
	}
	return storage.URLData{}, nil
}
