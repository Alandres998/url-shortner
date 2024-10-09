package config

import (
	"flag"
	"fmt"
	"os"
)

// Options общая конфигурация проекта
var Options struct {
	ServerAdress ServerConfig
	FileStorage  FileStorageConfig
	DatabaseDSN  string
	StorageType  string
}

// ServerConfig конфигурация сервера
type ServerConfig struct {
	MainURLServer string
	ShortURL      string
}

// FileStorageConfig конфигурация файлового хранилища
type FileStorageConfig struct {
	Path string
	Mode int
}

// StorageTypeDB константа для факторки БД
const StorageTypeDB = "database"

// StorageTypeFile константа для факторки файлового хранилища
const StorageTypeFile = "file"

// StorageTypeFile константа для факторки хранилища в памяти
const StorageTypeMemory = "memory"

// InitConfig инициализация конфига
func InitConfig() {
	if os.Getenv("RUN_MODE") == "test" {
		return
	}
	parseFlags()
	loadEnv()
	loadConfigFile()
	determineStorageType()
}

// InitConfigExample инициализация конфига для теста
func InitConfigExample() {
	if os.Getenv("RUN_MODE") == "test" {
		return
	}
	loadEnv()
	loadConfigFile()
	Options.StorageType = StorageTypeMemory
}

// GetAdressServer функция получения адреса сервера
func GetAdressServer(Port string) string {
	text := fmt.Sprintf("http://localhost%s", Port)
	return text
}

// parseFlags Устанавливаем конфиг из флагов командой строки
func parseFlags() {
	flag.StringVar(&Options.ServerAdress.MainURLServer, "a", ":8080", "basic main address")
	flag.StringVar(&Options.ServerAdress.ShortURL, "b", "http://localhost:8080", "short response address")
	flag.StringVar(&Options.FileStorage.Path, "f", "", "storage file")
	flag.StringVar(&Options.DatabaseDSN, "d", "", "database DSN")
	flag.Parse()
}

// loadEnv Устанавливаем конфиг из env
func loadEnv() {
	if envMainURLServer := os.Getenv("SERVER_ADDRESS"); envMainURLServer != "" {
		Options.ServerAdress.MainURLServer = envMainURLServer
	}
	if envShortURL := os.Getenv("BASE_URL"); envShortURL != "" {
		Options.ServerAdress.ShortURL = envShortURL
	}
	if envFileStorage := os.Getenv("FILE_STORAGE_PATH"); envFileStorage != "" {
		Options.FileStorage.Path = envFileStorage
	}

	if envDatabaseDSN := os.Getenv("DATABASE_DSN"); envDatabaseDSN != "" {
		Options.DatabaseDSN = envDatabaseDSN
	}
}

// loadConfigFile Устанавливаем конфиг для файла хранилища
func loadConfigFile() {
	if flag.Lookup("f").Value.String() == "" && os.Getenv("FILE_STORAGE_PATH") == "" {
		Options.FileStorage.Path = "/tmp/short-url-db.json"
		Options.FileStorage.Mode = os.O_RDONLY
	} else {
		Options.FileStorage.Mode = os.O_RDWR
	}
}

// determineStorageType Авто выбор хранилища на основе конфига
func determineStorageType() {
	if Options.DatabaseDSN != "" {
		Options.StorageType = StorageTypeDB
	} else if Options.FileStorage.Path != "" {
		Options.StorageType = StorageTypeFile
	} else {
		Options.StorageType = StorageTypeMemory
	}
}
