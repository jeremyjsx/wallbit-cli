package cli

import (
	"context"
	"fmt"
	"strings"

	wallbitwallets "github.com/jeremyjsx/wallbit-go/services/wallets"
	"github.com/spf13/cobra"
)

var (
	walletsCurrency string
	walletsNetwork  string
)

var walletsCmd = &cobra.Command{
	Use:   "wallets",
	Short: "Query wallet addresses",
}

var walletsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get wallet addresses with optional filters",
	RunE:  runWalletsGet,
}

func init() {
	walletsGetCmd.Flags().StringVar(&walletsCurrency, "currency", "", "Filter by currency code (e.g. USDC)")
	walletsGetCmd.Flags().StringVar(&walletsNetwork, "network", "", "Filter by blockchain network (e.g. POLYGON)")

	walletsCmd.AddCommand(walletsGetCmd)
	rootCmd.AddCommand(walletsCmd)
}

func runWalletsGet(cmd *cobra.Command, args []string) error {
	ctx, cancel := context.WithTimeout(cmd.Context(), app.Timeout())
	defer cancel()

	req := &wallbitwallets.GetRequest{
		Currency: strings.ToUpper(strings.TrimSpace(walletsCurrency)),
		Network:  strings.ToUpper(strings.TrimSpace(walletsNetwork)),
	}

	svc, err := app.Services()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	var out any
	err = runWithLoading(cmd.ErrOrStderr(), func() error {
		res, err := svc.Wallets.Get(ctx, req)
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
