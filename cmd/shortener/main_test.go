package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShorter(t *testing.T) {
	testCases := []struct {
		method       string
		body         string
		expectedCode int
		expectedBody string
	}{
		{method: http.MethodPost, body: "http://valhalla.com", expectedCode: http.StatusCreated},
		{method: http.MethodPost, body: "", expectedCode: http.StatusBadRequest, expectedBody: Error400DefaultText},
		{method: http.MethodGet, body: "", expectedCode: http.StatusBadRequest, expectedBody: Error400DefaultText},
		{method: http.MethodPut, body: "", expectedCode: http.StatusBadRequest, expectedBody: Error400DefaultText},
		{method: http.MethodTrace, body: "", expectedCode: http.StatusBadRequest, expectedBody: Error400DefaultText},
		{method: http.MethodOptions, body: "", expectedCode: http.StatusBadRequest, expectedBody: Error400DefaultText},
		{method: http.MethodConnect, body: "", expectedCode: http.StatusBadRequest, expectedBody: Error400DefaultText},
		{method: http.MethodPatch, body: "", expectedCode: http.StatusBadRequest, expectedBody: Error400DefaultText},
		{method: http.MethodDelete, body: "", expectedCode: http.StatusBadRequest, expectedBody: Error400DefaultText},
	}

	for _, tc := range testCases {
		t.Run(tc.method, func(t *testing.T) {
			var req *http.Request
			if tc.method == http.MethodPost {
				req = httptest.NewRequest(tc.method, "/", strings.NewReader(tc.body))
			} else {
				req = httptest.NewRequest(tc.method, "/", nil)
			}
			w := httptest.NewRecorder()

			// Вызовем хендлер
			http.DefaultServeMux.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedCode, w.Code, "Код ответа не совпадает с ожидаемым")
			if tc.expectedBody != "" {
				assert.Equal(t, tc.expectedBody, strings.TrimSpace(w.Body.String()), "Тело ответа не совпадает с ожидаемым")
			} else if tc.method == http.MethodPost && tc.body != "" {
				// Проверка на то, что URL был сокращен
				assert.Contains(t, w.Body.String(), "http://localhost:8080/", "Сокращенный URL не содержит ожидаемый префикс")
			}
		})
	}
}

func TestFuller(t *testing.T) {
	// Предварительно создадим сокращенный URL для теста
	originalURL := "http://valhalla.com"
	shortCode := generateShortURL()
	urlStore.Lock()
	urlStore.m[shortCode] = originalURL
	urlStore.Unlock()
	testCases := []struct {
		method       string
		path         string
		expectedCode int
		expectedLoc  string
		expectedBody string
	}{
		{method: http.MethodGet, path: "/" + shortCode, expectedCode: http.StatusTemporaryRedirect, expectedLoc: originalURL},
		{method: http.MethodGet, path: "/nonexistent", expectedCode: http.StatusBadRequest, expectedBody: Error400DefaultText},
	}

	for _, tc := range testCases {
		t.Run(tc.method+" "+tc.path, func(t *testing.T) {
			req := httptest.NewRequest(tc.method, tc.path, nil)
			w := httptest.NewRecorder()

			// Вызовем хендлер
			http.DefaultServeMux.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedCode, w.Code, "Код ответа не совпадает с ожидаемым")
			if tc.expectedBody != "" {
				assert.Equal(t, tc.expectedBody, strings.TrimSpace(w.Body.String()), "Тело ответа не совпадает с ожидаемым")
			}
			if tc.expectedLoc != "" {
				assert.Equal(t, tc.expectedLoc, w.Header().Get("Location"), "Заголовок Location не совпадает с ожидаемым")
			}
		})
	}
}

func init() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPost {
			Shorter(res, req)
		} else if req.Method == http.MethodGet {
			Fuller(res, req)
		} else {
			getErrorCode400(res, Error400DefaultText)
		}
	})
}
