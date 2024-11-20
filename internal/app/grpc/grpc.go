package grpc

import (
	"context"
	"log"

	pb "github.com/Alandres998/url-shortner/internal/app/proto"
	"github.com/Alandres998/url-shortner/internal/app/repository"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type ShortenerServer struct {
	pb.UnimplementedShortenerServiceServer
	repo repository.URLRepository // Предполагается, что у вас есть репозиторий для работы с URL
}

func (s *ShortenerServer) ShortenURL(ctx context.Context, req *pb.ShortenRequest) (*pb.ShortenResponse, error) {
	// Логика для сокращения URL
	shortURL := uuid.New().String() // Генерация случайного идентификатора для сокращенного URL
	err := s.repo.SaveURL(req.OriginalUrl, shortURL)
	if err != nil {
		return &pb.ShortenResponse{Error: err.Error()}, nil
	}
	return &pb.ShortenResponse{ShortUrl: shortURL}, nil
}

func (s *ShortenerServer) GetFullURL(ctx context.Context, req *pb.GetFullRequest) (*pb.GetFullResponse, error) {
	// Логика для получения оригинального URL по короткому
	originalURL, err := s.repo.GetOriginalURL(req.ShortUrl)
	if err != nil {
		return &pb.GetFullResponse{Error: err.Error()}, nil
	}
	return &pb.GetFullResponse{OriginalUrl: originalURL}, nil
}

func (s *ShortenerServer) GetAllShortURLByCookie(ctx context.Context, req *pb.EmptyRequest) (*pb.GetAllShortURLResponse, error) {
	// Логика для получения всех URL для пользователя по cookie
	urls, err := s.repo.GetAllURLsByUser(ctx)
	if err != nil {
		return nil, err
	}
	responseURLs := make([]*pb.ShortUserResponse, len(urls))
	for i, url := range urls {
		responseURLs[i] = &pb.ShortUserResponse{
			ShortUrl:    url.ShortURL,
			OriginalUrl: url.OriginalURL,
		}
	}
	return &pb.GetAllShortURLResponse{Urls: responseURLs}, nil
}

func (s *ShortenerServer) DeleteShortURL(ctx context.Context, req *pb.DeleteShortURLRequest) (*pb.EmptyResponse, error) {
	// Логика для удаления URL
	for _, shortURL := range req.ShortUrls {
		err := s.repo.DeleteURL(shortURL)
		if err != nil {
			return nil, err
		}
	}
	return &pb.EmptyResponse{}, nil
}

func (s *ShortenerServer) Ping(ctx context.Context, req *pb.EmptyRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{Status: "OK"}, nil
}

func NewShortenerServer(repo repository.URLRepository) *ShortenerServer {
	return &ShortenerServer{repo: repo}
}

func StartGRPCServer() {
	server := grpc.NewServer()
	repo := repository.NewURLRepository() // Предполагается, что есть репозиторий для работы с базой данных
	shortenerServer := NewShortenerServer(repo)
	pb.RegisterShortenerServiceServer(server, shortenerServer)

	log.Println("Starting gRPC server on port 50051...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
