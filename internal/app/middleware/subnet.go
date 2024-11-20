package middlewares

import (
	"net"
	"net/http"

	"github.com/Alandres998/url-shortner/internal/config"
	"github.com/gin-gonic/gin"
)

// CheckTrustedSubnet Проверяет входит ли ИП в подсеть
func CheckTrustedSubnet() gin.HandlerFunc {
	return func(c *gin.Context) {
		trustedSubnet := config.Options.TrustedSubnet
		if trustedSubnet == "" {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		ip := c.GetHeader("X-Real-IP")
		if ip == "" {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		_, subnet, err := net.ParseCIDR(trustedSubnet)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		clientIP := net.ParseIP(ip)
		if clientIP == nil || !subnet.Contains(clientIP) {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Next()
	}
}
