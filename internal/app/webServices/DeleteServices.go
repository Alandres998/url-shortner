package webservices

import (
	"net/http"

	"github.com/Alandres998/url-shortner/internal/app/service/auth"
	"github.com/Alandres998/url-shortner/internal/app/service/logger"
	"github.com/Alandres998/url-shortner/internal/app/service/shortener"
	"github.com/gin-gonic/gin"
)

// DeleteShortURL Веб-Сервис по удалению ссылки
func DeleteShortURL(c *gin.Context) {
	var shortURLs []string

	// Получаем список URL для удаления из тела запроса
	if err := c.BindJSON(&shortURLs); err != nil {
		logger.LogError("DeleteShortURL", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Получаем идентификатор пользователя из куки
	userID, err := auth.GetUserID(c)
	if err != nil {
		logger.LogError("DeleteShortURL", err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Запускаем асинхронное удаление
	go shortener.DeleteShortURL(userID, shortURLs)

	// Возвращаем статус 202 Accepted немедленно
	c.Status(http.StatusAccepted)
}
