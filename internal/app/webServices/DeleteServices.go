package webservices

import (
	"github.com/Alandres998/url-shortner/internal/app/service/auth"
	"github.com/Alandres998/url-shortner/internal/app/service/logger"
	"github.com/Alandres998/url-shortner/internal/app/service/shortener"
	"github.com/gin-gonic/gin"
)

func DeleteShortURL(c *gin.Context) error {
	var shortURLs []string

	if err := c.BindJSON(&shortURLs); err != nil {
		logger.LogError("Shorter Delete", err.Error())
		return err
	}

	userID, err := auth.GetUserIDByCookie(c)
	if err != nil {
		logger.LogError("Shorter Delete", err.Error())
		return err
	}
	go shortener.DeleteShortURL(userID, shortURLs)
	return nil
}
