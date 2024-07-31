package serverservices

import (
	"log"
	"net/http"

	"github.com/Alandres998/url-shortner/internal/app/db/db"
	fileservices "github.com/Alandres998/url-shortner/internal/app/db/fileServices"
	"github.com/Alandres998/url-shortner/internal/app/routers"
	"github.com/Alandres998/url-shortner/internal/config"
	"github.com/gin-gonic/gin"
)

func RunServer() {
	config.InitConfig()
	db.InitDB()
	cfg := config.Options.ServerAdress
	gin.SetMode(gin.ReleaseMode)
	//syncservices.InitURLStorage()
	fileservices.InitFileStorage()
	routersInit := routers.InitRouter()

	server := &http.Server{
		Addr:    cfg.MainURLServer,
		Handler: routersInit,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Ахтунг сервер прилег: %s\n", err)
	}
}
