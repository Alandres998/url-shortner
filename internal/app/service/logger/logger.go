package logger

import (
	"log"

	"go.uber.org/zap"
)

// LoginInfo записывает информационное сообщение в лог.
// title - заголовок сообщения, info - сообщение
func LoginInfo(title string, info string) {
	logger, errLog := zap.NewProduction()
	defer logger.Sync()
	if errLog != nil {
		log.Fatalf("Не смог инициализировать логгер")
	}

	logger.Info("Внимание",
		zap.String(title, info),
	)
}

// LogError записывает сообщение об ошибке в лог.
// title - заголовок сообщения об ошибке, info - сообщение
func LogError(title string, info string) {
	logger, errLog := zap.NewProduction()
	defer logger.Sync()
	if errLog != nil {
		log.Fatalf("Не смог инициализировать логгер")
	}

	logger.Error("Ошибка",
		zap.String(title, info),
	)
}
