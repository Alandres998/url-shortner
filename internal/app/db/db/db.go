package db

import (
	"log"

	"github.com/Alandres998/url-shortner/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

var DB *sqlx.DB

func InitDB() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Не смог иницировать логгер")
	}
	defer logger.Sync()
	DB, err = sqlx.Connect("postgres", config.Options.DatabaseDSN)
	if err != nil {
		logger.Error("Проблемы при подключении к БД",
			zap.String("Не смог подключиться к БД", err.Error()),
		)
	}
}
