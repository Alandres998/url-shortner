package webservices_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Alandres998/url-shortner/internal/app/routers"
	"github.com/stretchr/testify/assert"
)

func TestWebInterfaceShorten(t *testing.T) {
	router := routers.InitRouter()

	w := httptest.NewRecorder()
	body := `{"url":"https://practicum.yandex.ru"}`
	req, _ := http.NewRequest("POST", "/api/shorten", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Contains(t, response["result"], "http://localhost:8080/")
}

func TestWebInterfaceShorten_InvalidJSON(t *testing.T) {
	router := routers.InitRouter()

	w := httptest.NewRecorder()
	body := `{"invalid"}`
	req, _ := http.NewRequest("POST", "/api/shorten", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
