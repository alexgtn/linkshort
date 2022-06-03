package link

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alexgtn/go-linkshort/infra/sqlite"
)

var long = "https://jsonplaceholder.typicode.com/albums"

func TestLinkRepo_Create(t *testing.T) {
	c := sqlite.OpenEnt("file:mockrepo?mode=memory&cache=shared&_fk=1")
	r := NewLinkRepo(c)

	_, err := r.Create(context.Background(), long)
	assert.NoError(t, err)

	// error on duplicates
	_, err = r.Create(context.Background(), long)
	assert.Error(t, err)

	existing, err := r.GetOne(context.Background(), long)
	assert.NoError(t, err)

	assert.Equal(t, existing.LongURL(), long)
}

func TestLinkRepo_GetOne(t *testing.T) {
	c := sqlite.OpenEnt("file:mockrepo?mode=memory&cache=shared&_fk=1")
	r := NewLinkRepo(c)

	l, err := r.Create(context.Background(), long)
	assert.NoError(t, err)

	existing, err := r.GetOne(context.Background(), l.LongURL())
	assert.NoError(t, err)
	assert.Equal(t, existing.LongURL(), long)

	_, err = r.GetOne(context.Background(), "nonexistent")
	assert.Error(t, err)
}
