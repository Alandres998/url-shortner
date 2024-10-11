package webservices

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

type MockLogger struct{}

func (m *MockLogger) LogError(msg string, err string) {

}

func MockGetUserID(c *gin.Context) (string, error) {
	return "user123", nil
}

func MockDeleteShortURL(userID string, shortURLs []string) {

}

func BenchmarkDeleteShortURL(b *testing.B) {
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(recorder)

	shortURLs := []string{"short1", "short2", "short3"}
	body, _ := json.Marshal(shortURLs)
	c.Request = &http.Request{Body: io.NopCloser(bytes.NewReader(body))}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DeleteShortURL(c)
	}
}
