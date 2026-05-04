package cli

import (
	"context"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var revokeConfirm bool

var apiKeyCmd = &cobra.Command{
	Use:   "apikey",
	Short: "Manage API key actions",
}

var apiKeyRevokeCmd = &cobra.Command{
	Use:   "revoke",
	Short: "Revoke the current API key",
	RunE:  runAPIKeyRevoke,
}

func init() {
	apiKeyRevokeCmd.Flags().BoolVar(&revokeConfirm, "confirm", false, "Required safety confirmation for API key revocation")
	apiKeyCmd.AddCommand(apiKeyRevokeCmd)
	rootCmd.AddCommand(apiKeyCmd)
}

func runAPIKeyRevoke(cmd *cobra.Command, args []string) error {
	if !revokeConfirm {
		return errors.New("refusing to revoke API key without --confirm")
	}

	ctx, cancel := context.WithTimeout(cmd.Context(), app.Timeout())
	defer cancel()

	svc, err := app.Services()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	var out any
	err = runWithLoading(cmd.ErrOrStderr(), func() error {
		res, err := svc.APIKey.Revoke(ctx)
		if err != nil {
			return err
		}
		out = res
		return nil
	})
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	return writeJSON(out, cmd)
}
