package webservices

import (
	"errors"

	syncservices "github.com/Alandres998/url-shortner/internal/app/db/syncServices"
	"github.com/gin-gonic/gin"
)

// Обработчик для возврата полной строки
func Fuller(c *gin.Context) (string, error) {
	id := c.Param("id")
	originalURL, exists := syncservices.URLStorage.Get(id)
	if !exists {
		return "", errors.New(Error400DefaultText)
	}
	return originalURL, nil
}
