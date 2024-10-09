package webservices_test

import (
	"context"
	"testing"

	"net/http/httptest"

	"github.com/Alandres998/url-shortner/internal/app/db/storage"
	"github.com/Alandres998/url-shortner/internal/app/webservices"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

// MockStorage - мок для тестирования
type MockStorage struct {
	mock.Mock
}

func (m *MockStorage) Get(ctx context.Context, id string) (string, error) {
	args := m.Called(ctx, id)
	return args.String(0), args.Error(1)
}

func (m *MockStorage) DeleteUserURL(ctx context.Context, userID, urlID string) error {
	args := m.Called(ctx, userID, urlID)
	return args.Error(0)
}

func (m *MockStorage) GetUserURLs(ctx context.Context, userID string) ([]storage.Store, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([]storage.Store), args.Error(1)
}

// Добавьте другие методы интерфейса по мере необходимости

func BenchmarkFuller(b *testing.B) {
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)

	// Создаем мок хранилища
	mockStorage := new(MockStorage)

	// Настраиваем ожидания
	id := "test-id"
	expectedURL := "http://original-url.com"
	mockStorage.On("Get", mock.Anything, id).Return(expectedURL, nil)

	// Установите ваш мок в глобальную переменную или передайте в функцию
	// Например, если у вас есть функция, которая принимает хранилище:
	storage.SetStore(mockStorage) // Реализуйте SetStore в пакете storage

	// Устанавливаем параметр id в контексте
	c.Params = gin.Params{gin.Param{Key: "id", Value: id}}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result, err := webservices.Fuller(c)
		if err != nil {
			b.Error(err) // Обработка ошибки
		}
		if result != expectedURL {
			b.Errorf("expected %s, got %s", expectedURL, result)
		}
	}

	// Проверяем ожидания
	mockStorage.AssertExpectations(b)
}
