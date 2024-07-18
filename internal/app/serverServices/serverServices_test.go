package serverservices

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	syncservices "github.com/Alandres998/url-shortner/internal/app/db/syncServices"
	"github.com/Alandres998/url-shortner/internal/app/routers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	syncservices.InitURLStorage()
	return routers.InitRouter()
}

func TestMain(m *testing.M) {
	os.Setenv("RUN_MODE", "test")
	code := m.Run()
	os.Exit(code)
}

func TestWebInterfaceShort(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	body := bytes.NewBufferString(`{"url":"http://valhalla.com"}`)
	req, _ := http.NewRequest("POST", "/", body)
	req.Header.Set("Content-Type", "text/plain")
	router.ServeHTTP(w, req)
	log.Print(w.Body.String())
	assert.Equal(t, http.StatusCreated, w.Code)
}

// func TestWebInterfaceFull(t *testing.T) {
// 	router := setupRouter()

// 	w := httptest.NewRecorder()
// 	body := bytes.NewBufferString(`{"url":"http://valhalla.com"}`)
// 	req, _ := http.NewRequest("POST", "/", body)
// 	req.Header.Set("Content-Type", "text/plain")
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusCreated, w.Code)
// 	shortURL := w.Body.String()

// 	w = httptest.NewRecorder()
// 	req, _ = http.NewRequest("GET", shortURL, nil)
// 	router.ServeHTTP(w, req)
// 	assert.Equal(t, http.StatusTemporaryRedirect, w.Code)
// }

func TestTestWebInterfaceShortFail(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()

	body := bytes.NewBufferString(`{"url":"http://example.com"}`)
	req, _ := http.NewRequest("PUT", "/", body)
	req.Header.Set("Content-Type", "text/plain")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestWebInterfaceFullFail(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "testfail", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
