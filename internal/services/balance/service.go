package balance

import (
	"context"

	wallbitbalance "github.com/jeremyjsx/wallbit-go/services/balance"
	"github.com/jeremyjsx/wallbit-go/wallbit"
)

type ClientProvider func() (*wallbit.Client, error)

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
	res, err := c.Balance.GetChecking(ctx)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

func (s *Service) GetStocks(ctx context.Context) (*wallbitbalance.StocksBalanceResponse, error) {
	c, err := s.clientProvider()
	if err != nil {
		return nil, err
	}
	res, err := c.Balance.GetStocks(ctx)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}
