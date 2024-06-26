package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

var DynamicHostDns = false

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
	rand.Seed(time.Now().UnixNano()) //А на каком вообще golang написаны тексты на 1.20 или ниже ?
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	shortURL := make([]byte, 8)
	for i := range shortURL {
		shortURL[i] = chars[rand.Intn(len(chars))]
	}
	return string(shortURL)
}

// Хендлер сокращения
func Shorter(res http.ResponseWriter, req *http.Request) {
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

	var originalURL string
	if DynamicHostDns {
		originalURL = string(body)
	} else {
		originalURL = "http://localhost:8080"
	}

	codeUrl := generateShortURL()
	shortedCode := fmt.Sprintf("%s/%s", originalURL, codeUrl)

	urlStore.Lock()
	urlStore.m[shortedCode] = originalURL
	urlStore.Unlock()

	res.WriteHeader(http.StatusCreated)
	res.Header().Set("Content-Type", "text/plain")
	res.Write([]byte(shortedCode))
}

func Fuller(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Расшировать"))
}

func main() {
	http.HandleFunc("/", Shorter)
	http.HandleFunc("/{id}", Fuller)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Ахтунг сервер прилег: %s\n", err)
	}
}
