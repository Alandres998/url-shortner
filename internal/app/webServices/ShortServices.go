package webservices

import (
	"errors"
	"fmt"
	"io"

	shorterservices "github.com/Alandres998/url-shortner/internal/app/buslogic/shorterServices"
	syncservices "github.com/Alandres998/url-shortner/internal/app/db/syncServices"
	"github.com/gin-gonic/gin"
)

var Error400DefaultText = "Ошибка"
var MainURL = "http://localhost:8080"

func GetErrorWithCode(c *gin.Context, errorText string, codeError int) {
	c.Writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
	c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
	c.Writer.WriteHeader(codeError)
	fmt.Fprintln(c.Writer, errorText)
}

func Shorter(c *gin.Context) (string, error) {
	req := c.Request

	//Проверка на метод и тело содержимого
	body, err := io.ReadAll(req.Body)
	if err != nil || len(body) == 0 {
		return "", errors.New(Error400DefaultText)
	}

	codeURL := shorterservices.GenerateShortURL()
	shortedCode := fmt.Sprintf("%s/%s", MainURL, codeURL)
	originalURL := string(body)

	syncservices.UrlStorage.Set(codeURL, originalURL)
	return shortedCode, nil
}
