package v1

import (
	"context"
	"errors"
	"net/http"

	"github.com/Alandres998/url-shortner/internal/app/db/storage"
	"github.com/Alandres998/url-shortner/internal/app/service/auth"
	"github.com/Alandres998/url-shortner/internal/app/service/logger"
	"github.com/Alandres998/url-shortner/internal/app/service/shortener"
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
		if errors.Is(err, storage.ErrURLDeleted) {
			webservices.GetErrorWithCode(c, err.Error(), http.StatusGone)
			return
		} else {
			webservices.GetErrorWithCode(c, err.Error(), http.StatusBadRequest)
			return
		}

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
	ctx := context.Background()
	err := storage.Store.Ping(ctx)
	if err != nil {
		webservices.GetErrorWithCode(c, err.Error(), http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func WebInterfaceGetAllShortURLByCookie(c *gin.Context) {
	var statusCode int
	var responseJSON []webservices.ShortUserResponse

	userURLs, err := webservices.GetAllUserShorterURL(c)
	if err != nil {
		if errors.Is(err, webservices.Error401DefaultText) {
			statusCode = http.StatusUnauthorized
			responseJSON = []webservices.ShortUserResponse{}
		} else if errors.Is(err, webservices.Error204DefaultText) {
			statusCode = http.StatusNoContent
			responseJSON = []webservices.ShortUserResponse{}
		} else {
			statusCode = http.StatusBadRequest
			responseJSON = []webservices.ShortUserResponse{}
		}
	} else {
		statusCode = http.StatusOK
		responseJSON = userURLs
	}
	c.JSON(statusCode, responseJSON)
}

func WebInterfaceDeleteShortURL(c *gin.Context) {
	var shortURLs []string
	if err := c.BindJSON(&shortURLs); err != nil {
		logger.LogError("Shorter Delete", err.Error())
		c.String(http.StatusBadRequest, "")
		return
	}

	userID, err := auth.GetUserIDByCookie(c)
	if err != nil {
		logger.LogError("Shorter Delete", err.Error())
		c.String(http.StatusBadRequest, "")
		return
	}

	go func() {
		shortener.DeleteShortURL(userID, shortURLs)
	}()
	c.String(http.StatusAccepted, "")
}
