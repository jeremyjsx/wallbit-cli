package operations

import (
	"context"

	wallbitoperations "github.com/jeremyjsx/wallbit-go/services/operations"
)

type Service struct {
	sdk *wallbitoperations.Service
}

func New(sdk *wallbitoperations.Service) *Service {
	return &Service{sdk: sdk}
}

func (s *Service) DepositInvestment(ctx context.Context, req wallbitoperations.InvestmentDepositRequest) (*wallbitoperations.Transaction, error) {
	res, err := s.sdk.DepositInvestment(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

func (s *Service) WithdrawInvestment(ctx context.Context, req wallbitoperations.InvestmentWithdrawRequest) (*wallbitoperations.Transaction, error) {
	res, err := s.sdk.WithdrawInvestment(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}
