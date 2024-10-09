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

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Ахтунг сервер прилег: %s\n", err)
	}
}
