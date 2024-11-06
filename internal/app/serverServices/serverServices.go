package serverservices

import (
	"context"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Alandres998/url-shortner/internal/app/db/storagefactory"
	"github.com/Alandres998/url-shortner/internal/app/routers"
	"github.com/Alandres998/url-shortner/internal/config"
	"github.com/gin-gonic/gin"
)

// RunServer запускает сервер
func RunServer() {
	config.InitConfig()

	// Создание канала для приема сигналов
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	shutdownTimeout := 15 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	storagefactory.NewStorage()
	cfg := config.Options.ServerAdress
	gin.SetMode(gin.ReleaseMode)

	routersInit := routers.InitRouter()

	server := &http.Server{
		Addr:    cfg.MainURLServer,
		Handler: routersInit,
	}

	go func() {
		if config.Options.EnableHTTPS {
			log.Println("Запуск сервера с HTTPS...")
			if err := server.ListenAndServeTLS(config.Options.SSLConfig.CertFile, config.Options.SSLConfig.KeyFile); err != nil && err != http.ErrServerClosed {
				log.Fatalf("Ошибка запуска HTTPS-сервера: %s\n", err)
			}
		} else {
			log.Println("Запуск сервера с HTTP...")
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("Ошибка запуска HTTP-сервера: %s\n", err)
			}
		}
	}()

	<-signalChan
	log.Println("Получен сигнал для завершения работы сервера...")

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Ошибка при завершении работы сервера: %s", err)
	} else {
		log.Println("Сервер завершил работу корректно.")
	}
}
