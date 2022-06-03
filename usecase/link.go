package usecase

import (
	"context"
	"fmt"
	"strings"

	errors2 "github.com/pkg/errors"

	"github.com/alexgtn/go-linkshort/domain/link"
	pb "github.com/alexgtn/go-linkshort/proto"
)

type linkRepo interface {
	Create(ctx context.Context, long string) (*link.Link, error)
	GetOne(ctx context.Context, long string) (*link.Link, error)
}

var errCreateLink = func(err error, link string) error {
	return errors2.Wrapf(err, fmt.Sprintf("error creating link %s", link))
}

var errRedirect = func(err error, link string) error {
	return errors2.Wrapf(err, fmt.Sprintf("error redirecting to link %s", link))
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

func (s *service) Redirect(ctx context.Context, r *pb.RedirectRequest) (*pb.RedirectReply, error) {
	long := strings.TrimSpace(r.ShortPath)

	err := r.ValidateAll()
	if err != nil {
		return nil, errRedirect(err, long)
	}

	existingLink, err := s.linkRepo.GetOne(ctx, long)
	if err != nil {
		return nil, errRedirect(err, long)
	}
	// return existing
	return &pb.RedirectReply{
		LongUri: existingLink.Long(),
	}, nil
}

func (s *service) Create(ctx context.Context, r *pb.CreateLinkRequest) (*pb.CreateLinkReply, error) {
	long := strings.TrimSpace(r.LongUri)

	err := r.ValidateAll()
	if err != nil {
		return nil, errCreateLink(err, long)
	}

	existingLink, err := s.linkRepo.GetOne(ctx, long)
	if err != nil {
		// create link
		newLink, err := s.linkRepo.Create(ctx, long)
		if err != nil {
			return nil, errCreateLink(err, long)
		}

		// return new link
		return &pb.CreateLinkReply{
			ShortUri: shortURI(s.baseURL, newLink.Short()),
		}, nil
	}
	// return existing
	return &pb.CreateLinkReply{
		ShortUri: shortURI(s.baseURL, existingLink.Short()),
	}, nil
}

func shortURI(baseURL string, path string) string {
	return fmt.Sprintf("%s/%s", baseURL, path)
}
