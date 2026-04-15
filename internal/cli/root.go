package cli

import (
	"os"

	"github.com/spf13/cobra"
)

var apiKey string

var rootCmd = &cobra.Command{
	Use:   "wallbit",
	Short: "Wallbit CLI",
	Long:  "Command-line interface for the Wallbit API.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&apiKey, "api-key", "", "Wallbit API key (optional; overrides env and stored credentials)")

	rootCmd.AddCommand(authCmd)
}
