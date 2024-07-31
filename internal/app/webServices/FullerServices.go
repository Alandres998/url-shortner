package webservices

import (
	"errors"

	"github.com/Alandres998/url-shortner/internal/app/db/storage"
	"github.com/gin-gonic/gin"
)

// Обработчик для возврата полной строки
func Fuller(c *gin.Context) (string, error) {
	id := c.Param("id")
	urlOriginal, err := storage.Store.Get(id)
	if err != nil {
		return "", errors.New(Error400DefaultText)
	}
	return urlOriginal, nil
}
