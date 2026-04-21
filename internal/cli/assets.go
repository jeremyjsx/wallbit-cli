package cli

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	wallbitassets "github.com/jeremyjsx/wallbit-go/services/assets"
	"github.com/spf13/cobra"
)

var (
	assetsCategory string
	assetsSearch   string
	assetsPage     int
	assetsLimit    int
)

var assetsCmd = &cobra.Command{
	Use:   "assets",
	Short: "Query supported assets",
}

var assetsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List assets with optional filters",
	RunE:  runAssetsList,
}

var assetsGetCmd = &cobra.Command{
	Use:   "get <symbol>",
	Short: "Get a single asset by symbol",
	Args:  cobra.ExactArgs(1),
	RunE:  runAssetsGet,
}

func init() {
	assetsListCmd.Flags().StringVar(&assetsCategory, "category", "", "Asset category filter")
	assetsListCmd.Flags().StringVar(&assetsSearch, "search", "", "Search text for symbol/name")
	assetsListCmd.Flags().IntVar(&assetsPage, "page", 0, "Page number (1-based)")
	assetsListCmd.Flags().IntVar(&assetsLimit, "limit", 0, "Page size")

	assetsCmd.AddCommand(assetsListCmd, assetsGetCmd)
	rootCmd.AddCommand(assetsCmd)
}

func runAssetsList(cmd *cobra.Command, args []string) error {
	ctx, cancel := context.WithTimeout(cmd.Context(), app.Timeout())
	defer cancel()

	req := &wallbitassets.ListRequest{
		Category: strings.ToUpper(strings.TrimSpace(assetsCategory)),
		Search:   strings.TrimSpace(assetsSearch),
	}
	if assetsPage > 0 {
		page := assetsPage
		req.Page = &page
	}
	if assetsLimit > 0 {
		limit := assetsLimit
		req.Limit = &limit
	}

	svc, err := app.Services()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	out, err := svc.Assets.List(ctx, req)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	enc := json.NewEncoder(cmd.OutOrStdout())
	enc.SetIndent("", "  ")
	return enc.Encode(out)
}

func runAssetsGet(cmd *cobra.Command, args []string) error {
	ctx, cancel := context.WithTimeout(cmd.Context(), app.Timeout())
	defer cancel()

	symbol := strings.ToUpper(strings.TrimSpace(args[0]))
	if symbol == "" {
		return errors.New("symbol is required")
	}

	svc, err := app.Services()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	out, err := svc.Assets.Get(ctx, symbol)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	enc := json.NewEncoder(cmd.OutOrStdout())
	enc.SetIndent("", "  ")
	return enc.Encode(out)
}
