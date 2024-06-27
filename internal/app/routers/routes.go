package routers

import (
	v1 "github.com/Alandres998/url-shortner/internal/app/routers/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	mainRouteGroup := r.Group("/")
	{
		mainRouteGroup.POST("testnewpost", v1.WebInterfaceShort)
		mainRouteGroup.GET("/:id", v1.WebInterfaceFull)
	}
	return r
}
