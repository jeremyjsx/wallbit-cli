package roboadvisor

import (
	"context"

	wallbitroboadvisor "github.com/jeremyjsx/wallbit-go/services/roboadvisor"
)

type Service struct {
	sdk *wallbitroboadvisor.Service
}

func New(sdk *wallbitroboadvisor.Service) *Service {
	return &Service{sdk: sdk}
}

func (s *Service) GetBalance(ctx context.Context) (*wallbitroboadvisor.GetBalanceResponse, error) {
	res, err := s.sdk.GetBalance(ctx)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

func (s *Service) Deposit(ctx context.Context, req wallbitroboadvisor.DepositRequest) (*wallbitroboadvisor.DepositResponse, error) {
	res, err := s.sdk.Deposit(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

func (s *Service) Withdraw(ctx context.Context, req wallbitroboadvisor.WithdrawRequest) (*wallbitroboadvisor.WithdrawResponse, error) {
	res, err := s.sdk.Withdraw(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}
