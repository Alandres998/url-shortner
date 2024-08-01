package storage

import (
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
	Set(key string, value string) error
	Get(key string) (string, error)
	GetbyOriginURL(key string) (URLData, error)
}

var ErrURLExists = errors.New("такой адрес уже есть")

var Store Storage
