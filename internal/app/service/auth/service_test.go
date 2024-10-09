package auth_test

import (
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

// Тестируем SetUserCookie
func TestSetUserCookie(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest("GET", "/", nil)

	auth.SetUserCookie(c, "user123")

	cookies := w.Result().Cookies()
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
	c.Request.AddCookie(&http.Cookie{
		Name:  auth.CookieName,
		Value: token,
	})
	userID, err := auth.GetUserID(c)
	require.NoError(t, err)
	assert.NotEmpty(t, userID)
}
