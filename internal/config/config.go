package config

import (
	"flag"
	"fmt"
	"os"
)

var Options struct {
	ServerAdress ServerConfig
	FileStorage  FileStorageConfig
}

type ServerConfig struct {
	MainURLServer string
	ShortURL      string
}

type FileStorageConfig struct {
	Path string
	Mode int
}

func InitConfig() {
	if os.Getenv("RUN_MODE") == "test" {
		return
	}
	parseFlags()
	loadEnv()
	loadConfigFile()
}

func GetAdressServer(Port string) string {
	text := fmt.Sprintf("http://localhost%s", Port)
	return text
}

func parseFlags() {
	flag.StringVar(&Options.ServerAdress.MainURLServer, "a", ":8080", "basic main address")
	flag.StringVar(&Options.ServerAdress.ShortURL, "b", "http://localhost:8080", "short response address")
	flag.StringVar(&Options.FileStorage.Path, "f", "", "storage file")
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
}

func loadConfigFile() {
	if flag.Lookup("f").Value.String() == "" && os.Getenv("FILE_STORAGE_PATH") == "" {
		Options.FileStorage.Path = "/tmp/short-url-db.json"
		Options.FileStorage.Mode = os.O_RDONLY
	} else {
		Options.FileStorage.Mode = os.O_RDWR
	}
}
