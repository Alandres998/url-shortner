package storage

import (
	"context"
	"errors"
	"time"
)

// URLData представляет данные о коротком URL.
type URLData struct {
	ID          int       `json:"uuid" db:"id"`                   // Идентификатор записи
	ShortURL    string    `json:"short_url" db:"short_url"`       // Короткий URL
	OriginalURL string    `json:"original_url" db:"original_url"` // Оригинальный URL
	UserID      string    `json:"user_id" db:"user_id"`           // Идентификатор пользователя
	DateCreated time.Time `db:"date_created"`                     // Дата создания записи
	Deleted     bool      `json:"is_deleted" db:"is_deleted"`     // Флаг, указывающий, был ли URL удален
}

// Storage интерфейс определяет методы для работы с хранилищем URL.
type Storage interface {
	// Set сохраняет значение по ключу для указанного пользователя.
	Set(ctx context.Context, userID, key, value string) error

	// Get возвращает значение, связанное с указанным ключом.
	Get(ctx context.Context, key string) (string, error)

	// GetbyOriginURL возвращает данные URL по оригинальному URL.
	GetbyOriginURL(ctx context.Context, key string) (URLData, error)

	// GetUserURLs возвращает все короткие URL, связанные с указанным пользователем.
	GetUserURLs(ctx context.Context, userID string) ([]URLData, error)

	// DeleteUserURL удаляет указанные короткие URL, принадлежащие пользователю.
	DeleteUserURL(ctx context.Context, shortURLs []string, userID string) error

	// Ping метод для проверки доступности
	Ping(ctx context.Context) error
}

// ErrURLExists возвращается, когда URL уже существует в хранилище.
var ErrURLExists = errors.New("такой адрес уже есть")

// ErrURLDeleted возвращается, когда запрашивается удаленный URL.
var ErrURLDeleted = errors.New("URL был удален")

// Store представляет текущее хранилище URL.
var Store Storage

// MockStorage представляет собой заглушку для тестирования интерфейса Storage.
type MockStorage struct {
	// DeleteUserURLFunc - функция, используемая для подмены метода DeleteUserURL в тестах.
	DeleteUserURLFunc func(ctx context.Context, urls []string, userID string) error
}

// Get возвращает значение, связанное с указанным ключом.
func (m *MockStorage) Get(ctx context.Context, key string) (string, error) {
	panic("тест")
}

// GetUserURLs возвращает все короткие URL, связанные с указанным пользователем.
func (m *MockStorage) GetUserURLs(ctx context.Context, userID string) ([]URLData, error) {
	panic("тест")
}

// GetbyOriginURL возвращает данные URL по оригинальному URL.
func (m *MockStorage) GetbyOriginURL(ctx context.Context, key string) (URLData, error) {
	panic("тест")
}

// Ping проверяет доступность хранилища.
func (m *MockStorage) Ping(ctx context.Context) error {
	panic("тест")
}

// Set сохраняет значение по ключу для указанного пользователя.
func (m *MockStorage) Set(ctx context.Context, userID string, key string, value string) error {
	panic("тест")
}

// DeleteUserURL удаляет указанные короткие URL, принадлежащие пользователю.
func (m *MockStorage) DeleteUserURL(ctx context.Context, urls []string, userID string) error {
	if m.DeleteUserURLFunc != nil {
		return m.DeleteUserURLFunc(ctx, urls, userID)
	}
	return nil
}
