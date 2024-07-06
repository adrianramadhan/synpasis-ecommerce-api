package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func Start() {
	rootCmd := &cobra.Command{}
	_, cancel := context.WithCancel(context.Background())

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
			fmt.Println("Starting REST Server");
		},
	},
	}

	rootCmd.AddCommand(cmd...)
	if err := rootCmd.Execute(); err != nil{
		log.Fatalf("Failed Start Server: %v", err)
	}
}