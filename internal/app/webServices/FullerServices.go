package webservices

import (
	"context"
	"errors"

	"github.com/Alandres998/url-shortner/internal/app/db/storage"
)

// Fuller для возврата полной строки
func Fuller(ctx context.Context, id string) (string, error) {
	urlOriginal, err := storage.Store.Get(ctx, id)
	if err != nil {
		if errors.Is(err, storage.ErrURLDeleted) {
			return "", storage.ErrURLDeleted
		}
		return "", err
	}
	return urlOriginal, nil
}
