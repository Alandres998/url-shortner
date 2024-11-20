package grpc

import (
	"context"
	"errors"

	"github.com/Alandres998/url-shortner/internal/app/db/storage"
	"github.com/Alandres998/url-shortner/internal/app/proto"
	webservices "github.com/Alandres998/url-shortner/internal/app/webServices"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type URLShortenerServer struct {
	proto.UnimplementedURLShortenerServiceServer
}

func (s *URLShortenerServer) ShortenURL(ctx context.Context, req *proto.CreateShortURLRequest) (*proto.CreateShortURLResponse, error) {

	shortURL, err := webservices.ShorterGeneral(ctx, req.UserId, req.OriginalUrl)
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
