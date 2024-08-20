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

const Error400DefaultText = "Ошибка"

// Определение типов для запросов и ответов
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

func GetErrorWithCode(c *gin.Context, errorText string, codeError int) {
	c.Writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
	c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
	c.Writer.WriteHeader(codeError)
	fmt.Fprintln(c.Writer, errorText)
}

func Shorter(c *gin.Context) (string, error) {
	ctx := context.Background()
	req := c.Request

	logger, errLog := zap.NewProduction()
	if errLog != nil {
		log.Fatalf("Не смог иницировать логгер")
	}

	defer logger.Sync()

	userID, err := auth.GetUserID(c)
	if err != nil {
		logger.Info("Shorter Save",
			zap.String("Внимание", err.Error()),
		)
	}

	body, err := io.ReadAll(req.Body)
	if err != nil || len(body) == 0 {
		return "", errors.New(Error400DefaultText)
	}

	codeURL := shortener.GenerateShortURL()
	shortedCode := fmt.Sprintf("%s/%s", config.Options.ServerAdress.ShortURL, codeURL)
	originalURL := string(body)

	err = storage.Store.Set(ctx, userID, codeURL, originalURL)
	if err != nil {
		if errors.Is(err, storage.ErrURLExists) {
			URLStore, err := storage.Store.GetbyOriginURL(ctx, originalURL)
			URLStore.ShortURL = fmt.Sprintf("%s/%s", config.Options.ServerAdress.ShortURL, URLStore.ShortURL)
			shortedCode = URLStore.ShortURL
			if err != nil {
				logger.Error("Shorter Save Dublicate",
					zap.String("Ошибка", string(err.Error())),
				)
			}
		} else {
			logger.Error("Shorter Save",
				zap.String("Ошибка", string(err.Error())),
			)
		}
	}

	return shortedCode, err
}

func ShorterJSON(c *gin.Context) (ShortenResponse, error) {
	ctx := context.Background()
	req := new(ShortenRequest)
	logger, errLog := zap.NewProduction()
	if errLog != nil {
		log.Fatalf("Не смог иницировать логгер")
	}

	defer logger.Sync()

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

func ShorterJSONBatch(c *gin.Context) ([]BatchResponse, error) {
	ctx := context.Background()
	var batchRequests []BatchRequest
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("не смог иницировать логгер")
	}
	defer logger.Sync()

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
