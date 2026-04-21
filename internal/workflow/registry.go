package workflow

import (
	"context"
	"fmt"
	"strings"

	"github.com/jeremyjsx/wallbit-cli/internal/services"
	wallbitaccountdetails "github.com/jeremyjsx/wallbit-go/services/accountdetails"
	wallbitassets "github.com/jeremyjsx/wallbit-go/services/assets"
	wallbittransactions "github.com/jeremyjsx/wallbit-go/services/transactions"
	wallbitrates "github.com/jeremyjsx/wallbit-go/services/rates"
	wallbitwallets "github.com/jeremyjsx/wallbit-go/services/wallets"
)

type StepHandler func(ctx context.Context, svc *services.Services, with map[string]any) (any, error)

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
