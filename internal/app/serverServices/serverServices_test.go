package serverservices

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Alandres998/url-shortner/internal/app/db/storagefactory"
	"github.com/Alandres998/url-shortner/internal/app/routers"
	webservices "github.com/Alandres998/url-shortner/internal/app/webServices"
	"github.com/Alandres998/url-shortner/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	return routers.InitRouter()
}

func TestMain(m *testing.M) {
	os.Setenv("RUN_MODE", "test")
	config.InitConfig()
	storagefactory.NewStorage()
	code := m.Run()
	os.Exit(code)
}

func TestShorten(t *testing.T) {
	router := setupRouter()

	body := bytes.NewBufferString(`http://valhalla.com`)
	req, _ := http.NewRequest("POST", "/", body)
	req.Header.Set("Content-Type", "text/plain")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestShortenJSON(t *testing.T) {
	router := setupRouter()

	body := `{"url": "http://valhalla.com"}`
	req, _ := http.NewRequest("POST", "/api/shorten", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	// Тестируем, что вернулся правильный URL
	var response webservices.ShortenResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.NotEmpty(t, response.Result)
}

func TestShortenJSONBatch(t *testing.T) {
	router := setupRouter()

	batchRequest := []webservices.BatchRequest{
		{CorrelationID: "1", OriginalURL: "http://example.com"},
		{CorrelationID: "2", OriginalURL: "http://example.org"},
	}
	body, _ := json.Marshal(batchRequest)

	req, _ := http.NewRequest("POST", "/api/shorten/batch", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	// Тестируем, что вернулся правильный формат ответа
	var responses []webservices.BatchResponse
	err := json.Unmarshal(w.Body.Bytes(), &responses)
	assert.NoError(t, err)
	assert.Len(t, responses, 2)
	for _, res := range responses {
		assert.NotEmpty(t, res.ShortURL)
	}
}
