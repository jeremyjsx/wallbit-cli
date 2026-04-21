package cli

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	tradessvc "github.com/jeremyjsx/wallbit-cli/internal/services/trades"
	"github.com/spf13/cobra"
)

var (
	tradeSymbol      string
	tradeDirection   string
	tradeCurrency    string
	tradeOrderType   string
	tradeAmount      float64
	tradeShares      float64
	tradeStopPrice   float64
	tradeLimitPrice  float64
	tradeTimeInForce string
)

var tradesCmd = &cobra.Command{
	Use:   "trades",
	Short: "Execute trading operations",
}

var tradesCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a trade order",
	RunE:  runTradesCreate,
}

func init() {
	tradesCreateCmd.Flags().StringVar(&tradeSymbol, "symbol", "", "Asset symbol (e.g. AAPL)")
	tradesCreateCmd.Flags().StringVar(&tradeDirection, "direction", "", "Trade direction: BUY or SELL")
	tradesCreateCmd.Flags().StringVar(&tradeCurrency, "currency", "USD", "Trade currency (currently USD)")
	tradesCreateCmd.Flags().StringVar(&tradeOrderType, "order-type", "", "Order type: MARKET, LIMIT, STOP, STOP_LIMIT")
	tradesCreateCmd.Flags().Float64Var(&tradeAmount, "amount", 0, "Amount in USD (mutually exclusive with --shares)")
	tradesCreateCmd.Flags().Float64Var(&tradeShares, "shares", 0, "Shares quantity (mutually exclusive with --amount)")
	tradesCreateCmd.Flags().Float64Var(&tradeStopPrice, "stop-price", 0, "Stop price (required for STOP and STOP_LIMIT)")
	tradesCreateCmd.Flags().Float64Var(&tradeLimitPrice, "limit-price", 0, "Limit price (required for LIMIT and STOP_LIMIT)")
	tradesCreateCmd.Flags().StringVar(&tradeTimeInForce, "time-in-force", "", "Time in force for LIMIT orders: DAY or GTC")

	_ = tradesCreateCmd.MarkFlagRequired("symbol")
	_ = tradesCreateCmd.MarkFlagRequired("direction")
	_ = tradesCreateCmd.MarkFlagRequired("order-type")

	tradesCmd.AddCommand(tradesCreateCmd)
	rootCmd.AddCommand(tradesCmd)
}

func runTradesCreate(cmd *cobra.Command, args []string) error {
	ctx, cancel := context.WithTimeout(cmd.Context(), app.Timeout())
	defer cancel()

	orderType := strings.ToUpper(strings.TrimSpace(tradeOrderType))
	direction := strings.ToUpper(strings.TrimSpace(tradeDirection))
	currency := strings.ToUpper(strings.TrimSpace(tradeCurrency))
	timeInForce := strings.ToUpper(strings.TrimSpace(tradeTimeInForce))

	amountProvided := cmd.Flags().Changed("amount")
	sharesProvided := cmd.Flags().Changed("shares")
	stopPriceProvided := cmd.Flags().Changed("stop-price")
	limitPriceProvided := cmd.Flags().Changed("limit-price")

	if amountProvided == sharesProvided {
		return errors.New("exactly one of --amount or --shares is required")
	}

	switch orderType {
	case "MARKET":
		// No additional fields required.
	case "LIMIT":
		if !limitPriceProvided {
			return errors.New("--limit-price is required when --order-type is LIMIT")
		}
		if timeInForce == "" {
			return errors.New("--time-in-force is required when --order-type is LIMIT")
		}
	case "STOP":
		if !stopPriceProvided {
			return errors.New("--stop-price is required when --order-type is STOP")
		}
	case "STOP_LIMIT":
		if !stopPriceProvided {
			return errors.New("--stop-price is required when --order-type is STOP_LIMIT")
		}
		if !limitPriceProvided {
			return errors.New("--limit-price is required when --order-type is STOP_LIMIT")
		}
	default:
		return fmt.Errorf("invalid --order-type %q, expected MARKET, LIMIT, STOP, STOP_LIMIT", orderType)
	}

	if direction != "BUY" && direction != "SELL" {
		return fmt.Errorf("invalid --direction %q, expected BUY or SELL", direction)
	}

	if currency == "" {
		return errors.New("--currency is required")
	}

	req := &tradessvc.CreateInput{
		Symbol:      strings.TrimSpace(tradeSymbol),
		Direction:   direction,
		Currency:    currency,
		OrderType:   orderType,
		TimeInForce: timeInForce,
	}

	if req.Symbol == "" {
		return errors.New("--symbol is required")
	}

	if amountProvided {
		amount := tradeAmount
		req.Amount = &amount
	}
	if sharesProvided {
		shares := tradeShares
		req.Shares = &shares
	}
	if stopPriceProvided {
		stopPrice := tradeStopPrice
		req.StopPrice = &stopPrice
	}
	if limitPriceProvided {
		limitPrice := tradeLimitPrice
		req.LimitPrice = &limitPrice
	}

	svc, err := app.Services()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	out, err := svc.Trades.Create(ctx, req)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	enc := json.NewEncoder(cmd.OutOrStdout())
	enc.SetIndent("", "  ")
	return enc.Encode(out)
}
