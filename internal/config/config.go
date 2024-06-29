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

func init() {
	flag.StringVar(&Options.ServerAdress.MainURLServer, "a", ":8080", "basic main address")
	flag.StringVar(&Options.ServerAdress.ShortURL, "b", "http://localhost:8080", "short response address")

	flag.StringVar(&Options.ServerAdress.MainURLServer, "a", ":8080", "basic main address")
	flag.StringVar(&Options.ServerAdress.ShortURL, "b", "http://localhost:8080", "short response address")

	if envMainURLServer := os.Getenv("SERVER_ADDRESS"); envMainURLServer != "" {
		Options.ServerAdress.MainURLServer = envMainURLServer
	}
	if envShortURL := os.Getenv("BASE_URL"); envShortURL != "" {
		Options.ServerAdress.ShortURL = envShortURL
	}

	flag.Parse()
}

func GetAdressServer(Port string) string {
	text := fmt.Sprintf("http://localhost%s", Port)
	return text
}
