package serverservices

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Alandres998/url-shortner/internal/app/db/storagefactory"
	"github.com/Alandres998/url-shortner/internal/app/routers"
	"github.com/Alandres998/url-shortner/internal/config"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func RunServer() {
	config.InitConfig()

	// Канал для обработки сигналов
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	shutdownTimeout := 15 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	storagefactory.NewStorage()
	cfg := config.Options.ServerAdress
	gin.SetMode(gin.ReleaseMode)

	routersInit := routers.InitRouter()

	// HTTP сервер
	server := &http.Server{
		Addr:    cfg.MainURLServer,
		Handler: routersInit,
	}

	// gRPC сервер
	grpcServer := grpc.NewServer()

	go func() {
		if err := startServer(server, config.Options.EnableHTTPS, config.Options.SSLConfig); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Ошибка запуска HTTP сервера: %s\n", err)
		}
	}()

	go func() {
		lis, err := net.Listen("tcp", config.Options.GRPCPort)
		if err != nil {
			log.Fatalf("Ошибка при запуске gRPC-сервера: %v", err)
		}

		log.Printf("Запуск gRPC-сервера на %s...\n", config.Options.GRPCPort)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Ошибка при работе gRPC-сервера: %v", err)
		}
	}()

	<-signalChan
	log.Println("Получен сигнал для завершения работы сервера...")

	// Завершаем работу серверов с обработкой таймаута
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Ошибка при завершении работы HTTP-сервера: %s", err)
	} else if ctx.Err() == context.DeadlineExceeded {
		log.Println("Ушел в таймаут при завершении HTTP-сервера")
	} else {
		log.Println("HTTP сервер завершил работу корректно.")
	}

	grpcServer.GracefulStop()
	log.Println("gRPC сервер завершил работу.")
}

func startServer(server *http.Server, enableHTTPS bool, sslConfig config.SSLConfig) error {
	if enableHTTPS {
		log.Println("Запуск сервера с HTTPS...")
		return server.ListenAndServeTLS(sslConfig.CertFile, sslConfig.KeyFile)
	}

	log.Println("Запуск сервера с HTTP...")
	return server.ListenAndServe()
}
