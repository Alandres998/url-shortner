package fileservices

import (
	"os"
	"testing"

	"github.com/Alandres998/url-shortner/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestFileStorage(t *testing.T) {
	// Устанавливаем путь и режим для файлового хранилища
	filePath := "/tmp/test-short-url-db.json"
	mode := os.O_RDWR | os.O_CREATE

	config.Options.FileStorage.Mode = mode
	config.Options.FileStorage.Path = filePath

	fs, err := NewFileStorage(filePath)
	assert.NoError(t, err)
	assert.NotNil(t, fs)

	defer os.Remove(filePath)

	shortURL := "testShort"
	originalURL := "http://valhalla.com"
	err = fs.Set(shortURL, originalURL)
	assert.NoError(t, err)

	retrievedOriginalURL, err := fs.Get(shortURL)
	assert.NoError(t, err)
	assert.Equal(t, originalURL, retrievedOriginalURL)

	_, err = fs.Get("Не существующий адрес")
	assert.Error(t, err)
	assert.Equal(t, "такого адреса нет", err.Error())

	urlData, err := fs.GetbyOriginURL(originalURL)
	assert.NoError(t, err)
	assert.Equal(t, shortURL, urlData.ShortURL)
	assert.Equal(t, originalURL, urlData.OriginalURL)

	// Тест 6: Проверяем Ping
	err = fs.Ping()
	assert.NoError(t, err)
}
