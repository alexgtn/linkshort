package link

import (
	"errors"
	"time"

	errors2 "github.com/pkg/errors"

	"github.com/alexgtn/go-linkshort/common/encoding"
)

const MaxLen = 2048

// Link is a link
type Link struct {
	id int
	// eg. platform-base-url/shortPath
	shortPath string
	// original long uri that was shortened
	longURL       string
	accessedTimes int
	createdAt     time.Time
}

var errLink = errors.New("could not create link")

func NewLink(id int, long string, createdAt time.Time) (*Link, error) {
	if createdAt.IsZero() {
		return nil, errors2.Wrap(errLink, "createdAt is zero")
	}

	if id <= 0 {
		return nil, errors2.Wrap(errLink, "id is not positive")
	}

	short := encoding.ToBase62(long)

	return &Link{id, short, long, 0, createdAt}, nil
}

func (u *Link) ID() int {
	return u.id
}

func (u *Link) ShortPath() string {
	return u.shortPath
}

func (u *Link) LongURL() string {
	return u.longURL
}

func (u *Link) CreatedAt() time.Time {
	return u.createdAt
}

func (u *Link) AccessedTimes() int {
	return u.id
}
