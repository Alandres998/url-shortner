package config

import (
	"flag"
	"fmt"
	"os"
)

var Options struct {
	ServerAdress ServerConfig
}

type ServerConfig struct {
	MainURLServer string
	ShortURL      string
}

func InitConfig() {
	if os.Getenv("RUN_MODE") == "test" {
		return
	}

	parseFlags()
	loadEnv()
}

func GetAdressServer(Port string) string {
	text := fmt.Sprintf("http://localhost%s", Port)
	return text
}

func parseFlags() {
	flag.StringVar(&Options.ServerAdress.MainURLServer, "a", ":8080", "basic main address")
	flag.StringVar(&Options.ServerAdress.ShortURL, "b", "http://localhost:8080", "short response address")
	flag.Parse()
}

func loadEnv() {
	if envMainURLServer := os.Getenv("SERVER_ADDRESS"); envMainURLServer != "" {
		Options.ServerAdress.MainURLServer = envMainURLServer
	}
	if envShortURL := os.Getenv("BASE_URL"); envShortURL != "" {
		Options.ServerAdress.ShortURL = envShortURL
	}
}
