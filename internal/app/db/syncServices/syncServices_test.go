package syncservices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMemoryStorage(t *testing.T) {
	storage := NewMemoryStorage()
	assert.NotNil(t, storage)
}

func TestURLMap_SetAndGet(t *testing.T) {
	storage := NewMemoryStorage()

	key := "Yandex"
	value := "testValue"

	err := storage.Set(key, value)
	assert.NoError(t, err)

	retrievedValue, err := storage.Get(key)
	assert.NoError(t, err)
	assert.Equal(t, value, retrievedValue)
}

func TestURLMap_GetNonExistentKey(t *testing.T) {
	storage := NewMemoryStorage()

	_, err := storage.Get("ТутШото")
	assert.Error(t, err)
	assert.Equal(t, "ключ не обнаружен", err.Error())
}

func TestURLMap_GetbyOriginURL(t *testing.T) {
	storage := NewMemoryStorage()

	key := "testKey"
	value := "http://valhalla.com"

	err := storage.Set(key, value)
	assert.NoError(t, err)

	urlData, err := storage.GetbyOriginURL(value)
	assert.NoError(t, err)
	assert.Equal(t, key, urlData.ShortURL)
	assert.Equal(t, value, urlData.OriginalURL)
}

func TestURLMap_GetbyOriginURLNonExistent(t *testing.T) {
	storage := NewMemoryStorage()

	_, err := storage.GetbyOriginURL("nonExistentURL")
	assert.NoError(t, err)
}

func TestURLMap_Ping(t *testing.T) {
	storage := NewMemoryStorage()
	err := storage.Ping()
	assert.NoError(t, err)
}
