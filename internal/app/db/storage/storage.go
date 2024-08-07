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
	DateCreated time.Time `db:"date_created"`
}

type Storage interface {
	Set(ctx context.Context, key string, value string) error
	Get(ctx context.Context, key string) (string, error)
	GetbyOriginURL(ctx context.Context, key string) (URLData, error)
	Ping(ctx context.Context) error
}

var ErrURLExists = errors.New("такой адрес уже есть")

var Store Storage
