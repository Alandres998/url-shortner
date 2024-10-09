package shortener

import (
	"context"
	"testing"
	"unicode"

	"github.com/Alandres998/url-shortner/internal/app/db/storage"
	"github.com/stretchr/testify/assert"
)

func TestGenerateShortURLLength(t *testing.T) {
	shortURL := GenerateShortURL()
	assert.Equal(t, 8, len(shortURL), "Expected short URL length to be 8")
}

func TestGenerateShortURLCharacterSet(t *testing.T) {
	validChars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	shortURL := GenerateShortURL()

	for _, char := range shortURL {
		assert.Contains(t, validChars, string(char), "Generated short URL contains invalid character")
	}
}

func TestGenerateShortURLUnique(t *testing.T) {
	generatedURLs := make(map[string]bool)
	numGenerations := 10000

	for i := 0; i < numGenerations; i++ {
		shortURL := GenerateShortURL()
		if generatedURLs[shortURL] {
			t.Fatalf("Generated duplicate short URL: %s", shortURL)
		}
		generatedURLs[shortURL] = true
	}

	assert.Equal(t, numGenerations, len(generatedURLs), "Expected all generated short URLs to be unique")
}

func TestGenerateShortURLAlphanumeric(t *testing.T) {
	shortURL := GenerateShortURL()
	for _, char := range shortURL {
		assert.True(t, unicode.IsLetter(char) || unicode.IsDigit(char), "Generated short URL contains non-alphanumeric character")
	}
}

/// Бенчмарки

func BenchmarkGenerateShortURL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GenerateShortURL()
	}
}

func BenchmarkDeleteShortURL(b *testing.B) {
	// Создаем мок для хранилища
	mockStore := &storage.MockStorage{
		DeleteUserURLFunc: func(ctx context.Context, urls []string, userID string) error {
			// Здесь вы можете добавить логику, если нужно, или просто вернуть nil
			return nil
		},
	}

	// Устанавливаем мок в хранилище
	storage.Store = mockStore // Предполагается, что у вас есть глобальная переменная Store в пакете storage

	userID := "test_user_id"
	shortURLs := make([]string, 100) // Генерируем 100 коротких URL для теста
	for i := 0; i < 100; i++ {
		shortURLs[i] = GenerateShortURL()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DeleteShortURL(userID, shortURLs)
	}
}
