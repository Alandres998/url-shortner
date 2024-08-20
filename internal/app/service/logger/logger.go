package logger

import (
	"log"

	"go.uber.org/zap"
)

func LoginInfo(title string, info string) {
	logger, errLog := zap.NewProduction()
	defer logger.Sync()
	if errLog != nil {
		log.Fatalf("Не смог иницировать логгер")
	}

	logger.Info("Внимание",
		zap.String(title, info),
	)
}

func LogError(title string, info string) {
	logger, errLog := zap.NewProduction()
	defer logger.Sync()
	if errLog != nil {
		log.Fatalf("Не смог иницировать логгер")
	}

	logger.Error("Ошибка",
		zap.String(title, info),
	)
}
