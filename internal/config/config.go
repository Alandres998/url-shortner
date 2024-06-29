package config

import "flag"

var Options struct {
	ServerAdress ServerConfig
}

type ServerConfig struct {
	MainURLServer string
	ShortURL      string
}

func init() {
	flag.StringVar(&Options.ServerAdress.MainURLServer, "a", "localhost:8080", "basic main address")
	flag.StringVar(&Options.ServerAdress.ShortURL, "b", "localhost:8080", "short response address")
}
