package syncservices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURLMap(t *testing.T) {
	InitURLStorage()

	t.Run("Чекаем успешный кейс", func(t *testing.T) {
		key := "shortURL"
		value := "https://example.com"

		URLStorage.Set(key, value)
		result, exists := URLStorage.Get(key)

		assert.True(t, exists, "Expected key to exist")
		assert.Equal(t, value, result, "Expected value to match")
	})

	t.Run("Чекаем любую бабайку", func(t *testing.T) {
		key := "nonExistentKey"
		result, exists := URLStorage.Get(key)

		assert.False(t, exists, "Expected key not to exist")
		assert.Empty(t, result, "Expected result to be empty")
	})
}
