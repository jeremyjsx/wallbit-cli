package cli

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var balanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "Query account balances",
}

var balanceGetCheckingCmd = &cobra.Command{
	Use:   "checking",
	Short: "Get checking account balances by currency",
	RunE:  runBalanceGetChecking,
}

var balanceGetStocksCmd = &cobra.Command{
	Use:   "stocks",
	Short: "Get stocks portfolio positions",
	RunE:  runBalanceGetStocks,
}

func init() {
	balanceCmd.AddCommand(balanceGetCheckingCmd, balanceGetStocksCmd)
	rootCmd.AddCommand(balanceCmd)
}

func runBalanceGetChecking(cmd *cobra.Command, args []string) error {
	ctx, cancel := context.WithTimeout(cmd.Context(), app.Timeout())
	defer cancel()

	svc, err := app.Services()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	out, err := svc.Balance.GetChecking(ctx)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	enc := json.NewEncoder(cmd.OutOrStdout())
	enc.SetIndent("", "  ")
	return enc.Encode(out)
}

func runBalanceGetStocks(cmd *cobra.Command, args []string) error {
	ctx, cancel := context.WithTimeout(cmd.Context(), app.Timeout())
	defer cancel()

	svc, err := app.Services()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	out, err := svc.Balance.GetStocks(ctx)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	enc := json.NewEncoder(cmd.OutOrStdout())
	enc.SetIndent("", "  ")
	return enc.Encode(out)
}
