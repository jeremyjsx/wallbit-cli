package cli

import (
	"context"
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
	var out any
	err = runWithLoading(cmd.ErrOrStderr(), func() error {
		res, err := svc.Balance.GetChecking(ctx)
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

func runBalanceGetStocks(cmd *cobra.Command, args []string) error {
	ctx, cancel := context.WithTimeout(cmd.Context(), app.Timeout())
	defer cancel()

	svc, err := app.Services()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	var out any
	err = runWithLoading(cmd.ErrOrStderr(), func() error {
		res, err := svc.Balance.GetStocks(ctx)
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
