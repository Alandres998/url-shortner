package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

// OptionsStruct структура с настройками
type OptionsStruct struct {
	ServerAdress ServerConfig
	FileStorage  FileStorageConfig
	DatabaseDSN  string `json:"database_dsn"`
	StorageType  string `json:"storage_type"`
	EnableHTTPS  bool   `json:"enable_https"`
	SSLConfig    SSLConfig
}

// Options общая конфигурация проекта
var Options OptionsStruct

// SSLConfig информация о том где искать сертификат
type SSLConfig struct {
	CertFile string `json:"ssl_cert_file"`
	KeyFile  string `json:"ssl_key_file"`
}

// ServerConfig конфигурация сервера
type ServerConfig struct {
	MainURLServer string `json:"server_address"`
	ShortURL      string `json:"base_url"`
}

// FileStorageConfig конфигурация файлового хранилища
type FileStorageConfig struct {
	Path string `json:"file_storage_path"`
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
	loadConfigJson()
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
	flag.BoolVar(&Options.EnableHTTPS, "s", false, "enable HTTPS")
	flag.StringVar(&Options.SSLConfig.CertFile, "cert", "server.crt", "path to SSL certificate") //Чет в задании не было про подсовывание ключей для http.ListenAndServeTLS
	flag.StringVar(&Options.SSLConfig.KeyFile, "key", "server.key", "path to SSL key")
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

	if envEnableHTTPS := os.Getenv("ENABLE_HTTPS"); envEnableHTTPS == "true" {
		Options.EnableHTTPS = true
	}

	if envCertFile := os.Getenv("SSL_CERT_FILE"); envCertFile != "" {
		Options.SSLConfig.CertFile = envCertFile
	}
	if envKeyFile := os.Getenv("SSL_KEY_FILE"); envKeyFile != "" {
		Options.SSLConfig.KeyFile = envKeyFile
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

// loadJson Чтение конфига из json
func loadConfigJson() {
	configFilePath := flag.String("c", "", "config file path")
	flag.StringVar(configFilePath, "config", "", "config file path")

	if envConfigPath := os.Getenv("CONFIG"); envConfigPath != "" && *configFilePath == "" {
		*configFilePath = envConfigPath
	}

	if *configFilePath != "" {
		var configFromFile = OptionsStruct{}
		var serverAdress = ServerConfig{}
		var fileStorage = FileStorageConfig{}
		var sslConfig = SSLConfig{}
		file, err := os.ReadFile(*configFilePath)
		if err != nil {
			log.Fatalf("Ошибка чтения конфигурационного файла: %v", err)
		}

		err = json.Unmarshal(file, &configFromFile)
		if err != nil {
			log.Fatalf("Ошибка парсинга конфигурационного файла: %v", err)
		}

		err = json.Unmarshal(file, &serverAdress)
		if err != nil {
			log.Fatalf("Ошибка парсинга конфигурационного файла сервера: %v", err)
		}

		err = json.Unmarshal(file, &fileStorage)
		if err != nil {
			log.Fatalf("Ошибка парсинга конфигурационного файла хранилища: %v", err)
		}

		err = json.Unmarshal(file, &sslConfig)
		if err != nil {
			log.Fatalf("Ошибка парсинга конфигурационного файла сертификата: %v", err)
		}
		configFromFile.ServerAdress = serverAdress
		configFromFile.FileStorage = fileStorage
		configFromFile.SSLConfig = sslConfig

		if Options.ServerAdress.MainURLServer == "" {
			Options.ServerAdress.MainURLServer = configFromFile.ServerAdress.MainURLServer
		}
		if Options.ServerAdress.ShortURL == "" {
			Options.ServerAdress.ShortURL = configFromFile.ServerAdress.ShortURL
		}
		if Options.FileStorage.Path == "" {
			Options.FileStorage.Path = configFromFile.FileStorage.Path
		}
		if Options.FileStorage.Mode != 0 {
			Options.FileStorage.Mode = configFromFile.FileStorage.Mode
		}
		if Options.DatabaseDSN == "" {
			Options.DatabaseDSN = configFromFile.DatabaseDSN
		}
		if !Options.EnableHTTPS {
			Options.EnableHTTPS = configFromFile.EnableHTTPS
		}
		if Options.SSLConfig.CertFile == "" {
			Options.SSLConfig.CertFile = configFromFile.SSLConfig.CertFile
		}
		if Options.SSLConfig.KeyFile == "" {
			Options.SSLConfig.KeyFile = configFromFile.SSLConfig.KeyFile
		}
	}
}
