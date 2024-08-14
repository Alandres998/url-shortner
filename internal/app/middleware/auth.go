package middlewares

import (
	"github.com/Alandres998/url-shortner/internal/app/service/auth"
	"github.com/gin-gonic/gin"
)

// Mидлварка на установка куки
func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		cookie, err := c.Cookie(auth.CookieName)
		if err != nil || !auth.ValidateCookie(cookie) {
			userID := auth.GenerateUserID()
			auth.SetUserCookie(c, userID)
			c.Set(auth.CookieName, userID)
			c.SetCookie(auth.CookieName, userID, 3600, "/", "localhost", false, true)
		} else {
			c.Set(auth.CookieName, cookie)
			c.SetCookie(auth.CookieName, cookie, 3600, "/", "localhost", false, true)
		}

		c.Next()
	}
}
