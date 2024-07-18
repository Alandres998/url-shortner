package webservices

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	fileservices "github.com/Alandres998/url-shortner/internal/app/db/fileServices"
	"github.com/Alandres998/url-shortner/internal/app/service/shortener"
	"github.com/Alandres998/url-shortner/internal/config"
	"github.com/gin-gonic/gin"
)

// Эти две структуры хотел бы вынести в модели но пофакту это структура запроса и ответа
// поэтому не стал выносить в сущность  models
type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	Result string `json:"result"`
}

const Error400DefaultText = "Ошибка"

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

	codeURL := shortener.GenerateShortURL()
	shortedCode := fmt.Sprintf("%s/%s", config.Options.ServerAdress.ShortURL, codeURL)
	originalURL := string(body)

	fileservices.SaveUrl(codeURL, originalURL)
	return shortedCode, nil
}

func ShorterJSON(c *gin.Context) (ShortenResponse, error) {
	req := new(ShortenRequest)
	body, _ := io.ReadAll(c.Request.Body)

	err := json.Unmarshal(body, req)
	if err != nil {
		return ShortenResponse{}, errors.New(Error400DefaultText)
	}

	codeURL := shortener.GenerateShortURL()
	shortedCode := fmt.Sprintf("%s/%s", config.Options.ServerAdress.ShortURL, codeURL)
	res := ShortenResponse{Result: shortedCode}
	fileservices.SaveUrl(codeURL, req.URL)
	return res, nil
}
