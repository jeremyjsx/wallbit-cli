package cli

import (
	"bufio"
	"errors"
	"fmt"
	"strings"

	"github.com/jeremyjsx/wallbit-cli/internal/credentials"
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Manage API credentials",
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Save your Wallbit API key locally",
	RunE:  runAuthLogin,
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show whether an API key is configured (value is never printed)",
	RunE:  runAuthStatus,
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Remove locally stored API key",
	RunE:  runAuthLogout,
}

func init() {
	authCmd.AddCommand(loginCmd, statusCmd, logoutCmd)
}

func runAuthLogin(cmd *cobra.Command, args []string) error {
	key := strings.TrimSpace(apiKey)
	if key == "" {
		fmt.Fprint(cmd.ErrOrStderr(), "Enter API key: ")
		line, err := bufio.NewReader(cmd.InOrStdin()).ReadString('\n')
		if err != nil {
			return fmt.Errorf("read API key: %w", err)
		}
		key = strings.TrimSpace(line)
	}
	if err := credentials.Save(key); err != nil {
		return err
	}
	fmt.Fprintln(cmd.OutOrStdout(), "API key saved.")
	return nil
}

func runAuthStatus(cmd *cobra.Command, args []string) error {
	fmt.Fprintf(cmd.OutOrStdout(), "%s set: %v\n", credentials.EnvAPIKey, credentials.EnvConfigured())
	fmt.Fprintf(cmd.OutOrStdout(), "credentials file present: %v\n", credentials.FileStoreConfigured())
	_, src, err := credentials.Load("")
	switch {
	case err == nil:
		fmt.Fprintf(cmd.OutOrStdout(), "effective key source: %s\n", src)
	case errors.Is(err, credentials.ErrNotConfigured):
		fmt.Fprintln(cmd.OutOrStdout(), "effective key: not configured")
	default:
		return err
	}
	return nil
}

func runAuthLogout(cmd *cobra.Command, args []string) error {
	if err := credentials.Delete(); err != nil {
		return err
	}
	fmt.Fprintln(cmd.OutOrStdout(), "Local API key removed (if it existed).")
	return nil
}
