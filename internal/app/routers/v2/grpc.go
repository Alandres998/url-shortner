package v2

import (
	"context"
	"errors"
	"log"

	"github.com/gofrs/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/Alandres998/url-shortner/internal/app/db/storage"
	"github.com/Alandres998/url-shortner/internal/app/proto"
	webservices "github.com/Alandres998/url-shortner/internal/app/webServices"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const userIDKey = "user_id"

type URLShortenerServer struct {
	proto.UnimplementedURLShortenerServiceServer
}

// EnsureUserIDInterceptor проверяет наличие user_id в metadata или создает новый.
func EnsureUserIDInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, errors.New("metadata отсутствует")
		}

		userIDs := md.Get(userIDKey)
		if len(userIDs) == 0 {
			// Генерация нового user_id
			newUserID := uuid.Must(uuid.NewV4()).String()
			md.Append(userIDKey, newUserID)
			ctx = metadata.NewIncomingContext(ctx, md)
		}

		return handler(ctx, req)
	}
}

// Извлекает  user_id из metadata
func GetUserIDFromContext(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("metadata отсутствует")
	}

	userIDs := md.Get(userIDKey)
	if len(userIDs) == 0 {
		return "", errors.New("user_id отсутствует")
	}

	return userIDs[0], nil
}

func (s *URLShortenerServer) CreateShortURL(ctx context.Context, req *proto.CreateShortURLRequest) (*proto.CreateShortURLResponse, error) {
	logger, errLog := zap.NewProduction()
	if errLog != nil {
		log.Fatalf("Не смог иницировать логгер")

	}
	defer func() {
		if errLoger := logger.Sync(); errLoger != nil {
			logger.Error("Проблемы при закрытии логера",
				zap.String("Не смог закрыть логгер", errLoger.Error()),
			)
		}
	}()

	userID, err := GetUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	shortURL, err := webservices.ShorterGeneral(ctx, userID, req.OriginalUrl)
	if err != nil {
		if errors.Is(err, storage.ErrURLExists) {
			return &proto.CreateShortURLResponse{
				ShortUrl: shortURL,
			}, nil
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.CreateShortURLResponse{
		ShortUrl: shortURL,
	}, nil
}

// GetOriginalURL Получить полную ссылку
func (s *URLShortenerServer) GetOriginalURL(ctx context.Context, req *proto.GetOriginalURLRequest) (*proto.GetOriginalURLResponse, error) {
	originalURL, err := webservices.Fuller(ctx, req.ShortUrl)
	if err != nil {
		if err == storage.ErrURLDeleted {
			return nil, status.Error(codes.NotFound, "URL was deleted")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.GetOriginalURLResponse{
		OriginalUrl: originalURL,
	}, nil
}

// GetOriginalURL Получить ссылки пользователя
func (s *URLShortenerServer) GetUserURLs(ctx context.Context, req *proto.GetUserURLsRequest) (*proto.GetUserURLsResponse, error) {
	userID := req.UserId
	if userID == "" {
		return nil, status.Error(codes.InvalidArgument, "user_id не может быть пустым")
	}

	urls, err := storage.Store.GetUserURLs(context.Background(), userID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if len(urls) == 0 {
		return nil, status.Error(codes.NotFound, "URLs не найдены")
	}

	response := &proto.GetUserURLsResponse{}
	for _, url := range urls {
		response.Urls = append(response.Urls, &proto.UserURL{
			ShortUrl:    url.ShortURL,
			OriginalUrl: url.OriginalURL,
		})
	}

	return response, nil
}

// DeleteUserURLs удалить ссылку пользваоетля
func (s *URLShortenerServer) DeleteUserURLs(ctx context.Context, req *proto.DeleteUserURLsRequest) (*proto.DeleteUserURLsResponse, error) {
	if len(req.ShortUrls) == 0 {
		return nil, status.Error(codes.InvalidArgument, "список short_urls не может быть пустым")
	}

	userID, err := GetUserIDFromContext(ctx)
	if err != nil || userID == "" {
		return nil, status.Error(codes.Unauthenticated, "не удалось определить user_id")
	}

	// Вызов бизнес-логики для удаления ссылок
	err = storage.Store.DeleteUserURL(ctx, req.ShortUrls, userID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.DeleteUserURLsResponse{Message: "URLs успешно удалены"}, nil
}
