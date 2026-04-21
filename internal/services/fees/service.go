package fees

import (
	"context"

	wallbitfees "github.com/jeremyjsx/wallbit-go/services/fees"
)

type Service struct {
	sdk *wallbitfees.Service
}

func New(sdk *wallbitfees.Service) *Service {
	return &Service{sdk: sdk}
}

func (s *Service) Get(ctx context.Context, req wallbitfees.GetRequest) (*wallbitfees.GetResponse, error) {
	res, err := s.sdk.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}
