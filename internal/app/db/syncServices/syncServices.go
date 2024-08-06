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

func (store *URLMap) Set(key string, value string) error {
	store.s.Lock()
	defer store.s.Unlock()
	store.m[key] = value
	return nil
}

func (store *URLMap) Get(key string) (string, error) {
	store.s.RLock()
	defer store.s.RUnlock()
	value, exists := store.m[key]
	if !exists {
		return "", errors.New("Ключ не обнаружен")
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

func (store *URLMap) Ping() error {
	return nil
}
