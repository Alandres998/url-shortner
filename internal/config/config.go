package config

import (
	"flag"
	"fmt"
	"os"
)

var Options struct {
	ServerAdress ServerConfig
	FileStorage  FileStorageConfig
	DatabaseDSN  string
	StorageType  string
}

type ServerConfig struct {
	MainURLServer string
	ShortURL      string
}

type FileStorageConfig struct {
	Path string
	Mode int
}

const StorageTypeDB = "database"
const StorageTypeFile = "file"
const StorageTypeMemory = "memory"

func InitConfig() {
	if os.Getenv("RUN_MODE") == "test" {
		return
	}
	parseFlags()
	loadEnv()
	loadConfigFile()
	determineStorageType()
}

func InitConfigExample() {
	if os.Getenv("RUN_MODE") == "test" {
		return
	}
	loadEnv()
	loadConfigFile()
	Options.StorageType = StorageTypeMemory
}

func GetAdressServer(Port string) string {
	text := fmt.Sprintf("http://localhost%s", Port)
	return text
}

func parseFlags() {
	flag.StringVar(&Options.ServerAdress.MainURLServer, "a", ":8080", "basic main address")
	flag.StringVar(&Options.ServerAdress.ShortURL, "b", "http://localhost:8080", "short response address")
	flag.StringVar(&Options.FileStorage.Path, "f", "", "storage file")
	flag.StringVar(&Options.DatabaseDSN, "d", "", "database DSN")
	flag.Parse()
}

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

func loadConfigFile() {
	if flag.Lookup("f").Value.String() == "" && os.Getenv("FILE_STORAGE_PATH") == "" {
		Options.FileStorage.Path = "/tmp/short-url-db.json"
		Options.FileStorage.Mode = os.O_RDONLY
	} else {
		Options.FileStorage.Mode = os.O_RDWR
	}
}

func determineStorageType() {
	if Options.DatabaseDSN != "" {
		Options.StorageType = StorageTypeDB
	} else if Options.FileStorage.Path != "" {
		Options.StorageType = StorageTypeFile
	} else {
		Options.StorageType = StorageTypeMemory
	}
}
