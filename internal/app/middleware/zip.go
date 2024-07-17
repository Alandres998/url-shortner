package middlewares

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GzipMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.Contains(c.GetHeader("Content-Encoding"), "gzip") {
			reader, err := gzip.NewReader(c.Request.Body)
			if err != nil {
				c.String(http.StatusBadRequest, "Контент не заархивирован")
				c.Abort()
				return
			}
			defer reader.Close()
			c.Request.Body = io.NopCloser(reader)
		}

		c.Next()

		if strings.Contains(c.GetHeader("Accept-Encoding"), "gzip") {
			contentType := c.Writer.Header().Get("Content-Type")
			if shouldCompress(contentType) {
				c.Writer.Header().Set("Content-Encoding", "gzip")
				gz := gzip.NewWriter(c.Writer)
				defer gz.Close()
				c.Writer = &gzipWriter{Writer: gz, ResponseWriter: c.Writer}
			}
		}
	}
}

// Проверяем можем ли зиповать такой тип ответа/запроса
func shouldCompress(contentType string) bool {
	compressibleTypes := []string{
		"application/json",
		"text/plain",
	}
	for _, t := range compressibleTypes {
		if strings.Contains(contentType, t) {
			return true
		}
	}
	return false
}

type gzipWriter struct {
	gin.ResponseWriter
	Writer io.Writer
}

func (w *gzipWriter) Write(data []byte) (int, error) {
	return w.Writer.Write(data)
}
