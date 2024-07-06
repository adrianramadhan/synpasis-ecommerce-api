package cmd

import (
	"log"

	"github.com/adrianramadhan/synpasis-ecommerce-api/pkg/database"
	"github.com/spf13/cobra"
)

var (
	migrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "Run Migration",
		Run: func(cmd *cobra.Command, args []string) {
			db, err := database.Connect()
			if err != nil {
				log.Fatalf("Failed to connect to database: %v", err)
			}

			err = database.AutoMigrate(db)
			if err != nil {
				log.Fatalf("Failed to migrate database: %v", err)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}
