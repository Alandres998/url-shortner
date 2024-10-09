package main

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWebInterfaceShort(t *testing.T) {
	// Подготовка данных для POST запроса
	data := []byte("https://valhalla.ru/")

	// Создание нового запроса
	req, err := http.NewRequest("POST", "http://localhost:8080/", bytes.NewBuffer(data))
	assert.NoError(t, err)

	// Установка заголовков
	req.Header.Set("Content-Type", "text/plain")
	req.Header.Set("Cookie", "user_id=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiZTdiY2U4MWQtYmU4NC00MjgwLTkxODctOGMzYjY3OGYzZjM5In0._TJR8eieyvf61NFTHkAu0o2ZHeacLviAYHzHH3WFgaM")

	// Отправка запроса
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer resp.Body.Close()

}

func TestWebInterfacePing(t *testing.T) {
	// Создание GET запроса для проверки состояния сервера
	req, err := http.NewRequest("GET", "http://localhost:8080/ping", nil)
	assert.NoError(t, err)

	// Отправка запроса
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer resp.Body.Close()
}

func TestWebInterfaceFull(t *testing.T) {
	// Создание GET запроса с параметром id
	req, err := http.NewRequest("GET", "http://localhost:8080/some_short_url_id", nil)
	assert.NoError(t, err)

	// Отправка запроса
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer resp.Body.Close()

}

func TestWebInterfaceShortenJSON(t *testing.T) {
	// Подготовка данных для POST запроса
	data := []byte(`{"url": "https://valhallajson/"}`)

	// Создание нового запроса
	req, err := http.NewRequest("POST", "http://localhost:8080/api/shorten", bytes.NewBuffer(data))
	assert.NoError(t, err)

	// Установка заголовков
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "user_id=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiZTdiY2U4MWQtYmU4NC00MjgwLTkxODctOGMzYjY3OGYzZjM5In0._TJR8eieyvf61NFTHkAu0o2ZHeacLviAYHzHH3WFgaM")

	// Отправка запроса
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer resp.Body.Close()
}

func TestWebInterfaceShortenJSONBatch(t *testing.T) {
	// Подготовка данных для POST запроса
	data := []byte(`[{"url": "https://valhallajson23/"}, {"url": "https://valhallajson21/"}]`)

	// Создание нового запроса
	req, err := http.NewRequest("POST", "http://localhost:8080/api/shorten/batch", bytes.NewBuffer(data))
	assert.NoError(t, err)

	// Установка заголовков
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "user_id=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiZTdiY2U4MWQtYmU4NC00MjgwLTkxODctOGMzYjY3OGYzZjM5In0._TJR8eieyvf61NFTHkAu0o2ZHeacLviAYHzHH3WFgaM")

	// Отправка запроса
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer resp.Body.Close()
}

func TestWebInterfaceGetAllShortURLByCookie(t *testing.T) {
	// Создание GET запроса для получения всех коротких URL
	req, err := http.NewRequest("GET", "http://localhost:8080/api/user/urls", nil)
	assert.NoError(t, err)

	// Установка заголовков
	req.Header.Set("Cookie", "user_id=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiZTdiY2U4MWQtYmU4NC00MjgwLTkxODctOGMzYjY3OGYzZjM5In0._TJR8eieyvf61NFTHkAu0o2ZHeacLviAYHzHH3WFgaM")

	// Отправка запроса
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer resp.Body.Close()
}

func TestWebInterfaceDeleteShortURL(t *testing.T) {
	// Подготовка данных для DELETE запроса
	data := []byte(`["short_url_id_1", "short_url_id_2"]`)

	// Создание нового запроса
	req, err := http.NewRequest("DELETE", "http://localhost:8080/api/user/urls", bytes.NewBuffer(data))
	assert.NoError(t, err)

	// Установка заголовков
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "user_id=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiZTdiY2U4MWQtYmU4NC00MjgwLTkxODctOGMzYjY3OGYzZjM5In0._TJR8eieyvf61NFTHkAu0o2ZHeacLviAYHzHH3WFgaM")

	// Отправка запроса
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer resp.Body.Close()

}
