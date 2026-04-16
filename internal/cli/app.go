package cli

import (
	"fmt"
	"time"

	"github.com/jeremyjsx/wallbit-cli/internal/credentials"
	apikeysvc "github.com/jeremyjsx/wallbit-cli/internal/services/apikey"
	balancesvc "github.com/jeremyjsx/wallbit-cli/internal/services/balance"
	transactionssvc "github.com/jeremyjsx/wallbit-cli/internal/services/transactions"
	wallbit "github.com/jeremyjsx/wallbit-go/client"
)

type App struct {
	apiKeyFlag string
	baseURL    string
	timeout    time.Duration

	client          *wallbit.Client
	apiKeySvc       *apikeysvc.Service
	balanceSvc      *balancesvc.Service
	transactionsSvc *transactionssvc.Service
}

func NewApp(apiKeyFlag, baseURL string, timeout time.Duration) *App {
	return &App{
		apiKeyFlag: apiKeyFlag,
		baseURL:    baseURL,
		timeout:    timeout,
	}
}

func (a *App) APIKeyFlag() string {
	return a.apiKeyFlag
}

func (a *App) Timeout() time.Duration {
	return a.timeout
}

func (a *App) Client() (*wallbit.Client, error) {
	if a.client != nil {
		return a.client, nil
	}

	key, _, err := credentials.Load(a.apiKeyFlag)
	if err != nil {
		return nil, err
	}

	c, err := wallbit.NewClient(
		key,
		wallbit.WithBaseURL(a.baseURL),
		wallbit.WithTimeout(a.timeout),
	)
	if err != nil {
		return nil, fmt.Errorf("%w: run wallbit auth login or set %s", err, credentials.EnvAPIKey)
	}

	a.client = c
	return a.client, nil
}

func (a *App) BalanceService() *balancesvc.Service {
	if a.balanceSvc != nil {
		return a.balanceSvc
	}
	a.balanceSvc = balancesvc.New(a.Client)
	return a.balanceSvc
}

func (a *App) TransactionsService() *transactionssvc.Service {
	if a.transactionsSvc != nil {
		return a.transactionsSvc
	}
	a.transactionsSvc = transactionssvc.New(a.Client)
	return a.transactionsSvc
}

func (a *App) APIKeyService() *apikeysvc.Service {
	if a.apiKeySvc != nil {
		return a.apiKeySvc
	}
	a.apiKeySvc = apikeysvc.New(a.Client)
	return a.apiKeySvc
}
