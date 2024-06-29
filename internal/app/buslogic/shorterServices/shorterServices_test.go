package shorterservices

import (
	"testing"
	"unicode"

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
