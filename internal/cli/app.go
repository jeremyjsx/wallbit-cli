package cli

import (
	"fmt"
	"time"

	"github.com/jeremyjsx/wallbit-cli/internal/credentials"
	wallbit "github.com/jeremyjsx/wallbit-go/client"
)

type App struct {
	apiKeyFlag string
	baseURL    string
	timeout    time.Duration

	client *wallbit.Client
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
