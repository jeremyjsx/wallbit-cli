package cli

import (
	"context"
	"errors"
	"fmt"
	"strings"

	wallbitfees "github.com/jeremyjsx/wallbit-go/services/fees"
	"github.com/spf13/cobra"
)

var feeType string

var feesCmd = &cobra.Command{
	Use:   "fees",
	Short: "Query account fee settings",
}

var feesGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get fee settings by type",
	RunE:  runFeesGet,
}

func init() {
	feesGetCmd.Flags().StringVar(&feeType, "type", "TRADE", "Fee type (e.g. TRADE)")
	feesCmd.AddCommand(feesGetCmd)
	rootCmd.AddCommand(feesCmd)
}

func runFeesGet(cmd *cobra.Command, args []string) error {
	ctx, cancel := context.WithTimeout(cmd.Context(), app.Timeout())
	defer cancel()

	req := wallbitfees.GetRequest{
		Type: strings.ToUpper(strings.TrimSpace(feeType)),
	}
	if req.Type == "" {
		return errors.New("--type is required")
	}

	svc, err := app.Services()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	var out any
	err = runWithLoading(cmd.ErrOrStderr(), func() error {
		res, err := svc.Fees.Get(ctx, req)
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
