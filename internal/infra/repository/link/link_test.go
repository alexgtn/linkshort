package link

import (
	"context"
	"log"
	"os"
	"testing"

	"entgo.io/ent/dialect/sql/schema"
	"github.com/stretchr/testify/assert"

	"github.com/alexgtn/go-linkshort/tools/ent/codegen/migrate"

	"github.com/alexgtn/go-linkshort/internal/infra/sqlite"
)

const long = "https://jsonplaceholder.typicode.com/albums"

const databaseUrl = "file:mockrepo?mode=memory&cache=shared&_fk=1"

var db *ent.Client

func TestMain(m *testing.M) {
	db = sqlite.OpenEnt(databaseUrl)
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {
			log.Fatal("error closing client")
		}
	}(db)

	// Run migration.
	err := db.Schema.Create(context.Background(),
		schema.WithAtlas(true),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true))
	if err != nil {
		log.Fatalf(err.Error())
	}

	code := m.Run()
	os.Exit(code)
}

func cleanDB() {
	_, err := db.Link.Delete().Exec(context.Background())
	if err != nil {
		log.Fatalf("failed to cleanup links table: %v", err)
	}
}

func TestLinkRepo_Create(t *testing.T) {
	r := NewLinkRepo(db)

	_, err := r.Create(context.Background(), long)
	assert.NoError(t, err)

	// error on duplicates
	_, err = r.Create(context.Background(), long)
	assert.Error(t, err)

	existing, err := r.GetOneByLongURL(context.Background(), long)
	assert.NoError(t, err)

	assert.Equal(t, existing.LongURL(), long)

	t.Cleanup(cleanDB)
}

func TestLinkRepo_GetOne(t *testing.T) {
	t.Run("get one by long URL", func(t *testing.T) {
		r := NewLinkRepo(db)

		l, err := r.Create(context.Background(), long)
		assert.NoError(t, err)

		existing, err := r.GetOneByLongURL(context.Background(), l.LongURL())
		assert.NoError(t, err)
		assert.Equal(t, existing.LongURL(), long)

		_, err = r.GetOneByLongURL(context.Background(), "nonexistent")
		assert.Error(t, err)

		t.Cleanup(cleanDB)
	})

	t.Run("get one by short path", func(t *testing.T) {
		r := NewLinkRepo(db)

		l, err := r.Create(context.Background(), long)
		assert.NoError(t, err)

		_, err = r.SetShortPath(context.Background(), l.ID(), l.ShortPath())
		assert.NoError(t, err)

		existing, err := r.GetOneByShortPath(context.Background(), l.ShortPath())
		assert.NoError(t, err)
		assert.Equal(t, existing.LongURL(), long)

		_, err = r.GetOneByShortPath(context.Background(), "nonexistent")
		assert.Error(t, err)

		t.Cleanup(cleanDB)
	})
}
