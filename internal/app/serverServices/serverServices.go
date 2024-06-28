package serverservices

import (
	"fmt"
	"log"
	"net/http"

	syncservices "github.com/Alandres998/url-shortner/internal/app/db/syncServices"
	"github.com/Alandres998/url-shortner/internal/app/routers"
	"github.com/gin-gonic/gin"
)

var Port = "8080"

func RunServer() {
	gin.SetMode(gin.ReleaseMode)
	syncservices.InitURLStorage()
	routersInit := routers.InitRouter()
	//Ех сейчас бы env
	endPoint := fmt.Sprintf(":%s", Port)
	server := &http.Server{
		Addr:    endPoint,
		Handler: routersInit,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Ахтунг сервер прилег: %s\n", err)
	}

}
