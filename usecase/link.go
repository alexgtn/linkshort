package usecase

import (
	"context"
	"fmt"

	"github.com/alexgtn/go-linkshort/domain/link"
	pb "github.com/alexgtn/go-linkshort/proto"
)

type linkRepo interface {
	Create(ctx context.Context, long string) (*link.Link, error)
	GetOne(ctx context.Context, long string) (*link.Link, error)
}

type service struct {
	pb.UnimplementedLinkshortServiceServer
	linkRepo linkRepo
	baseURL  string
}

func NewLinkService(r linkRepo, baseURL string) *service {
	return &service{
		linkRepo: r,
		baseURL:  baseURL,
	}
}

func (s *service) Redirect(context.Context, *pb.RedirectRequest) (*pb.RedirectReply, error) {
	return &pb.RedirectReply{
		LongUri: "https://jsonplaceholder.typicode.com/albums",
	}, nil
}

func (s *service) Create(context.Context, *pb.CreateLinkRequest) (*pb.CreateLinkReply, error) {
	return &pb.CreateLinkReply{
		ShortUri: fmt.Sprintf("%s/qqq", s.baseURL),
	}, nil
}
