package assets

import (
	"context"

	wallbitassets "github.com/jeremyjsx/wallbit-go/services/assets"
)

type Service struct {
	sdk *wallbitassets.Service
}

func New(sdk *wallbitassets.Service) *Service {
	return &Service{sdk: sdk}
}

func (s *Service) List(ctx context.Context, req *wallbitassets.ListRequest) (*wallbitassets.ListResponse, error) {
	res, err := s.sdk.List(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

func (s *Service) Get(ctx context.Context, symbol string) (*wallbitassets.GetResponse, error) {
	res, err := s.sdk.Get(ctx, symbol)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}
