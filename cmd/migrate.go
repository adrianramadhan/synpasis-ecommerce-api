package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run Migration",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run Migration")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
