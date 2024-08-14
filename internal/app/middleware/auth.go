package middlewares

import (
	"github.com/Alandres998/url-shortner/internal/app/service/auth"
	"github.com/gin-gonic/gin"
)

// Mидлварка на установка куки
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth.SetCookieUseInRequest(c)
		c.Next()
	}
}
