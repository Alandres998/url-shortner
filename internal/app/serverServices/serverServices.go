package serverservices

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/Alandres998/url-shortner/internal/app/db/storagefactory"
	"github.com/Alandres998/url-shortner/internal/app/routers"
	"github.com/Alandres998/url-shortner/internal/config"
	"github.com/gin-gonic/gin"
)

// RunServer запускает сервер
func RunServer() {
	config.InitConfig()
	storagefactory.NewStorage()
	cfg := config.Options.ServerAdress
	gin.SetMode(gin.ReleaseMode)

	routersInit := routers.InitRouter()

	server := &http.Server{
		Addr:    cfg.MainURLServer,
		Handler: routersInit,
	}

	if config.Options.EnableHTTPS {
		log.Println("Запуск сервера с HTTPS...")
		if err := server.ListenAndServeTLS(config.Options.SSLConfig.CertFile, config.Options.SSLConfig.KeyFile); err != nil {
			log.Fatalf("Ошибка запуска HTTPS-сервера: %s\n", err)
		}
	} else {
		log.Println("Запуск сервера с HTTP...")
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("Ошибка запуска HTTP-сервера: %s\n", err)
		}
	}
}
