package workflow

import (
	"context"
	"fmt"
	"strings"

	"github.com/jeremyjsx/wallbit-cli/internal/services"
	wallbitrates "github.com/jeremyjsx/wallbit-go/services/rates"
)

type StepHandler func(ctx context.Context, svc *services.Services, with map[string]any) (any, error)

var Registry = map[string]StepHandler{
	"rates.get":            runRatesGet,
	"balance.get_checking": runBalanceGetChecking,
	"balance.get_stocks":   runBalanceGetStocks,
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
