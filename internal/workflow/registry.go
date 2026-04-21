package workflow

import (
	"context"
	"fmt"
	"strings"

	"github.com/jeremyjsx/wallbit-cli/internal/services"
	wallbitaccountdetails "github.com/jeremyjsx/wallbit-go/services/accountdetails"
	wallbitassets "github.com/jeremyjsx/wallbit-go/services/assets"
	wallbitroboadvisor "github.com/jeremyjsx/wallbit-go/services/roboadvisor"
	wallbittrades "github.com/jeremyjsx/wallbit-go/services/trades"
	wallbittransactions "github.com/jeremyjsx/wallbit-go/services/transactions"
	wallbitrates "github.com/jeremyjsx/wallbit-go/services/rates"
	wallbitwallets "github.com/jeremyjsx/wallbit-go/services/wallets"
)

type StepHandler func(ctx context.Context, svc *services.Services, with map[string]any) (any, error)
type StepInputValidator func(with map[string]any) error

var Registry = map[string]StepHandler{
	"rates.get":               runRatesGet,
	"balance.get_checking":    runBalanceGetChecking,
	"balance.get_stocks":      runBalanceGetStocks,
	"wallets.get":             runWalletsGet,
	"assets.list":             runAssetsList,
	"assets.get":              runAssetsGet,
	"account_details.get":     runAccountDetailsGet,
	"transactions.list":       runTransactionsList,
	"cards.list":              runCardsList,
	"cards.block":             runCardsBlock,
	"cards.unblock":           runCardsUnblock,
	"trades.create":           runTradesCreate,
	"roboadvisor.deposit":     runRoboadvisorDeposit,
	"roboadvisor.withdraw":    runRoboadvisorWithdraw,
	"apikey.revoke":           runAPIKeyRevoke,
}

var InputValidators = map[string]StepInputValidator{
	"rates.get":            validateRatesGetInput,
	"assets.get":           validateAssetsGetInput,
	"cards.block":          validateCardsBlockInput,
	"cards.unblock":        validateCardsBlockInput,
	"trades.create":        validateTradesCreateInput,
	"roboadvisor.deposit":  validateRoboadvisorDepositInput,
	"roboadvisor.withdraw": validateRoboadvisorWithdrawInput,
}

func runRatesGet(ctx context.Context, svc *services.Services, with map[string]any) (any, error) {
	source, err := getRequiredString(with, "source")
	if err != nil {
		return nil, err
	}
	dest, err := getRequiredString(with, "dest")
	if err != nil {
		return nil, err
	}
	return svc.Rates.Get(ctx, wallbitrates.GetRequest{
		SourceCurrency: strings.ToUpper(strings.TrimSpace(source)),
		DestCurrency:   strings.ToUpper(strings.TrimSpace(dest)),
	})
}

func runBalanceGetChecking(ctx context.Context, svc *services.Services, with map[string]any) (any, error) {
	return svc.Balance.GetChecking(ctx)
}

func runBalanceGetStocks(ctx context.Context, svc *services.Services, with map[string]any) (any, error) {
	return svc.Balance.GetStocks(ctx)
}

func runWalletsGet(ctx context.Context, svc *services.Services, with map[string]any) (any, error) {
	return svc.Wallets.Get(ctx, &wallbitwallets.GetRequest{
		Currency: strings.ToUpper(getOptionalString(with, "currency")),
		Network:  strings.ToLower(getOptionalString(with, "network")),
	})
}

func runAssetsList(ctx context.Context, svc *services.Services, with map[string]any) (any, error) {
	req := &wallbitassets.ListRequest{
		Category: strings.ToUpper(getOptionalString(with, "category")),
		Search:   getOptionalString(with, "search"),
	}
	if v, ok, err := getOptionalInt(with, "page"); err != nil {
		return nil, err
	} else if ok {
		req.Page = &v
	}
	if v, ok, err := getOptionalInt(with, "limit"); err != nil {
		return nil, err
	} else if ok {
		req.Limit = &v
	}
	return svc.Assets.List(ctx, req)
}

func runAssetsGet(ctx context.Context, svc *services.Services, with map[string]any) (any, error) {
	symbol, err := getRequiredString(with, "symbol")
	if err != nil {
		return nil, err
	}
	return svc.Assets.Get(ctx, strings.ToUpper(strings.TrimSpace(symbol)))
}

func runAccountDetailsGet(ctx context.Context, svc *services.Services, with map[string]any) (any, error) {
	return svc.AccountDetails.Get(ctx, &wallbitaccountdetails.GetRequest{
		Country:  strings.ToUpper(getOptionalString(with, "country")),
		Currency: strings.ToUpper(getOptionalString(with, "currency")),
	})
}

