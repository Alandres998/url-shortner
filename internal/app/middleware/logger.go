package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Logger Логирует параметры запроса
func Logger() gin.HandlerFunc {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Не смог иницировать логгер")
	}

	defer func() {
		if errLoger := logger.Sync(); errLoger != nil {
			logger.Error("Проблемы при закрытии логера",
				zap.String("Не смог закрыть логгер", errLoger.Error()),
			)
		}
	}()

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
