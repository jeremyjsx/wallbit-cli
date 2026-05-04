package cli

import (
	"context"
	"errors"
	"fmt"
	"strings"

	wallbitrates "github.com/jeremyjsx/wallbit-go/services/rates"
	"github.com/spf13/cobra"
)

var (
	rateSource string
	rateDest   string
)

var ratesCmd = &cobra.Command{
	Use:   "rates",
	Short: "Query exchange rates",
}

var ratesGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the exchange rate between two currencies",
	RunE:  runRatesGet,
}

func init() {
	ratesGetCmd.Flags().StringVar(&rateSource, "source", "", "Source currency code (e.g. USD)")
	ratesGetCmd.Flags().StringVar(&rateDest, "dest", "", "Destination currency code (e.g. EUR)")
	_ = ratesGetCmd.MarkFlagRequired("source")
	_ = ratesGetCmd.MarkFlagRequired("dest")

	ratesCmd.AddCommand(ratesGetCmd)
	rootCmd.AddCommand(ratesCmd)
}

func runRatesGet(cmd *cobra.Command, args []string) error {
	ctx, cancel := context.WithTimeout(cmd.Context(), app.Timeout())
	defer cancel()

	req := wallbitrates.GetRequest{
		SourceCurrency: strings.ToUpper(strings.TrimSpace(rateSource)),
		DestCurrency:   strings.ToUpper(strings.TrimSpace(rateDest)),
	}
	if req.SourceCurrency == "" || req.DestCurrency == "" {
		return errors.New("--source and --dest are required")
	}

	svc, err := app.Services()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	var out any
	err = runWithLoading(cmd.ErrOrStderr(), func() error {
		res, err := svc.Rates.Get(ctx, req)
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
