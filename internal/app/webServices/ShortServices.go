package webservices

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/Alandres998/url-shortner/internal/app/db/storage"
	"github.com/Alandres998/url-shortner/internal/app/service/shortener"
	"github.com/Alandres998/url-shortner/internal/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Эти две структуры хотел бы вынести в модели но пофакту это структура запроса и ответа
// поэтому не стал выносить в сущность  models
type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	Result string `json:"result"`
}

type BatchRequest struct {
	CorrelationID string `json:"correlation_id"`
	OriginalURL   string `json:"original_url"`
}

type BatchResponse struct {
	CorrelationID string `json:"correlation_id"`
	ShortURL      string `json:"short_url"`
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

	err = storage.Store.Set(codeURL, originalURL)
	if err.Error() == storage.ErrURLExists.Error() {
		UrlStore, _ := storage.Store.GetbyOriginURL(originalURL)
		shortedCode = UrlStore.ShortURL
	}
	return shortedCode, err
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
	err = storage.Store.Set(codeURL, req.URL)
	if err.Error() == storage.ErrURLExists.Error() {
		UrlStore, _ := storage.Store.GetbyOriginURL(req.URL)
		shortedCode = UrlStore.ShortURL
	}
	return res, err
}

func ShorterJSONBatch(c *gin.Context) ([]BatchResponse, error) {
	var batchRequests []BatchRequest
	var batchResponses []BatchResponse
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("не смог иницировать логгер")
	}

	body, _ := io.ReadAll(c.Request.Body)

	err = json.Unmarshal(body, &batchRequests)
	if err != nil {
		return batchResponses, errors.New(Error400DefaultText)
	}

	for _, req := range batchRequests {
		codeURL := shortener.GenerateShortURL()
		shortedCode := fmt.Sprintf("%s/%s", config.Options.ServerAdress.ShortURL, codeURL)
		storage.Store.Set(codeURL, req.OriginalURL)
		if err != nil {
			logger.Error("запись в стор в баче",
				zap.String("ошибка", err.Error()),
			)
		}
		batchResponses = append(batchResponses, BatchResponse{
			CorrelationID: req.CorrelationID,
			ShortURL:      shortedCode,
		})
	}
	return batchResponses, nil
}
