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

	var storege storage.Storage

	if config.Options.DatabaseDSN != "" {
		storege, err = db.NewDBStorage(config.Options.DatabaseDSN)
	} else if config.Options.FileStorage.Path != "" {
		storege, err = fileservices.NewFileStorage(config.Options.FileStorage.Path)
	} else {
		storege, err = syncservices.NewMemoryStorage(), nil
	}

	if err != nil {
		logger.Error("Не удалось иницировать хранилище",
			zap.String("Ошибка", err.Error()),
		)
		return
	}
	storage.Store = storege
}
