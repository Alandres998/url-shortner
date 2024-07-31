package v1

import (
	"net/http"

	"github.com/Alandres998/url-shortner/internal/app/db/db"
	webservices "github.com/Alandres998/url-shortner/internal/app/webServices"
	"github.com/gin-gonic/gin"
)

func WebInterfaceShort(c *gin.Context) {
	responseText, err := webservices.Shorter(c)
	if err != nil {
		webservices.GetErrorWithCode(c, err.Error(), http.StatusBadRequest)
		return
	}

	c.String(http.StatusCreated, string(responseText))
}

func WebInterfaceFull(c *gin.Context) {
	responseHeaderLocation, err := webservices.Fuller(c)
	if err != nil {
		webservices.GetErrorWithCode(c, err.Error(), http.StatusBadRequest)
		return
	}
	c.Redirect(http.StatusTemporaryRedirect, responseHeaderLocation)
}

func WebInterfaceShortenJSON(c *gin.Context) {
	responseJSON, err := webservices.ShorterJSON(c)
	if err != nil {
		webservices.GetErrorWithCode(c, err.Error(), http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusCreated, responseJSON)
}

func WebInterfacePing(c *gin.Context) {
	err := db.DB.Ping()
	if err != nil {
		webservices.GetErrorWithCode(c, err.Error(), http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
