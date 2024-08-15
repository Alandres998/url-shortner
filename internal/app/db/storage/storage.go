package storage

import (
	"context"
	"errors"
	"time"
)

type URLData struct {
	ID          int       `json:"uuid" db:"id"`
	ShortURL    string    `json:"short_url" db:"short_url"`
	OriginalURL string    `json:"original_url" db:"original_url"`
	UserID      string    `json:"user_id" db:"user_id"`
	DateCreated time.Time `db:"date_created"`
	Deleted     bool      `json:"is_deleted" db:"is_deleted"`
}

type Storage interface {
	Set(ctx context.Context, userID, key, value string) error
	Get(ctx context.Context, key string) (string, error)
	GetbyOriginURL(ctx context.Context, key string) (URLData, error)
	GetUserURLs(ctx context.Context, userID string) ([]URLData, error)
	DeleteUserURL(ctx context.Context, shortURLs []string, userID string) error
	Ping(ctx context.Context) error
}

var ErrURLExists = errors.New("такой адрес уже есть")
var ErrURLDeleted = errors.New("URL был удален")

var Store Storage
