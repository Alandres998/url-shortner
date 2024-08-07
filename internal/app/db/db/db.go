package db

import (
	"context"
	"time"

	"github.com/Alandres998/url-shortner/internal/app/db/storage"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"go.uber.org/zap"
)

type DBStorage struct {
	db *sqlx.DB
}

func NewDBStorage(dsn string) (storage.Storage, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	defer logger.Sync()

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		logger.Error("Проблемы при подключении к БД",
			zap.String("Не смог подключиться к БД", err.Error()),
		)
		return nil, err
	}

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS short_url (
		id SERIAL PRIMARY KEY,
		short_url TEXT NOT NULL,
		original_url TEXT NOT NULL UNIQUE,
		date_created TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);`

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err = db.ExecContext(ctx, createTableQuery)
	if err != nil {
		logger.Error("Не удалось создать таблицу",
			zap.String("Ошибка", err.Error()),
		)
		return nil, err
	}

	return &DBStorage{db: db}, nil
}

func (s *DBStorage) Set(ctx context.Context, shortURL, originalURL string) error {
	query := `
	INSERT INTO short_url (short_url, original_url)
	VALUES ($1, $2);`

	_, err := s.db.ExecContext(ctx, query, shortURL, originalURL)
	if err != nil && isUniqueViolation(err) {
		return storage.ErrURLExists
	}
	return nil
}

func (s *DBStorage) Get(ctx context.Context, shortURL string) (string, error) {
	query := `
	SELECT original_url
	FROM short_url
	WHERE short_url = $1;`

	var originalURL string
	err := s.db.GetContext(ctx, &originalURL, query, shortURL)
	if err != nil {
		return "", err
	}
	return originalURL, nil
}

func (s *DBStorage) GetbyOriginURL(ctx context.Context, originalURL string) (storage.URLData, error) {
	query := `
	SELECT id, short_url, original_url, date_created
	FROM short_url
	WHERE original_url = $1;`

	var urlData storage.URLData
	err := s.db.GetContext(ctx, &urlData, query, originalURL)
	if err != nil {
		return urlData, err
	}
	return urlData, nil
}

func isUniqueViolation(err error) bool {
	if pgErr, ok := err.(*pq.Error); ok {
		if pgErr.Code == "23505" {
			return true
		}
	}
	return false
}

func (s *DBStorage) Ping(ctx context.Context) error {
	return s.db.PingContext(ctx)
}
