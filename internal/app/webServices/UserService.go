package webservices

import (
	"context"
	"errors"
	"fmt"

	"github.com/Alandres998/url-shortner/internal/app/db/storage"
	"github.com/Alandres998/url-shortner/internal/app/service/auth"
	"github.com/Alandres998/url-shortner/internal/config"
	"github.com/gin-gonic/gin"
)

type ShortUserResponse struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

var Error401DefaultText = errors.New("нет такого кука")
var Error204DefaultText = errors.New("у пользователя нет ни одного сокращенного URL")

// GetAllUserShorterURL Веб-Сервис получения списка всех сокращенных ссылок.
func GetAllUserShorterURL(c *gin.Context, authenticator auth.Authenticator) ([]ShortUserResponse, error) {
	ctx := context.Background()
	userID, err := authenticator.GetUserIDByCookie(c)
	if err != nil {
		return []ShortUserResponse{}, Error401DefaultText
	}

	storageArray, err := storage.Store.GetUserURLs(ctx, userID)
	if err != nil || len(storageArray) == 0 {
		return []ShortUserResponse{}, Error204DefaultText
	}

	response := make([]ShortUserResponse, 0, len(storageArray))
	for _, urlData := range storageArray {
		response = append(response, ShortUserResponse{
			ShortURL:    fmt.Sprintf("%s/%s", config.Options.ServerAdress.ShortURL, urlData.ShortURL),
			OriginalURL: urlData.OriginalURL,
		})
	}
	return response, nil
}
