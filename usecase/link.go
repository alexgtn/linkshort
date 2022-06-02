package usecase

import (
	"context"
	"fmt"

	"github.com/alexgtn/go-linkshort/domain/user"
	user_repo "github.com/alexgtn/go-linkshort/infra/repository/user"
	pb "github.com/alexgtn/go-linkshort/proto"
)

type linkRepo interface {
	GetByID(ctx context.Context, id int) (*user.User, error)
	Create(ctx context.Context, age int, name string) (*user.User, error)
	Update(ctx context.Context, id int, opts ...user_repo.Option) (*user.User, error)
}

type service struct {
	pb.UnimplementedLinkshortServiceServer
	userRepo linkRepo
	baseURL  string
}

func NewUserService(r linkRepo, baseURL string) *service {
	return &service{
		userRepo: r,
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
