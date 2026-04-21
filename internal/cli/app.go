package cli

import (
	"fmt"
	"time"

	"github.com/jeremyjsx/wallbit-cli/internal/credentials"
	"github.com/jeremyjsx/wallbit-cli/internal/services"
	"github.com/jeremyjsx/wallbit-go/wallbit"
)

type App struct {
	apiKeyFlag string
	baseURL    string
	timeout    time.Duration

	client   *wallbit.Client
	services *services.Services
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

func (a *App) Services() (*services.Services, error) {
	if a.services != nil {
		return a.services, nil
	}
	c, err := a.Client()
	if err != nil {
		return nil, err
	}
	a.services = services.New(c)
	return a.services, nil
}
