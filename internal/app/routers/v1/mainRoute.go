package v1

import (
	"net/http"

	webservices "github.com/Alandres998/url-shortner/internal/app/webServices"
	"github.com/gin-gonic/gin"
)

func WebInterfaceShort(c *gin.Context) {
	responseText, err := webservices.Shorter(c)
	if err != nil {
		webservices.GetErrorWithCode(c, err.Error(), http.StatusBadRequest)
		return
	}
	c.String(http.StatusCreated, responseText)
}

func WebInterfaceFull(c *gin.Context) {
	c.String(http.StatusOK, "test")
}
