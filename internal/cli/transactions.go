package cli

import (
	"context"
	"encoding/json"
	"fmt"

	wallbittx "github.com/jeremyjsx/wallbit-go/services/transactions"
	"github.com/spf13/cobra"
)

var (
	txPage     int
	txLimit    int
	txStatus   string
	txType     string
	txCurrency string
)

var transactionsCmd = &cobra.Command{
	Use:   "transactions",
	Short: "List account transactions",
}

var transactionsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List transactions with optional filters",
	RunE:  runTransactionsList,
}

func init() {
	transactionsListCmd.Flags().IntVar(&txPage, "page", 0, "Page number (1-based)")
	transactionsListCmd.Flags().IntVar(&txLimit, "limit", 0, "Page size")
	transactionsListCmd.Flags().StringVar(&txStatus, "status", "", "Filter by transaction status")
	transactionsListCmd.Flags().StringVar(&txType, "type", "", "Filter by transaction type")
	transactionsListCmd.Flags().StringVar(&txCurrency, "currency", "", "Filter by currency code")

	transactionsCmd.AddCommand(transactionsListCmd)
	rootCmd.AddCommand(transactionsCmd)
}

func runTransactionsList(cmd *cobra.Command, args []string) error {
	ctx, cancel := context.WithTimeout(cmd.Context(), app.Timeout())
	defer cancel()

	req := &wallbittx.ListRequest{
		Status:   txStatus,
		Type:     txType,
		Currency: txCurrency,
	}
	if txPage > 0 {
		page := txPage
		req.Page = &page
	}
	if txLimit > 0 {
		limit := txLimit
		req.Limit = &limit
	}

	svc, err := app.Services()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	out, err := svc.Transactions.List(ctx, req)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	enc := json.NewEncoder(cmd.OutOrStdout())
	enc.SetIndent("", "  ")
	return enc.Encode(out)
}
