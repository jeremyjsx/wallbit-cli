package cli

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	wallbitoperations "github.com/jeremyjsx/wallbit-go/services/operations"
	"github.com/spf13/cobra"
)

var (
	opsCurrency string
	opsAmount   float64
)

var operationsCmd = &cobra.Command{
	Use:   "operations",
	Short: "Internal account operations",
}

var operationsDepositInvestmentCmd = &cobra.Command{
	Use:   "deposit-investment",
	Short: "Move funds from default account to investment account",
	RunE:  runOperationsDepositInvestment,
}

var operationsWithdrawInvestmentCmd = &cobra.Command{
	Use:   "withdraw-investment",
	Short: "Move funds from investment account to default account",
	RunE:  runOperationsWithdrawInvestment,
}

func init() {
	operationsDepositInvestmentCmd.Flags().StringVar(&opsCurrency, "currency", "", "Currency code (e.g. USD)")
	operationsDepositInvestmentCmd.Flags().Float64Var(&opsAmount, "amount", 0, "Amount to move")
	_ = operationsDepositInvestmentCmd.MarkFlagRequired("currency")
	_ = operationsDepositInvestmentCmd.MarkFlagRequired("amount")

	operationsWithdrawInvestmentCmd.Flags().StringVar(&opsCurrency, "currency", "", "Currency code (e.g. USD)")
	operationsWithdrawInvestmentCmd.Flags().Float64Var(&opsAmount, "amount", 0, "Amount to move")
	_ = operationsWithdrawInvestmentCmd.MarkFlagRequired("currency")
	_ = operationsWithdrawInvestmentCmd.MarkFlagRequired("amount")

	operationsCmd.AddCommand(operationsDepositInvestmentCmd, operationsWithdrawInvestmentCmd)
	rootCmd.AddCommand(operationsCmd)
}

func runOperationsDepositInvestment(cmd *cobra.Command, args []string) error {
	ctx, cancel := context.WithTimeout(cmd.Context(), app.Timeout())
	defer cancel()

	currency := strings.ToUpper(strings.TrimSpace(opsCurrency))
	if currency == "" {
		return errors.New("--currency is required")
	}
	if opsAmount <= 0 {
		return errors.New("--amount must be positive")
	}

	req := wallbitoperations.InvestmentDepositRequest{
		Currency: currency,
		Amount:   opsAmount,
	}

	svc, err := app.Services()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	out, err := svc.Operations.DepositInvestment(ctx, req)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	enc := json.NewEncoder(cmd.OutOrStdout())
	enc.SetIndent("", "  ")
	return enc.Encode(out)
}

func runOperationsWithdrawInvestment(cmd *cobra.Command, args []string) error {
	ctx, cancel := context.WithTimeout(cmd.Context(), app.Timeout())
	defer cancel()

	currency := strings.ToUpper(strings.TrimSpace(opsCurrency))
	if currency == "" {
		return errors.New("--currency is required")
	}
	if opsAmount <= 0 {
		return errors.New("--amount must be positive")
	}

	req := wallbitoperations.InvestmentWithdrawRequest{
		Currency: currency,
		Amount:   opsAmount,
	}

	svc, err := app.Services()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	out, err := svc.Operations.WithdrawInvestment(ctx, req)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	enc := json.NewEncoder(cmd.OutOrStdout())
	enc.SetIndent("", "  ")
	return enc.Encode(out)
}
