package webservices

import (
	"log"

	"github.com/Alandres998/url-shortner/internal/app/db/storage"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// StatisticResponse Структура ответа
type StatisticResponse struct {
	TotalUrls  int `json:"total_urls"`
	TotalUsers int `json:"total_users"`
}

// GetStatisticsShortURL Возвращает статистику по shortURL для веб запроса
func GetStatisticsShortURL(c *gin.Context) (StatisticResponse, error) {
	ctx := c.Request.Context()

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

	var StatisticResponse StatisticResponse
	urlCount, userCount, err := storage.Store.GetStatistics(ctx)
	if err != nil {
		return StatisticResponse, Error401DefaultText
	}

	StatisticResponse.TotalUrls = urlCount
	StatisticResponse.TotalUsers = userCount
	return StatisticResponse, nil
}
