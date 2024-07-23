package fileservices

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Alandres998/url-shortner/internal/config"
	"go.uber.org/zap"
)

// URLData представляет данные о URL для сохранения в файл
type URLData struct {
	UUID        string `json:"uuid"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

var (
	urlData       []URLData
	lastIncrement int
)

func InitFileStorage() {

	lastIncrement = 0
	urlSlice, err := readOrCreateFile(config.Options.FileStorage.Path)
	if err != nil {
		log.Panic("Не смог проиницировать файловое хранилище")
	}
	urlData = urlSlice
}

func readOrCreateFile(filePath string) ([]URLData, error) {
	var items []URLData
	lastIncrement := 0
	file, err := os.Open(filePath)
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
		var item URLData
		if err := json.Unmarshal(line, &item); err != nil {
			return nil, fmt.Errorf("не удалось распарсить файл: %v", err)
		}

		num, err := strconv.Atoi(item.UUID)
		if err != nil {
			return nil, fmt.Errorf("ошибка при преобразовании строки в число: %v", err)
		}
		if num > lastIncrement {
			lastIncrement = num
		}
		items = append(items, item)
	}

	return items, nil
}

func GetURL(shortURL string) *URLData {
	for _, data := range urlData {
		if data.ShortURL == shortURL {
			return &data
		}
	}
	return nil
}

func SaveURL(ShortURL string, OriginalURL string) URLData {
	newShortURL := URLData{
		UUID:        strconv.Itoa(lastIncrement + 1),
		ShortURL:    ShortURL,
		OriginalURL: OriginalURL,
	}
	urlData = append(urlData, newShortURL)
	WriteInStorage(newShortURL)
	return newShortURL
}

func WriteInStorage(shortURL URLData) {
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
