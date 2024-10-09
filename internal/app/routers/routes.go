package routers

import (
	"net/http"
	"net/http/pprof"

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

	debugRouteGroup := r.Group("/debug/pprof")
	{
		debugRouteGroup.GET("/", gin.WrapF(pprof.Index))
		debugRouteGroup.GET("/cmdline", gin.WrapH(pprof.Handler("cmdline")))
		debugRouteGroup.GET("/block", gin.WrapH(pprof.Handler("block")))
		debugRouteGroup.GET("/goroutine", gin.WrapH(pprof.Handler("goroutine")))
		debugRouteGroup.GET("/mutex", gin.WrapH(pprof.Handler("mutex")))
		debugRouteGroup.GET("/threadcreate", gin.WrapH(pprof.Handler("threadcreate")))
		debugRouteGroup.GET("/heap", gin.WrapH(pprof.Handler("heap")))
		debugRouteGroup.GET("/profile", gin.WrapH(pprof.Handler("profile")))
		debugRouteGroup.GET("/symbol", gin.WrapH(pprof.Handler("symbol")))
		debugRouteGroup.GET("/trace", gin.WrapH(pprof.Handler("trace")))
		debugRouteGroup.GET("/allocs", gin.WrapH(pprof.Handler("allocs")))
	}

	return r
}
