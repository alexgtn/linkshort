package link

import (
	"encoding/base64"
	"errors"
	"time"

	errors2 "github.com/pkg/errors"
)

const MaxLen = 2048

// Link is a link
type Link struct {
	id            int
	short         string
	long          string
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

	short := base64.RawStdEncoding.EncodeToString([]byte(long))

	return &Link{id, short, long, 0, createdAt}, nil
}

func (u *Link) ID() int {
	return u.id
}

func (u *Link) Short() string {
	return u.short
}

func (u *Link) Long() string {
	return u.long
}

func (u *Link) CreatedAt() time.Time {
	return u.createdAt
}

func (u *Link) AccessedTimes() int {
	return u.id
}
