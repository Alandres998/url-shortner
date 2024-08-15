package webservices

import (
	"context"

	"github.com/Alandres998/url-shortner/internal/app/db/storage"
	"github.com/gin-gonic/gin"
)

// Обработчик для возврата полной строки
func Fuller(c *gin.Context) (string, error) {
	ctx := context.Background()
	id := c.Param("id")
	urlOriginal, err := storage.Store.Get(ctx, id)
	if err != nil {
		return "", err
	}
	return urlOriginal, nil
}
