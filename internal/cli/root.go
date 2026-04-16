package cli

import (
	"os"
	"time"

	"github.com/spf13/cobra"
)

var (
	apiKey         string
	baseURL        string
	requestTimeout time.Duration
)

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
	rootCmd.PersistentFlags().StringVar(&baseURL, "base-url", "https://api.wallbit.io", "Wallbit API base URL")
	rootCmd.PersistentFlags().DurationVar(&requestTimeout, "timeout", 30*time.Second, "HTTP client timeout")

	rootCmd.AddCommand(authCmd)

	defaultHelp := rootCmd.HelpFunc()
	rootCmd.SetHelpFunc(func(c *cobra.Command, args []string) {
		if c == rootCmd {
			fprintLogo(c.OutOrStdout())
		}
		defaultHelp(c, args)
	})
}
