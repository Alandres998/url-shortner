package auth_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Alandres998/url-shortner/internal/app/service/auth"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Тестируем GenerateUserID
func TestGenerateUserID(t *testing.T) {
	userID := auth.GenerateUserID()
	require.NotEmpty(t, userID)
}

// Тестируем GenerateJWT
func TestGenerateJWT(t *testing.T) {
	userID := auth.GenerateUserID()
	token, err := auth.GenerateJWT(userID)
	require.NoError(t, err)
	require.NotEmpty(t, token)
}

// Тестируем ValidateJWT
func TestValidateJWT(t *testing.T) {
	userID := auth.GenerateUserID()
	token, err := auth.GenerateJWT(userID)
	require.NoError(t, err)

	parsedToken, err := auth.ValidateJWT(token)
	require.NoError(t, err)
	require.True(t, parsedToken.Valid)
}

func TestSetUserCookie(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := httptest.NewRequest("GET", "/", nil)
	c.Request = req

	auth.SetUserCookie(c, "user123")

	result := w.Result()
	defer func() {
		if errBodyClose := result.Body.Close(); errBodyClose != nil {
			log.Printf("Ошибка при закрытии чтения тела ответа: %v", errBodyClose)
		}
	}()

	cookies := result.Cookies()

	require.Len(t, cookies, 1)
	assert.Equal(t, auth.CookieName, cookies[0].Name)
}

func TestGetUserID_Success(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	token, err := auth.GenerateJWT("user123")
	require.NoError(t, err)
	http.SetCookie(w, &http.Cookie{
		Name:     auth.CookieName,
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(24 * time.Hour),
	})

	c.Request = httptest.NewRequest("GET", "/", nil)
	err = c.Request.Body.Close()
	if err != nil {
		log.Printf("Ошибка при закрытии чтения тела ответа: %v", err)
	}

	c.Request.AddCookie(&http.Cookie{
		Name:  auth.CookieName,
		Value: token,
	})

	userID, err := auth.GetUserID(c)
	require.NoError(t, err)
	assert.NotEmpty(t, userID)
}

// //Бенчмарки
func BenchmarkGenerateJWT(b *testing.B) {
	userID := auth.GenerateUserID()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := auth.GenerateJWT(userID)
		if err != nil {
			b.Error(err) // Обработка ошибки
		}
	}
}

func BenchmarkValidateJWT(b *testing.B) {
	userID := auth.GenerateUserID()
	token, err := auth.GenerateJWT(userID)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := auth.ValidateJWT(token)
		if err != nil {
			b.Error(err) // Обработка ошибки
		}
	}
}

func BenchmarkSetUserCookie(b *testing.B) {
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)

	userID := auth.GenerateUserID()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		auth.SetUserCookie(c, userID)
	}
}

func BenchmarkGetUserID(b *testing.B) {
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)

	userID := auth.GenerateUserID()
	auth.SetUserCookie(c, userID)

	c.Request = httptest.NewRequest(http.MethodGet, "/", nil)

	err := c.Request.Body.Close()
	if err != nil {
		log.Printf("Ошибка при закрытии чтения тела ответа: %v", err)
	}

	c.Request.AddCookie(&http.Cookie{
		Name:  auth.CookieName,
		Value: userID,
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := auth.GetUserID(c)
		if err != nil {
			b.Error(err) // Обработка ошибки
		}
	}
}
func BenchmarkGetUserIDByCookie(b *testing.B) {
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)

	userID := auth.GenerateUserID()
	auth.SetUserCookie(c, userID)

	c.Request = httptest.NewRequest(http.MethodGet, "/", nil)

	err := c.Request.Body.Close()
	if err != nil {
		log.Printf("Ошибка при закрытии чтения тела ответа: %v", err)
	}

	c.Request.AddCookie(&http.Cookie{
		Name:  auth.CookieName,
		Value: userID,
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := auth.GetUserIDByCookie(c)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkLogHeader(b *testing.B) {
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)

	c.Request = &http.Request{
		Header: http.Header{
			"Authorization": []string{"Bearer some_token"},
			"Content-Type":  []string{"application/json"},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		auth.LogHeader(c, "TestAction")
	}
}

func BenchmarkSetCookieUseInRequest(b *testing.B) {
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)

	token, err := auth.GenerateJWT(auth.GenerateUserID())
	if err != nil {
		b.Fatal(err)
	}

	c.Request = httptest.NewRequest(http.MethodGet, "/", nil)

	err = c.Request.Body.Close()
	if err != nil {
		log.Printf("Ошибка при закрытии чтения тела ответа: %v", err)
	}

	c.Request.AddCookie(&http.Cookie{
		Name:  auth.CookieName,
		Value: token,
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		auth.SetCookieUseInRequest(c)
	}
}
