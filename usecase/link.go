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
	GetOneByShortPath(ctx context.Context, short string) (*link.Link, error)
	GetOneByLongURL(ctx context.Context, long string) (*link.Link, error)
	SetShortPath(ctx context.Context, id int, path string) (*link.Link, error)
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

// Redirect returns the long URL provided a short path
func (s *service) Redirect(ctx context.Context, r *pb.RedirectRequest) (*pb.RedirectReply, error) {
	short := strings.TrimSpace(r.ShortPath)

	err := r.ValidateAll()
	if err != nil {
		return nil, errRedirect(err, short)
	}

	existingLink, err := s.linkRepo.GetOneByShortPath(ctx, short)
	if err != nil {
		return nil, errRedirect(err, short)
	}
	// return existing
	return &pb.RedirectReply{
		LongUri: existingLink.LongURL(),
	}, nil
}

// Create creates a short link (if not exists), otherwise returns existing link
func (s *service) Create(ctx context.Context, r *pb.CreateLinkRequest) (*pb.CreateLinkReply, error) {
	err := r.ValidateAll()
	if err != nil {
		return nil, errCreateLink(err, r.LongUri)
	}

	existingLink, err := s.linkRepo.GetOneByLongURL(ctx, r.LongUri)
	if err != nil {
		// create link
		newLink, err := s.linkRepo.Create(ctx, r.LongUri)
		if err != nil {
			return nil, errCreateLink(err, r.LongUri)
		}

		// set short path
		_, err = s.linkRepo.SetShortPath(ctx, newLink.ID(), newLink.ShortPath())
		if err != nil {
			return nil, errCreateLink(err, r.LongUri)
		}

		// return new link
		return &pb.CreateLinkReply{
			ShortUri: shortURI(s.baseURL, newLink.ShortPath()),
		}, nil
	}
	// return existing
	return &pb.CreateLinkReply{
		ShortUri: shortURI(s.baseURL, existingLink.ShortPath()),
	}, nil
}

func shortURI(baseURL string, path string) string {
	return fmt.Sprintf("%s/%s", baseURL, path)
}
