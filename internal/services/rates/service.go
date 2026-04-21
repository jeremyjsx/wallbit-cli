package rates

import (
	"context"

	wallbitrates "github.com/jeremyjsx/wallbit-go/services/rates"
)

type Service struct {
	sdk *wallbitrates.Service
}

func New(sdk *wallbitrates.Service) *Service {
	return &Service{sdk: sdk}
}

func (s *Service) Get(ctx context.Context, req wallbitrates.GetRequest) (*wallbitrates.GetResponse, error) {
	res, err := s.sdk.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}
