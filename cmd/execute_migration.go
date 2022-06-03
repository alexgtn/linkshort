package cmd

import (
	"context"
	"fmt"
	"log"

	"entgo.io/ent/dialect/sql/schema"

	"github.com/alexgtn/go-linkshort/ent"
	"github.com/alexgtn/go-linkshort/ent/migrate"
	"github.com/alexgtn/go-linkshort/infra/sqlite"

	"github.com/spf13/cobra"
)

// executeMigrationCmd represents the executeMigration command
var executeMigrationCmd = &cobra.Command{
	Use: "execute-migration",
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