func runTransactionsList(ctx context.Context, svc *services.Services, with map[string]any) (any, error) {
	req := &wallbittransactions.ListRequest{
		Status:   getOptionalString(with, "status"),
		Type:     getOptionalString(with, "type"),
		Currency: getOptionalString(with, "currency"),
	}
	if v, ok, err := getOptionalInt(with, "page"); err != nil {
		return nil, err
	} else if ok {
		req.Page = &v
	}
	if v, ok, err := getOptionalInt(with, "limit"); err != nil {
		return nil, err
	} else if ok {
		req.Limit = &v
	}
	return svc.Transactions.List(ctx, req)
}

func runCardsList(ctx context.Context, svc *services.Services, with map[string]any) (any, error) {
	return svc.Cards.List(ctx)
}

func runCardsBlock(ctx context.Context, svc *services.Services, with map[string]any) (any, error) {
	cardUUID, err := getRequiredString(with, "card_uuid")
	if err != nil {
		return nil, err
	}
	return svc.Cards.Block(ctx, cardUUID)
}

func runCardsUnblock(ctx context.Context, svc *services.Services, with map[string]any) (any, error) {
	cardUUID, err := getRequiredString(with, "card_uuid")
	if err != nil {
		return nil, err
	}
	return svc.Cards.Unblock(ctx, cardUUID)
}

func runTradesCreate(ctx context.Context, svc *services.Services, with map[string]any) (any, error) {
	symbol, err := getRequiredString(with, "symbol")
	if err != nil {
		return nil, err
	}
	direction, err := getRequiredString(with, "direction")
	if err != nil {
		return nil, err
	}
	currency, err := getRequiredString(with, "currency")
	if err != nil {
		return nil, err
	}
	orderType, err := getRequiredString(with, "order_type")
	if err != nil {
		return nil, err
	}

	req := wallbittrades.CreateRequest{
		Symbol:    strings.ToUpper(strings.TrimSpace(symbol)),
		Direction: strings.ToUpper(strings.TrimSpace(direction)),
		Currency:  strings.ToUpper(strings.TrimSpace(currency)),
		OrderType: strings.ToUpper(strings.TrimSpace(orderType)),
	}
	if v, ok, err := getOptionalFloat(with, "amount"); err != nil {
		return nil, err
	} else if ok {
		req.Amount = &v
	}
	if v, ok, err := getOptionalFloat(with, "shares"); err != nil {
		return nil, err
	} else if ok {
		req.Shares = &v
	}
	if v, ok, err := getOptionalFloat(with, "stop_price"); err != nil {
		return nil, err
	} else if ok {
		req.StopPrice = &v
	}
	if v, ok, err := getOptionalFloat(with, "limit_price"); err != nil {
		return nil, err
	} else if ok {
		req.LimitPrice = &v
	}
	if v := getOptionalString(with, "time_in_force"); v != "" {
		timeInForce := strings.ToUpper(v)
		req.TimeInForce = &timeInForce
	}
	return svc.Trades.Create(ctx, req)
}

func runRoboadvisorDeposit(ctx context.Context, svc *services.Services, with map[string]any) (any, error) {
	roboAdvisorID, err := getRequiredInt(with, "robo_advisor_id")
	if err != nil {
		return nil, err
	}
	amount, err := getRequiredFloat(with, "amount")
	if err != nil {
		return nil, err
	}
	from, err := getRequiredString(with, "from")
	if err != nil {
		return nil, err
	}
	req := wallbitroboadvisor.DepositRequest{
		RoboAdvisorID: roboAdvisorID,
		Amount:        amount,
		From:          wallbitroboadvisor.AccountType(strings.ToUpper(strings.TrimSpace(from))),
	}
	return svc.RoboAdvisor.Deposit(ctx, req)
}

func runRoboadvisorWithdraw(ctx context.Context, svc *services.Services, with map[string]any) (any, error) {
	roboAdvisorID, err := getRequiredInt(with, "robo_advisor_id")
	if err != nil {
		return nil, err
	}
	amount, err := getRequiredFloat(with, "amount")
	if err != nil {
		return nil, err
	}
	to, err := getRequiredString(with, "to")
	if err != nil {
		return nil, err
	}
	req := wallbitroboadvisor.WithdrawRequest{
		RoboAdvisorID: roboAdvisorID,
		Amount:        amount,
		To:            wallbitroboadvisor.AccountType(strings.ToUpper(strings.TrimSpace(to))),
	}
	return svc.RoboAdvisor.Withdraw(ctx, req)
}

func runAPIKeyRevoke(ctx context.Context, svc *services.Services, with map[string]any) (any, error) {
	return svc.APIKey.Revoke(ctx)
}

