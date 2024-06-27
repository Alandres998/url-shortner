package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"sync"
)

var DynamicHostDNS = false

// Map для хранения сокращённых и оригинальных URL
var urlStore = struct {
	sync.RWMutex
	m map[string]string
}{m: make(map[string]string)}

type MyHandler struct{}

// Наверняка потом усложнится и нужно будет добавлять кастомный текст
func getErrorCode400(res http.ResponseWriter, errorText string) {
	http.Error(res, errorText, http.StatusBadRequest)
}

// Функция шортер
func generateShortURL() string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	shortURL := make([]byte, 8)
	for i := range shortURL {
		shortURL[i] = chars[rand.Intn(len(chars))]
	}
	return string(shortURL)
}

// Хендлер сокращения
func Shorter(res http.ResponseWriter, req *http.Request) {
	var mainURL string
	//Проверка на метод
	if req.Method != http.MethodPost {
		getErrorCode400(res, "Ошибка")
		return
	}
	//Проверка на метод и тело содержимого
	body, err := io.ReadAll(req.Body)
	if err != nil || len(body) == 0 {
		getErrorCode400(res, "Ошибка")
		return
	}

	originalURL := string(body)
	if DynamicHostDNS {
		mainURL = string(body)
	} else {
		mainURL = "http://localhost:8080"
	}

	codeURL := generateShortURL()
	shortedCode := fmt.Sprintf("%s/%s", mainURL, codeURL)

	urlStore.Lock()
	urlStore.m[codeURL] = originalURL
	urlStore.Unlock()

	res.WriteHeader(http.StatusCreated)
	res.Header().Set("Content-Type", "text/plain")
	res.Write([]byte(shortedCode))
}

// Обработчик для возврата полной строки
func Fuller(res http.ResponseWriter, req *http.Request) {
	//В первые плачу что не могу использовать регулярку :D
	id := req.URL.Path[1:]

	urlStore.RLock()
	originalURL, exists := urlStore.m[id]
	urlStore.RUnlock()

	if !exists {
		getErrorCode400(res, "Ошибка")
		return
	}

	// Перенаправление на оригинальный URL
	res.Header().Set("Location", originalURL)
	res.WriteHeader(http.StatusTemporaryRedirect)
}

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPost {
			Shorter(res, req)
		} else if req.Method == http.MethodGet {
			Fuller(res, req)
		} else {
			getErrorCode400(res, "Ошибка")
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Ахтунг сервер прилег: %s\n", err)
	}
}
