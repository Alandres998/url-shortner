package webservices

import (
	"errors"

	fileservices "github.com/Alandres998/url-shortner/internal/app/db/fileServices"
	"github.com/gin-gonic/gin"
)

// Обработчик для возврата полной строки
func Fuller(c *gin.Context) (string, error) {
	id := c.Param("id")
	urlShort := fileservices.GetURL(id)
	if urlShort.OriginalURL == "" {
		return "", errors.New(Error400DefaultText)
	}
	return urlShort.OriginalURL, nil
}
