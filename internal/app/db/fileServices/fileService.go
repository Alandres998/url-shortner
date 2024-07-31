package fileservices

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/Alandres998/url-shortner/internal/app/db/storage"
	"github.com/Alandres998/url-shortner/internal/config"
	"go.uber.org/zap"
)

type FileStorage struct {
	filePath      string
	urlData       []storage.URLData
	lastIncrement int
}

func NewFileStorage(filePath string) (storage.Storage, error) {
	fs := &FileStorage{
		filePath: filePath,
	}
	err := fs.initFileStorage()
	if err != nil {
		return nil, err
	}
	return fs, nil
}

func (fs *FileStorage) initFileStorage() error {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Не смог иницировать логгер")
	}
	defer logger.Sync()
	fs.lastIncrement = 0
	urlSlice, err := fs.readOrCreateFile(fs.filePath)
	if err != nil {
		logger.Error("Инициализация стора",
			zap.String("Ошибка при инициализации", err.Error()),
		)
		log.Panic("Не смог проиницировать файловое хранилище")
	}
	fs.urlData = urlSlice
	return nil
}

func (fs *FileStorage) readOrCreateFile(filePath string) ([]storage.URLData, error) {
	var items []storage.URLData
	fs.lastIncrement = 0
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть файл: %v", err)
	}
	defer file.Close()

	// Разбираем JSON данные в структуры(объекты)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}
		var item storage.URLData
		if err := json.Unmarshal(line, &item); err != nil {
			return nil, fmt.Errorf("не удалось распарсить файл: %v", err)
		}

		if item.ID > fs.lastIncrement {
			fs.lastIncrement = item.ID
		}
		items = append(items, item)
	}

	return items, nil
}

func (fs *FileStorage) Set(shortURL, originalURL string) error {
	newShortURL := storage.URLData{
		ID:          fs.lastIncrement + 1,
		ShortURL:    shortURL,
		OriginalURL: originalURL,
	}
	fs.urlData = append(fs.urlData, newShortURL)
	fs.WriteInStorage(newShortURL)
	return nil
}

func (fs *FileStorage) Get(shortURL string) (string, error) {
	for _, data := range fs.urlData {
		if data.ShortURL == shortURL {
			return data.OriginalURL, nil
		}
	}
	return "", errors.New("URL not found")
}

func (fs *FileStorage) WriteInStorage(shortURL storage.URLData) {
	if config.Options.FileStorage.Mode == os.O_RDONLY {
		return
	}

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Не смог иницировать логгер")
	}
	defer logger.Sync()
	// Открываем файл для записи в конец
	file, err := os.OpenFile(config.Options.FileStorage.Path, os.O_APPEND|config.Options.FileStorage.Mode, 0644)
	if err != nil {
		logger.Error("Запись в файл store",
			zap.String("ошибка", err.Error()),
		)
		return
	}
	defer file.Close()

	jsonData, err := json.Marshal(shortURL)
	if err != nil {
		logger.Error("Запись в файл store",
			zap.String("Ахтунг не преобразовал структуту в джсон", err.Error()),
		)
		return
	}
	jsonData = append(jsonData, '\n')
	if _, err := file.Write(jsonData); err != nil {
		logger.Error("Запись в файл store",
			zap.String("Не смог записать в файл структуру", err.Error()),
		)
	}
}
