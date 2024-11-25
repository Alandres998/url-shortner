package syncservices

import (
	"context"
	"errors"
	"sync"

	"github.com/Alandres998/url-shortner/internal/app/db/storage"
)

// URLMap представляет хранилище для хранения коротких и оригинальных URL в памяти.
type URLMap struct {
	s sync.RWMutex
	m map[string]storage.URLData
}

// NewMemoryStorage создает новое хранилище в памяти.
// Возвращает интерфейс storage.Storage.
func NewMemoryStorage() storage.Storage {
	return &URLMap{
		m: make(map[string]storage.URLData),
	}
}

// Set сохраняет пару ключ-значение в хранилище.
// userID - идентификатор пользователя, ключ - короткий URL, value - оригинальный URL.
// Возвращает ошибку, если не удалось сохранить данные.
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

// Get возвращает оригинальный URL для заданного ключа.
// Возвращает ошибку, если ключ не существует.
func (store *URLMap) Get(ctx context.Context, key string) (string, error) {
	store.s.RLock()
	defer store.s.RUnlock()
	urlData, exists := store.m[key]
	if !exists {
		return "", errors.New("ключ не обнаружен")
	}
	return urlData.OriginalURL, nil
}

// GetbyOriginURL возвращает данные URL по оригинальному URL.
// Возвращает пустую структуру и ошибку, если оригинальный URL не найден.
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

// GetUserURLs возвращает список всех коротких URL, принадлежащих указанному пользователю.
// Возвращает массив URLData и ошибку, если данные не найдены.
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

// Ping проверяет доступность хранилища.
// Возвращает nil, так как хранилище всегда доступно.
func (store *URLMap) Ping(ctx context.Context) error {
	return nil
}

// DeleteUserURL помечает указанные короткие URL как удаленные для указанного пользователя.
// Возвращает ошибку, если не удалось выполнить операцию.
func (store *URLMap) DeleteUserURL(ctx context.Context, shortURLs []string, userID string) error {
	store.s.Lock()
	defer store.s.Unlock()

	for _, shortURL := range shortURLs {
		urlData, exists := store.m[shortURL]
		if !exists || urlData.UserID != userID {
			continue
		}
		urlData.Deleted = true
		store.m[shortURL] = urlData
	}

	return nil
}

// GetStatistics получает информацию о количестве сокращенных ссылок и уникальных пользователях.
func (store *URLMap) GetStatistics(ctx context.Context) (int, int, error) {
	store.s.RLock()
	defer store.s.RUnlock()

	urlCount := 0
	userSet := make(map[string]struct{})

	for _, urlData := range store.m {
		if !urlData.Deleted {
			urlCount++
		}
		if urlData.UserID != "" {
			userSet[urlData.UserID] = struct{}{}
		}
	}

	return urlCount, len(userSet), nil
}
