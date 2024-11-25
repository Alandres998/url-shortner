package db

import (
	"context"
	"time"

	"github.com/Alandres998/url-shortner/internal/app/db/storage"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"go.uber.org/zap"
)

// DBStorage для работы с базой данных.
type DBStorage struct {
	db *sqlx.DB
}

// NewDBStorage создает новое соединение с базой данных и инициализирует таблицу short_url.
// Возвращает экземпляр Storage и ошибку, если подключение не удалось.
func NewDBStorage(dsn string) (storage.Storage, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	defer func() {
		if errLoger := logger.Sync(); errLoger != nil {
			logger.Error("Проблемы при закрытии логера",
				zap.String("Не смог закрыть логгер", errLoger.Error()),
			)
		}
	}()

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
		user_id TEXT,
		date_created TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		is_deleted boolean DEFAULT FALSE
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

// Set сохраняет новую пару (короткий URL, оригинальный URL) в базе данных для указанного пользователя.
// Возвращает ошибку, если URL уже существует.
func (s *DBStorage) Set(ctx context.Context, userID, shortURL, originalURL string) error {
	query := `
	INSERT INTO short_url (short_url, original_url, user_id)
	VALUES ($1, $2, $3);`

	_, err := s.db.ExecContext(ctx, query, shortURL, originalURL, userID)
	if err != nil && isUniqueViolation(err) {
		return storage.ErrURLExists
	}
	return nil
}

// Get получает оригинальный URL по короткому URL.
// Возвращает оригинальный URL и ошибку, если короткий URL не найден или был удален.
func (s *DBStorage) Get(ctx context.Context, shortURL string) (string, error) {
	query := `
	SELECT id, short_url, original_url, user_id, date_created, is_deleted
	FROM short_url
	WHERE short_url = $1;`

	var urlData storage.URLData
	err := s.db.GetContext(ctx, &urlData, query, shortURL)
	if err != nil {
		return "", err
	}

	if urlData.Deleted {
		return urlData.OriginalURL, storage.ErrURLDeleted
	}

	return urlData.OriginalURL, nil
}

// GetbyOriginURL получает данные URL по оригинальному URL.
// Возвращает структуру URLData и ошибку, если оригинальный URL не найден.
func (s *DBStorage) GetbyOriginURL(ctx context.Context, originalURL string) (storage.URLData, error) {
	query := `
	SELECT id, short_url, original_url, user_id, date_created, is_deleted
	FROM short_url
	WHERE original_url = $1;`

	var urlData storage.URLData
	err := s.db.GetContext(ctx, &urlData, query, originalURL)
	if err != nil {
		return urlData, err
	}
	return urlData, nil
}

// isUniqueViolation проверяет, является ли ошибка нарушением уникальности в базе данных.
// Возвращает true, если это нарушение уникальности, иначе false.
func isUniqueViolation(err error) bool {
	if pgErr, ok := err.(*pq.Error); ok {
		if pgErr.Code == "23505" {
			return true
		}
	}
	return false
}

// Ping проверяет доступность базы данных.
// Возвращает ошибку, если база данных недоступна.
func (s *DBStorage) Ping(ctx context.Context) error {
	return s.db.PingContext(ctx)
}

// GetUserURLs получает список всех сокращенных URL для указанного пользователя.
// Возвращает срез URLData и ошибку, если пользователь не найден или произошла другая ошибка.
func (s *DBStorage) GetUserURLs(ctx context.Context, userID string) ([]storage.URLData, error) {
	query := `
	SELECT id, short_url, original_url, user_id, date_created
	FROM short_url
	WHERE user_id = $1;`

	var urls []storage.URLData
	err := s.db.SelectContext(ctx, &urls, query, userID)
	if err != nil {
		return nil, err
	}
	return urls, nil
}

// DeleteUserURL удаляет указанные короткие URL для указанного пользователя, устанавливая флаг is_deleted.
// Возвращает ошибку, если что-то пошло не так.
func (s *DBStorage) DeleteUserURL(ctx context.Context, shortURLs []string, userID string) error {
	query := `
	UPDATE short_url
	SET is_deleted = TRUE
	WHERE short_url = ANY($1) AND user_id = $2;`

	_, err := s.db.ExecContext(ctx, query, pq.Array(shortURLs), userID)
	if err != nil {
		return err
	}

	return nil
}

// GetStatistics получить иннформацию о количестве сокращенных ссылок и уникальных пользователях
func (s *DBStorage) GetStatistics(ctx context.Context) (int, int, error) {
	urlCountQuery := `SELECT COUNT(*) FROM short_url WHERE NOT is_deleted;`
	userCountQuery := `SELECT COUNT(DISTINCT user_id) FROM short_url WHERE user_id IS NOT NULL;`

	var urlCount, userCount int

	if err := s.db.GetContext(ctx, &urlCount, urlCountQuery); err != nil {
		return 0, 0, err
	}

	if err := s.db.GetContext(ctx, &userCount, userCountQuery); err != nil {
		return 0, 0, err
	}

	return urlCount, userCount, nil
}
