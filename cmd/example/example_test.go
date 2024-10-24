package main

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Alandres998/url-shortner/internal/app/db/storagefactory"
	"github.com/Alandres998/url-shortner/internal/app/routers"
	"github.com/Alandres998/url-shortner/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var Router *gin.Engine

func setupRouter() {
	if Router == nil {
		config.InitConfig()
		storagefactory.NewStorage()
		routersInit := routers.InitRouter()
		Router = routersInit
	}
}

func TestWebInterfaceShort(t *testing.T) {
	setupRouter()
	ts := httptest.NewServer(Router)
	defer ts.Close()

	// Подготовка данных для POST запроса
	data := []byte("https://valhalla.ru/")

	// Создание нового запроса
	req, err := http.NewRequest("POST", ts.URL+"/", bytes.NewBuffer(data))
	assert.NoError(t, err)

	// Установка заголовков
	req.Header.Set("Content-Type", "text/plain")
	req.Header.Set("Cookie", "user_id=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiZTdiY2U4MWQtYmU4NC00MjgwLTkxODctOGMzYjY3OGYzZjM5In0._TJR8eieyvf61NFTHkAu0o2ZHeacLviAYHzHH3WFgaM")

	// Отправка запроса
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer func() {
		if errBodyClose := resp.Body.Close(); errBodyClose != nil {
			log.Printf("Ошибка при закрытии чтения тела ответа: %v", errBodyClose)
		}
	}()
}

func TestWebInterfacePing(t *testing.T) {
	setupRouter()
	ts := httptest.NewServer(Router)
	defer ts.Close()

	// Создание GET запроса для проверки состояния сервера
	req, err := http.NewRequest("GET", ts.URL+"/ping", nil)
	assert.NoError(t, err)

	// Отправка запроса
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer func() {
		if errBodyClose := resp.Body.Close(); errBodyClose != nil {
			log.Printf("Ошибка при закрытии чтения тела ответа: %v", errBodyClose)
		}
	}()
}

func TestWebInterfaceFull(t *testing.T) {
	setupRouter()
	ts := httptest.NewServer(Router)
	defer ts.Close()

	// Создание GET запроса с параметром id
	req, err := http.NewRequest("GET", ts.URL+"/some_short_url_id", nil)
	assert.NoError(t, err)

	// Отправка запроса
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer func() {
		if errBodyClose := resp.Body.Close(); errBodyClose != nil {
			log.Printf("Ошибка при закрытии чтения тела ответа: %v", errBodyClose)
		}
	}()

}

func TestWebInterfaceShortenJSON(t *testing.T) {
	setupRouter()
	ts := httptest.NewServer(Router)
	defer ts.Close()

	// Подготовка данных для POST запроса
	data := []byte(`{"url": "https://valhallajson/"}`)

	// Создание нового запроса
	req, err := http.NewRequest("POST", ts.URL+"/api/shorten", bytes.NewBuffer(data))
	assert.NoError(t, err)

	// Установка заголовков
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "user_id=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiZTdiY2U4MWQtYmU4NC00MjgwLTkxODctOGMzYjY3OGYzZjM5In0._TJR8eieyvf61NFTHkAu0o2ZHeacLviAYHzHH3WFgaM")

	// Отправка запроса
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer func() {
		if errBodyClose := resp.Body.Close(); errBodyClose != nil {
			log.Printf("Ошибка при закрытии чтения тела ответа: %v", errBodyClose)
		}
	}()
}

func TestWebInterfaceShortenJSONBatch(t *testing.T) {
	setupRouter()
	ts := httptest.NewServer(Router)
	defer ts.Close()

	// Подготовка данных для POST запроса
	data := []byte(`[{"url": "https://valhallajson23/"}, {"url": "https://valhallajson21/"}]`)

	// Создание нового запроса
	req, err := http.NewRequest("POST", ts.URL+"/api/shorten/batch", bytes.NewBuffer(data))
	assert.NoError(t, err)

	// Установка заголовков
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "user_id=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiZTdiY2U4MWQtYmU4NC00MjgwLTkxODctOGMzYjY3OGYzZjM5In0._TJR8eieyvf61NFTHkAu0o2ZHeacLviAYHzHH3WFgaM")

	// Отправка запроса
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer func() {
		if errBodyClose := resp.Body.Close(); errBodyClose != nil {
			log.Printf("Ошибка при закрытии чтения тела ответа: %v", errBodyClose)
		}
	}()

}

func TestWebInterfaceGetAllShortURLByCookie(t *testing.T) {
	setupRouter()
	ts := httptest.NewServer(Router)
	defer ts.Close()

	// Создание GET запроса для получения всех коротких URL
	req, err := http.NewRequest("GET", ts.URL+"/api/user/urls", nil)
	assert.NoError(t, err)

	// Установка заголовков
	req.Header.Set("Cookie", "user_id=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiZTdiY2U4MWQtYmU4NC00MjgwLTkxODctOGMzYjY3OGYzZjM5In0._TJR8eieyvf61NFTHkAu0o2ZHeacLviAYHzHH3WFgaM")

	// Отправка запроса
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer func() {
		if errBodyClose := resp.Body.Close(); errBodyClose != nil {
			log.Printf("Ошибка при закрытии чтения тела ответа: %v", errBodyClose)
		}
	}()

}

func TestWebInterfaceDeleteShortURL(t *testing.T) {
	setupRouter()
	ts := httptest.NewServer(Router)
	defer ts.Close()

	// Подготовка данных для DELETE запроса
	data := []byte(`["short_url_id_1", "short_url_id_2"]`)

	// Создание нового запроса
	req, err := http.NewRequest("DELETE", ts.URL+"/api/user/urls", bytes.NewBuffer(data))
	assert.NoError(t, err)

	// Установка заголовков
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "user_id=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiZTdiY2U4MWQtYmU4NC00MjgwLTkxODctOGMzYjY3OGYzZjM5In0._TJR8eieyvf61NFTHkAu0o2ZHeacLviAYHzHH3WFgaM")

	// Отправка запроса
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer func() {
		if errBodyClose := resp.Body.Close(); errBodyClose != nil {
			log.Printf("Ошибка при закрытии чтения тела ответа: %v", errBodyClose)
		}
	}()
}
