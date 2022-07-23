package sqlite

import (
	"database/sql"
	"log"

	ent "github.com/alexgtn/go-linkshort/tools/ent/codegen"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/mattn/go-sqlite3"
)

// OpenEnt new connection
func OpenEnt(databaseUrl string) *ent.Client {
	db, err := sql.Open(dialect.SQLite, databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.SQLite, db)

	return ent.NewClient(ent.Driver(drv))
}

// Open new connection
func Open(databaseUrl string) *entsql.Driver {
	db, err := entsql.Open(dialect.SQLite, databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
