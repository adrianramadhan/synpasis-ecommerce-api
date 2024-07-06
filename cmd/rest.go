package cmd

import (
	"github.com/adrianramadhan/synpasis-ecommerce-api/cmd/rest"
	"github.com/spf13/cobra"
)

var restCmd = &cobra.Command{
	Use:   "rest",
	Short: "Run REST Server",
	Run: func(cmd *cobra.Command, args []string) {
		rest.StartRest()
	},
}

func init() {
	rootCmd.AddCommand(restCmd)
}
