package auth

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/Alandres998/url-shortner/internal/app/service/logger"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

const CookieName = "user_id"
const secretKey = "kFHrlqA0"

type Authenticator interface {
	GetUserIDByCookie(c *gin.Context) (string, error)
}

// AuthService предоставляет методы для аутентификации пользователей.
type AuthService struct{}

// GetUserIDByCookie получает id пользователя из cookie.
func (a *AuthService) GetUserIDByCookie(c *gin.Context) (string, error) {
	return GetUserIDByCookie(c)
}

// GenerateUserID создает новый уникальный id пользователя.
func GenerateUserID() string {
	return uuid.Must(uuid.NewV4()).String()
}

// GenerateJWT создает новый JWT на основе id пользователя.
func GenerateJWT(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

// ValidateJWT проверяет действительность JWT и возвращает токен, если он валиден.
func ValidateJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("не прошла валидация")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if _, ok := claims[CookieName].(string); !ok {
			return nil, errors.New("нет ключа с id пользователя")
		}
	}
	return token, nil
}

// SetUserCookie устанавливает cookie с JWT для указанного id пользователя.
func SetUserCookie(c *gin.Context, userID string) {
	jwt, err := GenerateJWT(GenerateUserID())
	if err != nil {
		logger.LoginInfo("Не смог сгенирорвать токен", err.Error())
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     CookieName,
		Value:    jwt,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(24 * time.Hour),
	})
	c.Set(CookieName, jwt)
}

// GetUserID получает id пользователя из cookie или контекста Gin.
func GetUserID(c *gin.Context) (string, error) {
	cookie, err := c.Cookie(CookieName)
	if err != nil {
		jwt, exists := c.Get(CookieName)
		if exists {
			if jwtString, ok := jwt.(string); ok {
				return jwtString, nil
			}
		}
		return "", errors.New("нет ключа в куках")
	}

	return cookie, nil
}

// GetUserIDByCookie получает id пользователя из cookie.
func GetUserIDByCookie(c *gin.Context) (string, error) {
	cookie, err := c.Cookie(CookieName)
	if err != nil {
		return "", errors.New("нет ключа в куках")
	}
	return cookie, nil
}

// LogHeader логирует все заголовки HTTP-запроса.
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

// SetCookieUseInRequest устанавливает cookie для текущего запроса, если она отсутствует или недействительна.
func SetCookieUseInRequest(c *gin.Context) {
	token, err := c.Cookie(CookieName)

	//Кейс токен отсутствует или ошибка генерирует новый
	if token == "" || err != nil {
		token, err = GenerateJWT(GenerateUserID())
		if err != nil {
			logger.LoginInfo("Не смог сгенирорвать токен", err.Error())
			return
		}
		SetUserCookie(c, token)
		logger.LoginInfo("Установлен новый токен", token)
		return
	}

	if token != "" {
		_, err = ValidateJWT(token)
		if err != nil {
			token, err = GenerateJWT(GenerateUserID())
			if err != nil {
				logger.LoginInfo("Не смог сгенирорвать токен", err.Error())
				return
			}
			logger.LoginInfo("Установлен токен (не прошел валидацию)", token)
			SetUserCookie(c, token)
		}
	}
}
