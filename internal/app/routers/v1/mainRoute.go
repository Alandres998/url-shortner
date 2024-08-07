package v1

import (
	"errors"
	"net/http"

	"github.com/Alandres998/url-shortner/internal/app/db/storage"
	webservices "github.com/Alandres998/url-shortner/internal/app/webServices"
	"github.com/gin-gonic/gin"
)

func WebInterfaceShort(c *gin.Context) {
	responseText, err := webservices.Shorter(c)
	statusCode := http.StatusCreated
	if err != nil && errors.Is(err, storage.ErrURLExists) {
		err = nil
		statusCode = http.StatusConflict
	}
	if err != nil {
		webservices.GetErrorWithCode(c, err.Error(), http.StatusBadRequest)
		return
	}

	c.String(statusCode, string(responseText))
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
	statusCode := http.StatusCreated
	if err != nil && errors.Is(err, storage.ErrURLExists) {
		err = nil
		statusCode = http.StatusConflict
	}
	if err != nil {
		webservices.GetErrorWithCode(c, err.Error(), http.StatusBadRequest)
		return
	}
	c.JSON(statusCode, responseJSON)
}

func WebInterfaceShortenJSONBatch(c *gin.Context) {
	responseJSON, err := webservices.ShorterJSONBatch(c)
	if err != nil {
		webservices.GetErrorWithCode(c, err.Error(), http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusCreated, responseJSON)
}

func WebInterfacePing(c *gin.Context) {
	err := storage.Store.Ping()
	if err != nil {
		webservices.GetErrorWithCode(c, err.Error(), http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
