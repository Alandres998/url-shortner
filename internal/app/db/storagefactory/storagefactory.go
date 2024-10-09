package storagefactory

import (
	"log"

	"github.com/Alandres998/url-shortner/internal/app/db/db"
	fileservices "github.com/Alandres998/url-shortner/internal/app/db/fileServices"
	"github.com/Alandres998/url-shortner/internal/app/db/storage"
	syncservices "github.com/Alandres998/url-shortner/internal/app/db/syncServices"
	"github.com/Alandres998/url-shortner/internal/config"
	"go.uber.org/zap"
)

// NewStorage фабрика создает стор в зависимости от конфига
func NewStorage() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()

	var store storage.Storage
	var message string
	if config.Options.DatabaseDSN != "" {
		store, err = db.NewDBStorage(config.Options.DatabaseDSN)
		message = "Выбрано хранилище с БД"
	} else if config.Options.FileStorage.Path != "" {
		store, err = fileservices.NewFileStorage(config.Options.FileStorage.Path)
		message = "Выбрано файловое хранилище"
	} else {
		store, err = syncservices.NewMemoryStorage(), nil
		message = "Выбрано хранилище - память"
	}
	logger.Info("Request",
		zap.String("store-server", message),
	)
	if err != nil {
		logger.Error("Не удалось иницировать хранилище",
			zap.String("Ошибка", err.Error()),
		)
		return
	}
	storage.Store = store
}
