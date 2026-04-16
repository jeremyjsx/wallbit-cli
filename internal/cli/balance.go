package cli

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jeremyjsx/wallbit-cli/internal/balance"
	"github.com/jeremyjsx/wallbit-cli/internal/client"
	"github.com/jeremyjsx/wallbit-cli/internal/credentials"
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

func init() {
	balanceCmd.AddCommand(balanceGetCheckingCmd)
	rootCmd.AddCommand(balanceCmd)
}

func runBalanceGetChecking(cmd *cobra.Command, args []string) error {
	key, _, err := credentials.Load(apiKey)
	if err != nil {
		return err
	}
	c, err := client.New(baseURL, key, requestTimeout)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(cmd.Context(), requestTimeout)
	defer cancel()

	out, err := balance.GetChecking(ctx, c)
	if err != nil {
		var apiErr client.APIStatusError
		if errors.As(err, &apiErr) && apiErr.Status == 401 {
			return fmt.Errorf("%w: run wallbit auth login or set %s", err, credentials.EnvAPIKey)
		}
		return err
	}
	enc := json.NewEncoder(cmd.OutOrStdout())
	enc.SetIndent("", "  ")
	return enc.Encode(out)
}
