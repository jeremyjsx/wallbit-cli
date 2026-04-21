package transactions

import (
	"context"

	wallbittx "github.com/jeremyjsx/wallbit-go/services/transactions"
)

type Service struct {
	sdk *wallbittx.Service
}

func New(sdk *wallbittx.Service) *Service {
	return &Service{sdk: sdk}
}

func (s *Service) List(ctx context.Context, req *wallbittx.ListRequest) (*wallbittx.ListResponse, error) {
	res, err := s.sdk.List(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}
