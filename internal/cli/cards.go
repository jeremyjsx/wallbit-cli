package cli

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var cardsCmd = &cobra.Command{
	Use:   "cards",
	Short: "Manage payment cards",
}

var cardsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List cards for the account",
	RunE:  runCardsList,
}

var cardsBlockCmd = &cobra.Command{
	Use:   "block <card-uuid>",
	Short: "Suspend a card by UUID",
	Args:  cobra.ExactArgs(1),
	RunE:  runCardsBlock,
}

var cardsUnblockCmd = &cobra.Command{
	Use:   "unblock <card-uuid>",
	Short: "Activate a suspended card by UUID",
	Args:  cobra.ExactArgs(1),
	RunE:  runCardsUnblock,
}

func init() {
	cardsCmd.AddCommand(cardsListCmd, cardsBlockCmd, cardsUnblockCmd)
	rootCmd.AddCommand(cardsCmd)
}

func runCardsList(cmd *cobra.Command, args []string) error {
	ctx, cancel := context.WithTimeout(cmd.Context(), app.Timeout())
	defer cancel()

	svc, err := app.Services()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	out, err := svc.Cards.List(ctx)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	enc := json.NewEncoder(cmd.OutOrStdout())
	enc.SetIndent("", "  ")
	return enc.Encode(out)
}

func runCardsBlock(cmd *cobra.Command, args []string) error {
	ctx, cancel := context.WithTimeout(cmd.Context(), app.Timeout())
	defer cancel()

	uuid := strings.TrimSpace(args[0])
	if uuid == "" {
		return errors.New("card uuid is required")
	}

	svc, err := app.Services()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	out, err := svc.Cards.Block(ctx, uuid)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	enc := json.NewEncoder(cmd.OutOrStdout())
	enc.SetIndent("", "  ")
	return enc.Encode(out)
}

func runCardsUnblock(cmd *cobra.Command, args []string) error {
	ctx, cancel := context.WithTimeout(cmd.Context(), app.Timeout())
	defer cancel()

	uuid := strings.TrimSpace(args[0])
	if uuid == "" {
		return errors.New("card uuid is required")
	}

	svc, err := app.Services()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	out, err := svc.Cards.Unblock(ctx, uuid)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	enc := json.NewEncoder(cmd.OutOrStdout())
	enc.SetIndent("", "  ")
	return enc.Encode(out)
}
