package syncservices

import (
	"context"
	"errors"
	"sync"

	"github.com/Alandres998/url-shortner/internal/app/db/storage"
)

type URLMap struct {
	s sync.RWMutex
	m map[string]storage.URLData
}

func NewMemoryStorage() storage.Storage {
	return &URLMap{
		m: make(map[string]storage.URLData),
	}
}

func (store *URLMap) Set(ctx context.Context, userID, key, value string) error {
	store.s.Lock()
	defer store.s.Unlock()
	store.m[key] = storage.URLData{
		ShortURL:    key,
		OriginalURL: value,
		UserID:      userID,
	}
	return nil
}

func (store *URLMap) Get(ctx context.Context, key string) (string, error) {
	store.s.RLock()
	defer store.s.RUnlock()
	urlData, exists := store.m[key]
	if !exists {
		return "", errors.New("ключ не обнаружен")
	}
	return urlData.OriginalURL, nil
}

func (store *URLMap) GetbyOriginURL(ctx context.Context, originalURL string) (storage.URLData, error) {
	store.s.RLock()
	defer store.s.RUnlock()

	for _, data := range store.m {
		if data.OriginalURL == originalURL {
			return data, nil
		}
	}
	return storage.URLData{}, nil
}

func (store *URLMap) GetUserURLs(ctx context.Context, userID string) ([]storage.URLData, error) {
	store.s.RLock()
	defer store.s.RUnlock()

	var userURLs []storage.URLData
	for _, urlData := range store.m {
		if urlData.UserID == userID {
			userURLs = append(userURLs, urlData)
		}
	}
	return userURLs, nil
}

func (store *URLMap) Ping(ctx context.Context) error {
	return nil
}
