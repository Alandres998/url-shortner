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
	loadConfigJSON()
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
	setOptionIfEmpty(&Options.ServerAdress.MainURLServer, os.Getenv("SERVER_ADDRESS"))
	setOptionIfEmpty(&Options.ServerAdress.ShortURL, os.Getenv("BASE_URL"))
	setOptionIfEmpty(&Options.FileStorage.Path, os.Getenv("FILE_STORAGE_PATH"))
	setOptionIfEmpty(&Options.DatabaseDSN, os.Getenv("DATABASE_DSN"))
	setOptionIfEmptyBool(&Options.EnableHTTPS, stringToBool(os.Getenv("ENABLE_HTTPS")))
	setOptionIfEmpty(&Options.SSLConfig.CertFile, os.Getenv("SSL_CERT_FILE"))
	setOptionIfEmpty(&Options.SSLConfig.KeyFile, os.Getenv("SSL_KEY_FILE"))
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

// loadConfigJSON Чтение и установка конфига из json
func loadConfigJSON() {
	configFilePath := flag.String("c", "", "config file path")
	flag.StringVar(configFilePath, "config", "", "config file path")

	if envConfigPath := os.Getenv("CONFIG"); envConfigPath != "" && *configFilePath == "" {
		*configFilePath = envConfigPath
	}

	if *configFilePath == "" {
		return
	}

	configFromFile := parseFileConfigJSON(configFilePath)

	setOptionIfEmpty(&Options.ServerAdress.MainURLServer, configFromFile.ServerAdress.MainURLServer)
	setOptionIfEmpty(&Options.ServerAdress.ShortURL, configFromFile.ServerAdress.ShortURL)
	setOptionIfEmpty(&Options.FileStorage.Path, configFromFile.FileStorage.Path)
	setOptionIfEmptyInt(&Options.FileStorage.Mode, configFromFile.FileStorage.Mode)
	setOptionIfEmpty(&Options.DatabaseDSN, configFromFile.DatabaseDSN)
	setOptionIfEmptyBool(&Options.EnableHTTPS, configFromFile.EnableHTTPS)
	setOptionIfEmpty(&Options.SSLConfig.CertFile, configFromFile.SSLConfig.CertFile)
	setOptionIfEmpty(&Options.SSLConfig.KeyFile, configFromFile.SSLConfig.KeyFile)
}

// parseFileConfigJSON парсинг файла JSON конфигурации
func parseFileConfigJSON(configFilePath *string) OptionsStruct {
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

	return configFromFile
}

// setOptionIfEmpty Устанавливает значение, если оно пустое для string
func setOptionIfEmpty(target *string, value string) {
	if *target == "" {
		*target = value
	}
}

// setOptionIfEmptyBool Устанавливает значение, если оно пустое для bool
func setOptionIfEmptyBool(target *bool, value bool) {
	if !*target {
		*target = value
	}
}

// setOptionIfEmptyInt Устанавливает значение, если оно пустое для int
func setOptionIfEmptyInt(target *int, value int) {
	if *target == 0 {
		*target = value
	}
}

// stringToBool Костыль для env с целью преобразования текста в bool
func stringToBool(value string) bool {
	return value == "true"
}
