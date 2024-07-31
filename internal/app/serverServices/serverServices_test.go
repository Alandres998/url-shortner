package serverservices

import (
	"testing"
)

// func setupRouter() *gin.Engine {
// 	return routers.InitRouter()
// }

func TestMain(m *testing.M) {
	// os.Setenv("RUN_MODE", "test")
	// code := m.Run()
	// os.Exit(code)
}

func TestWebInterfaceShort(t *testing.T) {
	// router := setupRouter()

	// w := httptest.NewRecorder()
	// body := bytes.NewBufferString(`{"url":"http://valhalla.com"}`)
	// req, _ := http.NewRequest("POST", "/", body)
	// req.Header.Set("Content-Type", "text/plain")
	// router.ServeHTTP(w, req)
	// assert.Equal(t, http.StatusCreated, w.Code)
}

func TestTestWebInterfaceShortFail(t *testing.T) {
	// router := setupRouter()
	// w := httptest.NewRecorder()

	// body := bytes.NewBufferString(`{"url":"http://example.com"}`)
	// req, _ := http.NewRequest("PUT", "/", body)
	// req.Header.Set("Content-Type", "text/plain")
	// router.ServeHTTP(w, req)

	// assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestWebInterfaceFullFail(t *testing.T) {
	// router := setupRouter()

	// w := httptest.NewRecorder()
	// req, _ := http.NewRequest("GET", "testfail", nil)
	// router.ServeHTTP(w, req)

	// assert.Equal(t, http.StatusBadRequest, w.Code)
}
