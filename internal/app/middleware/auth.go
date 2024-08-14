package middlewares

import (
	"log"

	"github.com/Alandres998/url-shortner/internal/app/service/auth"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Mидлварка на установка куки
func AuthMiddleware() gin.HandlerFunc {
	logger, errLog := zap.NewProduction()
	defer logger.Sync()
	if errLog != nil {
		log.Fatalf("Не смог иницировать логгер")
	}
	return func(c *gin.Context) {
		cookie, err := c.Cookie(auth.CookieName)
		if err != nil || !auth.ValidateCookie(cookie) {
			userID := auth.GenerateUserID()
			auth.SetUserCookie(c, userID)
			c.Set(auth.CookieName, userID)
		} else {
			c.Set(auth.CookieName, cookie)
		}
		c.Next()
	}
}
