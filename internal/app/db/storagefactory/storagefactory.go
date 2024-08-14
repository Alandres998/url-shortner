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

func NewStorage() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()

	var store storage.Storage
	if config.Options.DatabaseDSN != "" {
		store, err = db.NewDBStorage(config.Options.DatabaseDSN)
	} else if config.Options.FileStorage.Path != "" {
		store, err = fileservices.NewFileStorage(config.Options.FileStorage.Path)
	} else {
		store, err = syncservices.NewMemoryStorage(), nil
	}

	if err != nil {
		logger.Error("Не удалось иницировать хранилище",
			zap.String("Ошибка", err.Error()),
		)
		return
	}
	storage.Store = store
}
