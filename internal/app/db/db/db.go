package db

import (
	"github.com/Alandres998/url-shortner/internal/app/db/storage"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type DBStorage struct {
	db *sqlx.DB
}

var DB *sqlx.DB

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
		original_url TEXT NOT NULL,
		date_created TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		logger.Error("Не удалось создать таблицу",
			zap.String("Ошибка", err.Error()),
		)
		return nil, err
	}
	DB = db
	return &DBStorage{db: db}, nil
}

func (s *DBStorage) Set(shortURL, originalURL string) error {
	query := `
	INSERT INTO short_url (short_url, original_url)
	VALUES ($1, $2);`

	_, err := s.db.Exec(query, shortURL, originalURL)
	return err
}

func (s *DBStorage) Get(shortURL string) (string, error) {
	query := `
	SELECT original_url
	FROM short_url
	WHERE short_url = $1;`

	var originalURL string
	err := s.db.Get(&originalURL, query, shortURL)
	if err != nil {
		return "", err
	}
	return originalURL, nil
}
