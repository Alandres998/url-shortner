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
		auth.SetCookieUseInRequest(c)
		c.Next()
	}
}
