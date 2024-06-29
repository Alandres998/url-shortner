package config

import (
	"flag"
	"fmt"
)

var Options struct {
	ServerAdress ServerConfig
}

type ServerConfig struct {
	MainURLServer string
	ShortURL      string
}

func init() {
	flag.StringVar(&Options.ServerAdress.MainURLServer, "a", "8080", "basic main address")
	flag.StringVar(&Options.ServerAdress.ShortURL, "b", "8080", "short response address")
}

func GetAdressServer(Port string) string {
	text := fmt.Sprintf("http://localhost:%s", Port)
	return text
}
