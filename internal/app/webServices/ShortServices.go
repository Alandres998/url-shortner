package webservices

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/Alandres998/url-shortner/internal/app/db/storage"
	"github.com/Alandres998/url-shortner/internal/app/service/auth"
	"github.com/Alandres998/url-shortner/internal/app/service/shortener"
	"github.com/Alandres998/url-shortner/internal/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Error400DefaultText Текст обычной ошибки
const Error400DefaultText = "Ошибка"

// ShortenRequest Определение типов для запросов и ответов
type ShortenRequest struct {
	URL string `json:"url"`
}

// ShortenRequest структура ответа на сокращение
type ShortenResponse struct {
	Result string `json:"result"`
}

// BatchRequest структура параметров запроса на батч сокращения ссылок json
type BatchRequest struct {
	CorrelationID string `json:"correlation_id"`
	OriginalURL   string `json:"original_url"`
}

// BatchRequest структура параметров ответа на батч сокращения ссылок json
type BatchResponse struct {
	CorrelationID string `json:"correlation_id"`
	ShortURL      string `json:"short_url"`
}

// GetErrorWithCode вызов кастомноый ошибки с нужным кодом и текстом
func GetErrorWithCode(c *gin.Context, errorText string, codeError int) {
	c.Writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
	c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
	c.Writer.WriteHeader(codeError)
	if _, err := fmt.Fprintln(c.Writer, errorText); err != nil {
		log.Printf("Ошибка при записи в хендлер ответа: %v", err)
	}
}

// Shorter Веб-Сервис сокращения ссылки присланной text/plain
func Shorter(c *gin.Context) (string, error) {
	ctx := c.Request.Context()
	req := c.Request

	userID, err := auth.GetUserID(c)
	if err != nil {
		logger, errLog := zap.NewProduction()
		if errLog != nil {
			log.Fatalf("Не смог иницировать логгер")
		}

		defer func() {
			if errLoger := logger.Sync(); errLoger != nil {
				logger.Error("Проблемы при закрытии логера",
					zap.String("Не смог закрыть логгер", errLoger.Error()),
				)
			}
		}()

		logger.Info("Shorter Save", zap.String("Внимание", err.Error()))
	}

	const maxBodySize = 1024 // 1MB
	body, err := io.ReadAll(io.LimitReader(req.Body, maxBodySize))
	if err != nil || len(body) == 0 {
		return "", errors.New(Error400DefaultText)
	}

	codeURL := shortener.GenerateShortURL()
	shortedCode := fmt.Sprintf("%s/%s", config.Options.ServerAdress.ShortURL, codeURL)
	originalURL := string(body)

	err = storage.Store.Set(ctx, userID, codeURL, originalURL)
	if err != nil {
		logger, errLog := zap.NewProduction()
		if errLog != nil {
			log.Fatalf("Не смог иницировать логгер")
		}

		defer func() {
			if errLoger := logger.Sync(); errLoger != nil {
				logger.Error("Проблемы при закрытии логера",
					zap.String("Не смог закрыть логгер", errLoger.Error()),
				)
			}
		}()

		if errors.Is(err, storage.ErrURLExists) {
			URLStore, errDB := storage.Store.GetbyOriginURL(ctx, originalURL)
			if errDB == nil {
				URLStore.ShortURL = fmt.Sprintf("%s/%s", config.Options.ServerAdress.ShortURL, URLStore.ShortURL)
				shortedCode = URLStore.ShortURL
			} else {
				logger.Error("Shorter Save Duplicate", zap.Error(errDB))
			}
		} else {
			logger.Error("Shorter Save", zap.Error(err))
		}
	}

	return shortedCode, err
}

func ShorterGeneral(ctx context.Context, userID string, originalURL string) (string, error) {
	codeURL := shortener.GenerateShortURL()
	shortedCode := fmt.Sprintf("%s/%s", config.Options.ServerAdress.ShortURL, codeURL)

	err := storage.Store.Set(ctx, userID, codeURL, originalURL)
	if err != nil {
		if errors.Is(err, storage.ErrURLExists) {
			URLStore, errDB := storage.Store.GetbyOriginURL(ctx, originalURL)
			if errDB == nil {
				URLStore.ShortURL = fmt.Sprintf("%s/%s", config.Options.ServerAdress.ShortURL, URLStore.ShortURL)
				shortedCode = URLStore.ShortURL
			} else {
				log.Printf("Error getting URL from storage: %v", errDB)
			}
		} else {
			log.Printf("Error saving shortened URL: %v", err)
		}
	}

	return shortedCode, err
}

// ShorterJSON Веб-Сервис сокращения ссылок присланных application/json
func ShorterJSON(c *gin.Context) (ShortenResponse, error) {
	ctx := context.Background()
	req := new(ShortenRequest)
	logger, errLog := zap.NewProduction()
	if errLog != nil {
		log.Fatalf("Не смог иницировать логгер")
	}

	defer func() {
		if errLoger := logger.Sync(); errLoger != nil {
			logger.Error("Проблемы при закрытии логера",
				zap.String("Не смог закрыть логгер", errLoger.Error()),
			)
		}
	}()

	body, _ := io.ReadAll(c.Request.Body)

	userID, err := auth.GetUserID(c)
	if err != nil {
		logger.Info("ShorterJson Save",
			zap.String("Внимание", err.Error()),
		)
	}

	err = json.Unmarshal(body, req)
	if err != nil {
		return ShortenResponse{}, errors.New(Error400DefaultText)
	}

	codeURL := shortener.GenerateShortURL()
	shortedCode := fmt.Sprintf("%s/%s", config.Options.ServerAdress.ShortURL, codeURL)
	res := ShortenResponse{Result: shortedCode}
	err = storage.Store.Set(ctx, userID, codeURL, req.URL)

	if err != nil {
		if errors.Is(err, storage.ErrURLExists) {
			URLStore, _ := storage.Store.GetbyOriginURL(ctx, req.URL)
			URLStore.ShortURL = fmt.Sprintf("%s/%s", config.Options.ServerAdress.ShortURL, URLStore.ShortURL)
			res.Result = URLStore.ShortURL
		} else {
			logger.Error("ShortJSON Save",
				zap.String("Ошибка", string(err.Error())),
			)
			return ShortenResponse{}, err
		}
	}

	logger.Info("Request",
		zap.String("body-Response", string(res.Result)),
		zap.String("body-Request", string(body)),
	)

	return res, err
}

// ShorterJSONBatch Веб-Сервис сокращения батча ссылок присланных application/json
func ShorterJSONBatch(c *gin.Context) ([]BatchResponse, error) {
	ctx := context.Background()
	var batchRequests []BatchRequest
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("не смог иницировать логгер")
	}

	defer func() {
		if errLoger := logger.Sync(); errLoger != nil {
			logger.Error("Проблемы при закрытии логера",
				zap.String("Не смог закрыть логгер", errLoger.Error()),
			)
		}
	}()

	userID, err := auth.GetUserID(c)
	if err != nil {
		logger.Info("ShorterJsonBatch Save",
			zap.String("Внимание", err.Error()),
		)
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return nil, errors.New(Error400DefaultText)
	}

	err = json.Unmarshal(body, &batchRequests)
	if err != nil {
		return nil, errors.New(Error400DefaultText)
	}

	batchResponses := make([]BatchResponse, 0, len(batchRequests))

	for _, req := range batchRequests {
		codeURL := shortener.GenerateShortURL()
		shortedCode := fmt.Sprintf("%s/%s", config.Options.ServerAdress.ShortURL, codeURL)
		err := storage.Store.Set(ctx, userID, codeURL, req.OriginalURL)
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
