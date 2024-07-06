package cmd

import (
	"context"
	"log"
	"os"

	"github.com/adrianramadhan/synpasis-ecommerce-api/pkg/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func Start() {
	_, cancel := context.WithCancel(context.Background())
	config.LoadEnv(".env")

	quit := make(chan os.Signal)
	go func() {
		<-quit
		cancel()
	}()

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Failed Start Server: %v", err)
	}

}
