package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	wallbitaccountdetails "github.com/jeremyjsx/wallbit-go/services/accountdetails"
	"github.com/spf13/cobra"
)

var (
	accountDetailsCountry  string
	accountDetailsCurrency string
)

var accountDetailsCmd = &cobra.Command{
	Use:   "account-details",
	Short: "Query account banking details",
}

var accountDetailsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get account details by optional country/currency filters",
	RunE:  runAccountDetailsGet,
}

func init() {
	accountDetailsGetCmd.Flags().StringVar(&accountDetailsCountry, "country", "", "Country code filter (e.g. US, EU)")
	accountDetailsGetCmd.Flags().StringVar(&accountDetailsCurrency, "currency", "", "Currency code filter (e.g. USD, EUR)")

	accountDetailsCmd.AddCommand(accountDetailsGetCmd)
	rootCmd.AddCommand(accountDetailsCmd)
}

func runAccountDetailsGet(cmd *cobra.Command, args []string) error {
	ctx, cancel := context.WithTimeout(cmd.Context(), app.Timeout())
	defer cancel()

	req := &wallbitaccountdetails.GetRequest{
		Country:  strings.ToUpper(strings.TrimSpace(accountDetailsCountry)),
		Currency: strings.ToUpper(strings.TrimSpace(accountDetailsCurrency)),
	}

	svc, err := app.Services()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	out, err := svc.AccountDetails.Get(ctx, req)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	enc := json.NewEncoder(cmd.OutOrStdout())
	enc.SetIndent("", "  ")
	return enc.Encode(out)
}
