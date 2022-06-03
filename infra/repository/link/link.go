package link

import (
	"context"

	"github.com/pkg/errors"

	errors2 "github.com/alexgtn/go-linkshort/domain/errors"
	"github.com/alexgtn/go-linkshort/domain/link"
	"github.com/alexgtn/go-linkshort/ent"
	link2 "github.com/alexgtn/go-linkshort/ent/link"
)

// linkRepo manages a sql-based link repo
type linkRepo struct {
	client *ent.Client
}

func NewLinkRepo(c *ent.Client) *linkRepo {
	return &linkRepo{c}
}

func (r *linkRepo) GetOneByShortPath(ctx context.Context, short string) (*link.Link, error) {
	l, err := r.client.Link.
		Query().
		Where(link2.ShortPath(short)).
		Only(ctx)
	if err != nil {
		switch err.(type) {
		case *ent.NotFoundError:
			return nil, errors2.WrapErrNotFound(err)
		case *ent.NotSingularError:
			return nil, errors2.WrapErrNotSingular(err)
		default:
			return nil, errors.Wrapf(err, "could not get link by short path %s", short)
		}
	}

	dto, err := link.NewLink(l.ID, l.LongURI, l.CreatedAt)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create link %s", l.LongURI)
	}

	return dto, nil
}

func (r *linkRepo) GetOneByLongURL(ctx context.Context, long string) (*link.Link, error) {
	l, err := r.client.Link.
		Query().
		Where(link2.LongURI(long)).
		Only(ctx)
	if err != nil {
		switch err.(type) {
		case *ent.NotFoundError:
			return nil, errors2.WrapErrNotFound(err)
		case *ent.NotSingularError:
			return nil, errors2.WrapErrNotSingular(err)
		default:
			return nil, errors.Wrapf(err, "could not get link %s", long)
		}
	}

	dto, err := link.NewLink(l.ID, l.LongURI, l.CreatedAt)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create link %s", l.LongURI)
	}

	return dto, nil
}

func (r *linkRepo) SetShortPath(ctx context.Context, id int, path string) (*link.Link, error) {
	l, err := r.client.Link.
		UpdateOneID(id).
		SetShortPath(path).
		Save(ctx)
	if err != nil {
		switch err.(type) {
		case *ent.NotFoundError:
			return nil, errors2.WrapErrNotFound(err)
		case *ent.NotSingularError:
			return nil, errors2.WrapErrNotSingular(err)
		default:
			return nil, errors.Wrapf(err, "could not update link %d", id)
		}
	}

	dto, err := link.NewLink(l.ID, l.LongURI, l.CreatedAt)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to update link %s", l.LongURI)
	}

	return dto, nil
}

func (r *linkRepo) Create(ctx context.Context, long string) (*link.Link, error) {
	l, err := r.client.Link.
		Create().
		SetLongURI(long).
		Save(ctx)
	if err != nil {
		switch err.(type) {
		case *ent.NotFoundError:
			return nil, errors2.WrapErrNotFound(err)
		case *ent.NotSingularError:
			return nil, errors2.WrapErrNotSingular(err)
		default:
			return nil, errors.Wrapf(err, "could not create link %s", long)
		}
	}

	dto, err := link.NewLink(l.ID, l.LongURI, l.CreatedAt)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create link %s", l.LongURI)
	}

	return dto, nil
}
