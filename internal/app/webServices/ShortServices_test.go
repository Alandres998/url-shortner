package webservices_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Alandres998/url-shortner/internal/app/routers"
	"github.com/stretchr/testify/assert"
)

func TestWebInterfaceShorten_InvalidJSON(t *testing.T) {
	router := routers.InitRouter()

	w := httptest.NewRecorder()
	body := `{"invalid"}`
	req, _ := http.NewRequest("POST", "/api/shorten", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
