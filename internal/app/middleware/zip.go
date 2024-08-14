package middlewares

import (
	"bytes"
	"compress/gzip"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/Alandres998/url-shortner/internal/app/service/auth"
	"github.com/gin-gonic/gin"
)

func GzipMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth.InfoCookie(c, "Действия до мидлварке Zip")
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
			auth.InfoCookie(c, "Действия сжатия Zip")
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
			auth.InfoCookie(c, "Действия без сжатия")
			_, err := c.Writer.Write(buffer.Bytes())
			if err != nil {
				c.String(http.StatusInternalServerError, "Не смог записать в ответ")
			}
		}
		auth.InfoCookie(c, "Действия до после Zip middleware")
	}
}

// Проверяем можем ли зиповать такой тип ответа/запроса
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

// Проверяем можем ли зиповать такой ответ с таким
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

type responseWriter struct {
	gin.ResponseWriter
	buffer *bytes.Buffer
}

func (w *responseWriter) Write(data []byte) (int, error) {
	w.buffer.Write(data)
	if w.Header().Get("Content-Encoding") == "gzip" {
		return len(data), nil
	}
	return w.ResponseWriter.Write(data)
}
