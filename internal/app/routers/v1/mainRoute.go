package v1

import (
	"context"
	"errors"
	"io"
	"net/http"

	"github.com/Alandres998/url-shortner/internal/app/db/storage"
	"github.com/Alandres998/url-shortner/internal/app/service/auth"
	"github.com/Alandres998/url-shortner/internal/app/service/logger"
	"github.com/Alandres998/url-shortner/internal/app/service/shortener"
	webservices "github.com/Alandres998/url-shortner/internal/app/webServices"
	"github.com/gin-gonic/gin"
)

// WebInterfaceShort Веб интерфейс сокращения ссылок
func WebInterfaceShort(c *gin.Context) {
	ctx := c.Request.Context()
	req := c.Request

	userID, err := auth.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Не авторизваон"})
		return
	}

	const maxBodySize = 1024 // 1MB
	body, err := io.ReadAll(io.LimitReader(req.Body, maxBodySize))
	if err != nil || len(body) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Чет не то написал в теле"})
		return
	}

	originalURL := string(body)
	responseText, err := webservices.ShorterGeneral(ctx, userID, originalURL)
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

// WebInterfaceShort Веб интерфейс вернуть полную ссылку
func WebInterfaceFull(c *gin.Context) {
	id := c.Param("id")
	responseHeaderLocation, err := webservices.Fuller(c, id)
	if err != nil {
		if errors.Is(err, storage.ErrURLDeleted) {
			webservices.GetErrorWithCode(c, err.Error(), http.StatusGone)
		} else {
			webservices.GetErrorWithCode(c, err.Error(), http.StatusBadRequest)
		}
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, responseHeaderLocation)
}

// WebInterfaceShortenJSON Веб интерфейс сокращения ссылок для json
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

// WebInterfaceShortenJSONBatch Веб интерфейс сокращения ссылок для json батчами
func WebInterfaceShortenJSONBatch(c *gin.Context) {
	responseJSON, err := webservices.ShorterJSONBatch(c)
	if err != nil {
		webservices.GetErrorWithCode(c, err.Error(), http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusCreated, responseJSON)
}

// WebInterfacePing Веб интерфейс проверка доступности хранилища
func WebInterfacePing(c *gin.Context) {
	ctx := context.Background()
	err := storage.Store.Ping(ctx)
	if err != nil {
		webservices.GetErrorWithCode(c, err.Error(), http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

// WebInterfacePing Веб интерфейс получения всех сокращалок пользователя
func WebInterfaceGetAllShortURLByCookie(c *gin.Context) {
	var statusCode int
	var responseJSON []webservices.ShortUserResponse

	authenticator := &auth.AuthService{}
	userURLs, err := webservices.GetAllUserShorterURL(c, authenticator)
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

// WebInterfacePing Веб интерфейс удаления ссылки
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

// WebInterfaceStats Веб инетрфейс для контролера по получению статистики
func WebInterfaceStats(c *gin.Context) {
	var statusCode int

	responseJSON, err := webservices.GetStatisticsShortURL(c)
	if err != nil {
		statusCode = http.StatusForbidden
	} else {
		statusCode = http.StatusOK
	}

	c.JSON(statusCode, responseJSON)
}
