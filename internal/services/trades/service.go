package trades

import (
	"context"

	wallbittrades "github.com/jeremyjsx/wallbit-go/services/trades"
)

type Service struct {
	sdk *wallbittrades.Service
}

func New(sdk *wallbittrades.Service) *Service {
	return &Service{sdk: sdk}
}

func (s *Service) Create(ctx context.Context, req wallbittrades.CreateRequest) (*wallbittrades.CreateResponse, error) {
	res, err := s.sdk.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}
