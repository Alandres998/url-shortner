package fileservices

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"sync"

	"github.com/Alandres998/url-shortner/internal/app/db/storage"
)

type FileStorage struct {
	filePath string
	mu       sync.RWMutex
	data     map[string]storage.URLData
}

func NewFileStorage(filePath string) (*FileStorage, error) {
	fs := &FileStorage{
		filePath: filePath,
		data:     make(map[string]storage.URLData),
	}
	return fs, fs.loadData()
}

func (fs *FileStorage) loadData() error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	file, err := os.Open(fs.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(&fs.data)
}

func (fs *FileStorage) saveData() error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	file, err := os.Create(fs.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(fs.data)
}

func (fs *FileStorage) Set(ctx context.Context, userID, shortURL, originalURL string) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	fs.data[shortURL] = storage.URLData{
		ShortURL:    shortURL,
		OriginalURL: originalURL,
		UserID:      userID,
	}

	return fs.saveData()
}

func (fs *FileStorage) Get(ctx context.Context, shortURL string) (string, error) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()

	urlData, exists := fs.data[shortURL]
	if !exists {
		return "", errors.New("URL not found")
	}
	return urlData.OriginalURL, nil
}

func (fs *FileStorage) GetbyOriginURL(ctx context.Context, originalURL string) (storage.URLData, error) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()

	for _, urlData := range fs.data {
		if urlData.OriginalURL == originalURL {
			return urlData, nil
		}
	}
	return storage.URLData{}, errors.New("URL not found")
}

func (fs *FileStorage) GetUserURLs(ctx context.Context, userID string) ([]storage.URLData, error) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()

	var urls []storage.URLData
	for _, urlData := range fs.data {
		if urlData.UserID == userID {
			urls = append(urls, urlData)
		}
	}
	return urls, nil
}

func (fs *FileStorage) Ping(ctx context.Context) error {
	return nil
}
