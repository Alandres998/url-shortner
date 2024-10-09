package fileservices

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/Alandres998/url-shortner/internal/app/db/storage"
	"github.com/Alandres998/url-shortner/internal/config"
	"go.uber.org/zap"
)

// FileStorage представляет собой структуру для хранения URL-адресов в файле.
type FileStorage struct {
	filePath      string            // Путь к файлу хранения.
	mu            sync.RWMutex      // Мьютекс для безопасного доступа к данным.
	urlData       []storage.URLData // Срез данных URL.
	lastIncrement int               // Последний инкремент ID для новых записей.
}

// NewFileStorage создает новый экземпляр FileStorage и инициализирует файловое хранилище.
// Возвращает указатель на FileStorage и ошибку, если возникли проблемы.
func NewFileStorage(filePath string) (*FileStorage, error) {
	fs := &FileStorage{
		filePath: filePath,
	}
	err := fs.initFileStorage()
	if err != nil {
		return nil, err
	}
	return fs, nil
}

// initFileStorage инициализирует файловое хранилище, читая существующие данные или создавая новый файл.
// Возвращает ошибку, если инициализация не удалась.
func (fs *FileStorage) initFileStorage() error {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Не смог инициализировать логгер")
	}
	defer logger.Sync()
	fs.lastIncrement = 0
	urlSlice, err := fs.readOrCreateFile(fs.filePath)
	if err != nil {
		logger.Error("Инициализация стора",
			zap.String("Ошибка при инициализации", err.Error()),
		)
		log.Panic("Не смог проинициализировать файловое хранилище")
	}
	fs.urlData = urlSlice
	return nil
}

// readOrCreateFile читает данные из файла или создает новый файл, если он не существует.
// Возвращает срез URLData и ошибку, если возникли проблемы.
func (fs *FileStorage) readOrCreateFile(filePath string) ([]storage.URLData, error) {
	var items []storage.URLData
	fs.lastIncrement = 0
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть файл: %v", err)
	}
	defer file.Close()

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

// WriteInStorage записывает новый короткий URL в хранилище.
// Если хранилище работает в режиме только для чтения, ничего не делает.
func (fs *FileStorage) WriteInStorage(shortURL storage.URLData) {
	if config.Options.FileStorage.Mode == os.O_RDONLY {
		return
	}

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Не смог инициализировать логгер")
	}
	defer logger.Sync()

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
			zap.String("Ахтунг не преобразовал структуру в джсон", err.Error()),
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

// Set добавляет новый короткий URL в хранилище и сохраняет его в файл.
// Возвращает ошибку, если добавление не удалось.
func (fs *FileStorage) Set(ctx context.Context, userID, shortURL, originalURL string) error {
	newShortURL := storage.URLData{
		ID:          fs.lastIncrement + 1,
		ShortURL:    shortURL,
		OriginalURL: originalURL,
		UserID:      userID,
	}
	fs.urlData = append(fs.urlData, newShortURL)
	fs.WriteInStorage(newShortURL)
	return nil
}

// Get получает оригинальный URL по короткому URL.
// Возвращает оригинальный URL и ошибку, если URL не найден.
func (fs *FileStorage) Get(ctx context.Context, shortURL string) (string, error) {
	for _, data := range fs.urlData {
		if data.ShortURL == shortURL {
			return data.OriginalURL, nil
		}
	}
	return "", errors.New("такого адреса нет")
}

// GetbyOriginURL получает данные URL по оригинальному URL.
// Возвращает структуру URLData и ошибку, если оригинальный URL не найден.
func (fs *FileStorage) GetbyOriginURL(ctx context.Context, originalURL string) (storage.URLData, error) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()

	for _, urlData := range fs.urlData {
		if urlData.OriginalURL == originalURL {
			return urlData, nil
		}
	}
	return storage.URLData{}, errors.New("URL not found")
}

// GetUserURLs получает все короткие URL, принадлежащие определенному пользователю.
// Возвращает срез URLData и ошибку, если данные не найдены.
func (fs *FileStorage) GetUserURLs(ctx context.Context, userID string) ([]storage.URLData, error) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()

	var urls []storage.URLData
	for _, urlData := range fs.urlData {
		if urlData.UserID == userID {
			urls = append(urls, urlData)
		}
	}
	return urls, nil
}

// Ping проверяет доступность хранилища. В данном случае всегда возвращает nil.
func (fs *FileStorage) Ping(ctx context.Context) error {
	return nil
}

// DeleteUserURL удаляет короткие URL, принадлежащие авторизованному пользователю.
// Возвращает ошибку, если URL не найдены для удаления.
func (fs *FileStorage) DeleteUserURL(ctx context.Context, shortURLs []string, userID string) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	urlSet := make(map[string]struct{}, len(shortURLs))
	for _, url := range shortURLs {
		urlSet[url] = struct{}{}
	}

	updated := false
	for i, urlData := range fs.urlData {
		if _, found := urlSet[urlData.ShortURL]; found && urlData.UserID == userID {
			fs.urlData[i].Deleted = true
			updated = true
		}
	}

	if updated {
		fs.writeAllData()
		return nil
	}

	return errors.New("не удалось найти соответствующие URL для удаления")
}

// writeAllData записывает все данные из памяти обратно в файл, очищая его перед записью.
func (fs *FileStorage) writeAllData() {
	file, err := os.OpenFile(fs.filePath, os.O_TRUNC|os.O_RDWR, 0666)
	if err != nil {
		zap.L().Error("Ошибка при открытии файла", zap.Error(err))
		return
	}
	defer file.Close()

	for _, urlData := range fs.urlData {
		jsonData, err := json.Marshal(urlData)
		if err != nil {
			zap.L().Error("ошибка при серилизации данных", zap.Error(err))
			return
		}
		_, err = file.Write(jsonData)
		if err != nil {
			zap.L().Error("ошибка записи в файл", zap.Error(err))
			return
		}
		_, err = file.WriteString("\n")
		if err != nil {
			zap.L().Error("не смог записать сущность в файл", zap.Error(err))
			return
		}
	}
}
