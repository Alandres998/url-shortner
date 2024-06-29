package serverservices

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	syncservices "github.com/Alandres998/url-shortner/internal/app/db/syncServices"
	"github.com/Alandres998/url-shortner/internal/app/routers"
	"github.com/Alandres998/url-shortner/internal/config"
	"github.com/gin-gonic/gin"
)

func RunServer() {
	flag.Parse()
	gin.SetMode(gin.ReleaseMode)
	syncservices.InitURLStorage()
	routersInit := routers.InitRouter()
	//Ех сейчас бы env
	endPoint := fmt.Sprintf(":%s", config.Options.ServerAdress.MainURLServer)

	server := &http.Server{
		Addr:    endPoint,
		Handler: routersInit,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Ахтунг сервер прилег: %s\n", err)
	}
}
