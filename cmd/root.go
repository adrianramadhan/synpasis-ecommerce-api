package cmd

import (
	"context"
	"log"
	"os"

	"github.com/adrianramadhan/synpasis-ecommerce-api/cmd/rest"
	"github.com/adrianramadhan/synpasis-ecommerce-api/pkg/config"
	"github.com/spf13/cobra"
)

func Start() {
	_, cancel := context.WithCancel(context.Background())
	config.LoadEnv(".env")
	rootCmd := &cobra.Command{}

	quit := make(chan os.Signal)
	go func() {
		<-quit
		cancel()
	}()

	cmd := []*cobra.Command{
		{
		Use:   "rest",
		Short: "Run REST Server",
		Run: func(cmd *cobra.Command, args []string) {
			rest.StartRest()
		},
	},
	}

	rootCmd.AddCommand(cmd...)
	if err := rootCmd.Execute(); err != nil{
		log.Fatalf("Failed Start Server: %v", err)
	}
}