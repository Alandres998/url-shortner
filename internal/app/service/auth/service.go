package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"
)

const CookieName = "user_id"
const secretKey = "kFHrlqA0"

func GenerateUserID() string {
	uid, _ := uuid.NewV4()
	return uid.String()
}

func SignCookie(userID string) string {
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(userID))
	return userID + ":" + hex.EncodeToString(h.Sum(nil))
}

func ValidateCookie(cookie string) bool {
	parts := strings.Split(cookie, ":")
	if len(parts) != 2 {
		return false
	}
	userID, signature := parts[0], parts[1]
	expectedSignature := SignCookie(userID)
	return hmac.Equal([]byte(signature), []byte(strings.Split(expectedSignature, ":")[1]))
}

func SetUserCookie(c *gin.Context, userID string) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     CookieName,
		Value:    SignCookie(userID),
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(24 * time.Hour),
	})
}

func GetUserID(c *gin.Context) (string, error) {
	cookie, err := c.Cookie(CookieName)

	if err != nil {
		InfoCookie(c, "Пытаюсь взять пользователя из БД")
		return "", errors.New("нет ключа в куках")
	}

	return cookie, nil
}

func InfoCookie(c *gin.Context, action string) {
	logger, errLog := zap.NewProduction()
	defer logger.Sync()
	if errLog != nil {
		log.Fatalf("Не смог иницировать логгер")
	}

	statusCode := c.Writer.Status()
	contentLength := c.Writer.Size()
	logger.Info("Request CookieInfo",
		zap.String("url", c.Request.RequestURI),
		zap.String("method", c.Request.Method),
		zap.Int("status_code", statusCode),
		zap.Int("content_length", contentLength),
	)

	cookies := c.Request.Cookies()
	logger.Info("Cookie Get",
		zap.String(action, "пытаюсь получить куки"),
	)
	for _, cookie := range cookies {
		logger.Info("Cookie Get",
			zap.String("Name", cookie.Name),
			zap.String("Value", cookie.Value),
		)
	}
}
