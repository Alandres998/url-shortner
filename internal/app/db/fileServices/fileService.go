package fileservices

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Alandres998/url-shortner/internal/config"
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

	// Проверяем существование файла
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		if err := os.WriteFile(filePath, []byte(""), 0644); err != nil {
			return nil, fmt.Errorf("файла не существует yt смог его создать: %v", err)
		}
	}

	// Считываем данные из файла
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("не смог открыть файл на чтение: %v", err)
	}

	// Разбираем JSON данные в структуры(объекты)
	lines := bytes.Split(fileData, []byte("\n"))
	for _, line := range lines {
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

func SaveUrl(shortUrl string, originalUrl string) URLData {
	newShortUrl := URLData{
		UUID:        strconv.Itoa(lastIncrement + 1),
		ShortURL:    shortUrl,
		OriginalURL: originalUrl,
	}
	urlData = append(urlData, newShortUrl)
	WriteInStorage(newShortUrl)
	return newShortUrl
}

func WriteInStorage(shortURL URLData) {
	if config.Options.FileStorage.Mode == os.O_RDONLY {
		return
	}

	// Открываем файл для записи в конец
	file, err := os.OpenFile(config.Options.FileStorage.Path, os.O_APPEND|config.Options.FileStorage.Mode, 0644)
	if err != nil {
		fmt.Printf("Ошибка не смог открыть файл для записи: %v\n", err)
		return
	}
	defer file.Close()

	jsonData, err := json.Marshal(shortURL)
	if err != nil {
		fmt.Printf("Ахтунг не преобразовал структуту в джсон %v\n", err)
		return
	}
	jsonData = append(jsonData, '\n')
	if _, err := file.Write(jsonData); err != nil {
		fmt.Printf("Не смог записать в файл структуру: %v\n", err)
	}
}
