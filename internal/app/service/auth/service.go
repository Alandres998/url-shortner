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

	"github.com/Alandres998/url-shortner/internal/app/service/logger"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"
)

const CookieName = "user_id"
const secretKey = "kFHrlqA0"

var UserIDTemp = ""

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
	logger.LoginInfo("Берем глобальную переменную", UserIDTemp)
	if err != nil {
		if UserIDTemp != "" {
			return UserIDTemp, nil
		} else {
			return "", errors.New("нет ключа в куках")
		}

	}

	return cookie, nil
}

func LogHeader(c *gin.Context, action string) {
	// Логируем все заголовки
	logger, errLog := zap.NewProduction()
	defer logger.Sync()
	if errLog != nil {
		log.Fatalf("Не смог иницировать логгер")
	}
	headers := c.Request.Header
	for key, values := range headers {
		for _, value := range values {
			logger.Info(action,
				zap.String("header", key),
				zap.String("value", value),
			)
		}
	}
}

func SetCookieUseInRequest(c *gin.Context) {
	cookie, err := c.Cookie(CookieName)
	if err != nil || !ValidateCookie(cookie) {
		userID := GenerateUserID()
		SetUserCookie(c, userID)
		c.Set(CookieName, userID)
		c.SetCookie(CookieName, cookie, 3600, "/", "localhost", false, true)
		//UserIDTemp = userID
	} else {
		c.Set(CookieName, cookie)
		c.SetCookie(CookieName, cookie, 3600, "/", "localhost", false, true)
		//UserIDTemp = cookie
	}
	logger.LoginInfo("Устанвка переменной с куки", UserIDTemp)
}
