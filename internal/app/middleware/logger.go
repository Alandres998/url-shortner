package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Logger() gin.HandlerFunc {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		// Логирование сведений о запросе и ответе
		endTime := time.Now()
		latency := endTime.Sub(startTime)

		statusCode := c.Writer.Status()
		contentLength := c.Writer.Size()

		logger.Info("Request",
			zap.String("url", c.Request.RequestURI),
			zap.String("method", c.Request.Method),
			zap.Duration("latency", latency),
			zap.Int("status_code", statusCode),
			zap.Int("content_length", contentLength),
		)
	}
}
