package shortener

import (
	"context"
	"math/rand"

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
	go func() {
		err := storage.Store.DeleteUserURL(context.Background(), shortURLs, userID)
		if err != nil {
			logger.LogError("DeleteShortURL", err.Error())
		}
	}()
	// const batchSize = 10

	// urlChan := make(chan string, len(shortURLs))
	// var wg sync.WaitGroup

	// // Заполняем канал URL-ами
	// go func() {
	// 	for _, shortURL := range shortURLs {
	// 		urlChan <- shortURL
	// 	}
	// 	close(urlChan)
	// }()

	// // Обрабатываем URL-ы пакетами
	// go func() {
	// 	defer wg.Wait()
	// 	var buffer []string
	// 	for shortURL := range urlChan {
	// 		buffer = append(buffer, shortURL)
	// 		if len(buffer) >= batchSize {
	// 			wg.Add(1)
	// 			go func(urls []string) {
	// 				defer wg.Done()
	// 				err := storage.Store.DeleteUserURL(context.Background(), urls, userID)
	// 				if err != nil {
	// 					logger.LogError("DeleteShortURL", err.Error())
	// 				}
	// 			}(buffer)
	// 			buffer = nil
	// 		}
	// 	}

	// 	// Обрабатываем оставшиеся URL-ы
	// 	if len(buffer) > 0 {
	// 		wg.Add(1)
	// 		go func(urls []string) {
	// 			defer wg.Done()
	// 			err := storage.Store.DeleteUserURL(context.Background(), urls, userID)
	// 			if err != nil {
	// 				logger.LogError("DeleteShortURL", err.Error())
	// 			}
	// 		}(buffer)
	// 	}
	// }()
}
