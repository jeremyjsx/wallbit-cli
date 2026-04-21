package wallets

import (
	"context"

	wallbitwallets "github.com/jeremyjsx/wallbit-go/services/wallets"
)

type Service struct {
	sdk *wallbitwallets.Service
}

func New(sdk *wallbitwallets.Service) *Service {
	return &Service{sdk: sdk}
}

func (s *Service) Get(ctx context.Context, req *wallbitwallets.GetRequest) (*wallbitwallets.GetResponse, error) {
	res, err := s.sdk.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}
