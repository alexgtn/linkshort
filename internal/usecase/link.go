package usecase

import (
	"context"
	"fmt"
	"strings"

	errors2 "github.com/pkg/errors"

	"github.com/alexgtn/go-linkshort/internal/domain/link"
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
func (s *service) Redirect(ctx context.Context, shortPath string) (string, error) {
	short := strings.TrimSpace(shortPath)

	existingLink, err := s.linkRepo.GetOneByShortPath(ctx, short)
	if err != nil {
		return "", errRedirect(err, short)
	}
	// return existing
	return existingLink.LongURL(), nil
}

// CreateLink creates a short link (if not exists), otherwise returns existing link
func (s *service) CreateLink(ctx context.Context, longURL string) (string, error) {
	existingLink, err := s.linkRepo.GetOneByLongURL(ctx, longURL)
	if err != nil {
		// create link
		newLink, err := s.linkRepo.Create(ctx, longURL)
		if err != nil {
			// try re-fetch existing
			// mostly in highly concurrent scenarios
			existingLink, err := s.linkRepo.GetOneByLongURL(ctx, longURL)
			if err != nil {
				return "", errCreateLink(err, longURL)
			}

			// return new link
			return shortURI(s.baseURL, existingLink.ShortPath()), nil
		}

		// set short path
		_, err = s.linkRepo.SetShortPath(ctx, newLink.ID(), newLink.ShortPath())
		if err != nil {
			return "", errCreateLink(err, longURL)
		}

		// return new link
		return shortURI(s.baseURL, newLink.ShortPath()), nil
	}
	// return existing
	return shortURI(s.baseURL, existingLink.ShortPath()), nil
}

func shortURI(baseURL string, path string) string {
	return fmt.Sprintf("%s/%s", baseURL, path)
}
