package grpc

import (
	"context"
	"fmt"
	"strings"

	errors2 "github.com/pkg/errors"

	pb "github.com/alexgtn/go-linkshort/proto"
)

type linkService interface {
	Redirect(ctx context.Context, shortPath string) (string, error)
	CreateLink(ctx context.Context, longURL string) (string, error)
}

var errCreateLink = func(err error, link string) error {
	return errors2.Wrapf(err, fmt.Sprintf("error creating link %s", link))
}

var errRedirect = func(err error, link string) error {
	return errors2.Wrapf(err, fmt.Sprintf("error redirecting to link %s", link))
}

type delivery struct {
	pb.UnimplementedLinkshortServiceServer
	linkService linkService
}

func NewLinkDeliveryGrpc(r linkService) *delivery {
	return &delivery{
		linkService: r,
	}
}

// Redirect returns the long URL provided a short path
func (d *delivery) Redirect(ctx context.Context, r *pb.RedirectRequest) (*pb.RedirectReply, error) {
	short := strings.TrimSpace(r.ShortPath)

	err := r.ValidateAll()
	if err != nil {
		return nil, errRedirect(err, short)
	}

	existingLink, err := d.linkService.Redirect(ctx, short)
	if err != nil {
		return nil, errRedirect(err, short)
	}
	// return existing
	return &pb.RedirectReply{
		LongUri: existingLink,
	}, nil
}

// CreateLink creates a short link (if not exists), otherwise returns existing link
func (d *delivery) CreateLink(ctx context.Context, r *pb.CreateLinkRequest) (*pb.CreateLinkReply, error) {
	err := r.ValidateAll()
	if err != nil {
		return nil, errCreateLink(err, r.LongUri)
	}

	shortURI, err := d.linkService.CreateLink(ctx, r.LongUri)
	if err != nil {
		return nil, errCreateLink(err, r.LongUri)
	}

	// return existing
	return &pb.CreateLinkReply{
		ShortUri: shortURI,
	}, nil
}
