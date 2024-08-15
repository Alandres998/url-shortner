package routers

import (
	"net/http"

	middlewares "github.com/Alandres998/url-shortner/internal/app/middleware"
	v1 "github.com/Alandres998/url-shortner/internal/app/routers/v1"
	webservices "github.com/Alandres998/url-shortner/internal/app/webServices"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.Logger())
	r.Use(middlewares.GzipMiddleware())
	r.Use(middlewares.AuthMiddleware())

	r.NoRoute(func(c *gin.Context) {
		webservices.GetErrorWithCode(c, webservices.Error400DefaultText, http.StatusBadRequest)
	})

	mainRouteGroup := r.Group("/")
	{
		mainRouteGroup.POST("/", v1.WebInterfaceShort)
		mainRouteGroup.GET("/ping", v1.WebInterfacePing)
		mainRouteGroup.GET("/:id", v1.WebInterfaceFull)
	}

	apiRouteGroup := r.Group("/api")
	{
		apiRouteGroup.POST("/shorten", v1.WebInterfaceShortenJSON)
		apiRouteGroup.POST("/shorten/batch", v1.WebInterfaceShortenJSONBatch)
		apiRouteGroup.GET("/user/urls", v1.WebInterfaceGetAllShortURLByCookie)
		apiRouteGroup.DELETE("/user/urls", v1.WebInterfaceDeleteShortURL)
	}
	return r
}