func getRequiredString(with map[string]any, key string) (string, error) {
	if with == nil {
		return "", fmt.Errorf("with.%s is required", key)
	}
	raw, ok := with[key]
	if !ok {
		return "", fmt.Errorf("with.%s is required", key)
	}
	s, ok := raw.(string)
	if !ok {
		return "", fmt.Errorf("with.%s must be a string", key)
	}
	s = strings.TrimSpace(s)
	if s == "" {
		return "", fmt.Errorf("with.%s is required", key)
	}
	return s, nil
}

func getOptionalString(with map[string]any, key string) string {
	if with == nil {
		return ""
	}
	raw, ok := with[key]
	if !ok || raw == nil {
		return ""
	}
	s, ok := raw.(string)
	if !ok {
		return ""
	}
	return strings.TrimSpace(s)
}

func getOptionalInt(with map[string]any, key string) (int, bool, error) {
	if with == nil {
		return 0, false, nil
	}
	raw, ok := with[key]
	if !ok || raw == nil {
		return 0, false, nil
	}
	switch v := raw.(type) {
	case int:
		return v, true, nil
	case int64:
		return int(v), true, nil
	case float64:
		return int(v), true, nil
	default:
		return 0, false, fmt.Errorf("with.%s must be a number", key)
	}
}

func getRequiredInt(with map[string]any, key string) (int, error) {
	v, ok, err := getOptionalInt(with, key)
	if err != nil {
		return 0, err
	}
	if !ok {
		return 0, fmt.Errorf("with.%s is required", key)
	}
	return v, nil
}

func getOptionalFloat(with map[string]any, key string) (float64, bool, error) {
	if with == nil {
		return 0, false, nil
	}
	raw, ok := with[key]
	if !ok || raw == nil {
		return 0, false, nil
	}
	switch v := raw.(type) {
	case float64:
		return v, true, nil
	case float32:
		return float64(v), true, nil
	case int:
		return float64(v), true, nil
	case int64:
		return float64(v), true, nil
	default:
		return 0, false, fmt.Errorf("with.%s must be a number", key)
	}
}

func getRequiredFloat(with map[string]any, key string) (float64, error) {
	v, ok, err := getOptionalFloat(with, key)
	if err != nil {
		return 0, err
	}
	if !ok {
		return 0, fmt.Errorf("with.%s is required", key)
	}
	return v, nil
}

func validateRatesGetInput(with map[string]any) error {
	if _, err := getRequiredString(with, "source"); err != nil {
		return err
	}
	if _, err := getRequiredString(with, "dest"); err != nil {
		return err
	}
	return nil
}

func validateAssetsGetInput(with map[string]any) error {
	_, err := getRequiredString(with, "symbol")
	return err
}

func validateCardsBlockInput(with map[string]any) error {
	_, err := getRequiredString(with, "card_uuid")
	return err
}

func validateTradesCreateInput(with map[string]any) error {
	if _, err := getRequiredString(with, "symbol"); err != nil {
		return err
	}
	if _, err := getRequiredString(with, "direction"); err != nil {
		return err
	}
	if _, err := getRequiredString(with, "currency"); err != nil {
		return err
	}
	if _, err := getRequiredString(with, "order_type"); err != nil {
		return err
	}
	_, hasAmount, err := getOptionalFloat(with, "amount")
	if err != nil {
		return err
	}
	_, hasShares, err := getOptionalFloat(with, "shares")
	if err != nil {
		return err
	}
	if hasAmount == hasShares {
		return fmt.Errorf("with.amount or with.shares must be provided (exactly one)")
	}
	return nil
}

func validateRoboadvisorDepositInput(with map[string]any) error {
	if _, err := getRequiredInt(with, "robo_advisor_id"); err != nil {
		return err
	}
	if amount, err := getRequiredFloat(with, "amount"); err != nil {
		return err
	} else if amount <= 0 {
		return fmt.Errorf("with.amount must be positive")
	}
	from, err := getRequiredString(with, "from")
	if err != nil {
		return err
	}
	v := strings.ToUpper(strings.TrimSpace(from))
	if v != "DEFAULT" && v != "INVESTMENT" {
		return fmt.Errorf("with.from must be DEFAULT or INVESTMENT")
	}
	return nil
}

func validateRoboadvisorWithdrawInput(with map[string]any) error {
	if _, err := getRequiredInt(with, "robo_advisor_id"); err != nil {
		return err
	}
	if amount, err := getRequiredFloat(with, "amount"); err != nil {
		return err
	} else if amount <= 0 {
		return fmt.Errorf("with.amount must be positive")
	}
	to, err := getRequiredString(with, "to")
	if err != nil {
		return err
	}
	v := strings.ToUpper(strings.TrimSpace(to))
	if v != "DEFAULT" && v != "INVESTMENT" {
		return fmt.Errorf("with.to must be DEFAULT or INVESTMENT")
	}
	return nil
}
