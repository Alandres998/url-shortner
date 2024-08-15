package shortener

import (
	"context"
	"math/rand"
	"sync"

	"github.com/Alandres998/url-shortner/internal/app/db/storage"
	"github.com/Alandres998/url-shortner/internal/app/service/logger"
)

func GenerateShortURL() string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	shortURL := make([]byte, 8)
	for i := range shortURL {
		shortURL[i] = chars[rand.Intn(len(chars))]
	}
	return string(shortURL)
}

func DeleteShortURL(userID string, shortURLs []string) {
	batchSize := 2
	urlChan := make(chan string, len(shortURLs))
	var wg sync.WaitGroup

	for _, shortURL := range shortURLs {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			urlChan <- url
		}(shortURL)
	}

	go func() {
		wg.Wait()
		close(urlChan)
	}()

	go func() {
		var buffer []string
		for shortURL := range urlChan {
			buffer = append(buffer, shortURL)
			if len(buffer) >= batchSize {
				err := storage.Store.DeleteUserURL(context.Background(), buffer, userID)
				if err != nil {
					logger.LogError("Delete Short URL", err.Error())
				}
				buffer = buffer[:0]
			}
		}

		if len(buffer) > 0 {
			err := storage.Store.DeleteUserURL(context.Background(), buffer, userID)
			if err != nil {
				logger.LogError("Delete Short URL", err.Error())
			}
		}
	}()
}
