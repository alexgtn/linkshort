package cmd

import (
	"context"
	"fmt"
	"log"

	"entgo.io/ent/dialect/sql/schema"

	"github.com/alexgtn/go-linkshort/internal/infra/sqlite"
	ent "github.com/alexgtn/go-linkshort/tools/ent/codegen"
	"github.com/alexgtn/go-linkshort/tools/ent/codegen/migrate"

	"github.com/spf13/cobra"
)

// executeMigrationCmd represents the executeMigration command
var executeMigrationCmd = &cobra.Command{
	Use:   "execute-migration",
	Short: "Runs migration against a local db",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing migration")
		client := sqlite.OpenEnt(cfg.DatabaseURL)

		defer func(client *ent.Client) {
			err := client.Close()
			if err != nil {
				log.Fatal("error closing client")
			}
		}(client)
		ctx := context.Background()
		// Run migration.
		err := client.Schema.Create(ctx,
			schema.WithAtlas(true),
			migrate.WithDropIndex(true),
			migrate.WithDropColumn(true))
		if err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(executeMigrationCmd)
}
