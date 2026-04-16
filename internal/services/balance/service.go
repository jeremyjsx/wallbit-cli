package balance

import (
	"context"

	wallbitclient "github.com/jeremyjsx/wallbit-go/client"
	wallbitbalance "github.com/jeremyjsx/wallbit-go/services/balance"
)

type ClientProvider func() (*wallbitclient.Client, error)

type Service struct {
	clientProvider ClientProvider
}

func New(clientProvider ClientProvider) *Service {
	return &Service{clientProvider: clientProvider}
}

func (s *Service) GetChecking(ctx context.Context) (*wallbitbalance.CheckingBalanceResponse, error) {
	c, err := s.clientProvider()
	if err != nil {
		return nil, err
	}
	return c.Balance.GetChecking(ctx)
}

func (s *Service) GetStocks(ctx context.Context) (*wallbitbalance.StocksBalanceResponse, error) {
	c, err := s.clientProvider()
	if err != nil {
		return nil, err
	}
	return c.Balance.GetStocks(ctx)
}
