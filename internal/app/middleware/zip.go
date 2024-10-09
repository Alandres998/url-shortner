package middlewares

import (
	"bytes"
	"compress/gzip"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// GzipMiddleware Сжимает ответ, если это поддерживает клиент
func GzipMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.Contains(c.GetHeader("Content-Encoding"), "gzip") {
			reader, err := gzip.NewReader(c.Request.Body)
			if err != nil {
				c.AbortWithError(http.StatusBadRequest, errors.New("контент не заархивирован"))
				return
			}
			defer reader.Close()
			c.Request.Body = io.NopCloser(reader)
		}

		buffer := new(bytes.Buffer)
		writer := &responseWriter{ResponseWriter: c.Writer, buffer: buffer}
		c.Writer = writer

		c.Next()

		if strings.Contains(c.GetHeader("Accept-Encoding"), "gzip") {
			contentType := c.Writer.Header().Get("Content-Type")
			if shouldCompressContent(contentType) && avaibleCompressCode(c.Writer.Status()) {
				c.Writer.Header().Set("Content-Encoding", "gzip")
				c.Writer.Header().Del("Content-Length")
				gz := gzip.NewWriter(c.Writer)
				defer gz.Close()

				_, err := gz.Write(buffer.Bytes())
				if err != nil {
					c.AbortWithError(http.StatusBadRequest, errors.New("не смог записать в ответ"))
					return
				}
				return
			}
		}

		if !strings.Contains(c.GetHeader("Accept-Encoding"), "identity") && !strings.Contains(c.GetHeader("Accept-Encoding"), "") {

			_, err := c.Writer.Write(buffer.Bytes())
			if err != nil {
				c.String(http.StatusInternalServerError, "Не смог записать в ответ")
			}
		}
	}
}

// shouldCompressContent Проверяем можем ли зиповать такой тип ответа/запроса
func shouldCompressContent(contentType string) bool {
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

// avaibleCompressCode Проверяем можем ли зиповать такой ответ с таким
func avaibleCompressCode(CodeResponse int) bool {
	notAvaibleTypeCode := []int{
		http.StatusTemporaryRedirect,
	}
	for _, t := range notAvaibleTypeCode {
		if CodeResponse == t {
			return false
		}
	}
	return true
}

// responseWriter структура для расширения  ответа gin
type responseWriter struct {
	gin.ResponseWriter
	buffer *bytes.Buffer
}

// Write записать данные в ответ
func (w *responseWriter) Write(data []byte) (int, error) {
	w.buffer.Write(data)
	if w.Header().Get("Content-Encoding") == "gzip" {
		return len(data), nil
	}
	return w.ResponseWriter.Write(data)
}
