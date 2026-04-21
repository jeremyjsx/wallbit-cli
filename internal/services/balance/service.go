package balance

import (
	"context"

	wallbitbalance "github.com/jeremyjsx/wallbit-go/services/balance"
)

type Service struct {
	sdk *wallbitbalance.Service
}

func New(sdk *wallbitbalance.Service) *Service {
	return &Service{sdk: sdk}
}

func (s *Service) GetChecking(ctx context.Context) (*wallbitbalance.CheckingBalanceResponse, error) {
	res, err := s.sdk.GetChecking(ctx)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

func (s *Service) GetStocks(ctx context.Context) (*wallbitbalance.StocksBalanceResponse, error) {
	res, err := s.sdk.GetStocks(ctx)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}
