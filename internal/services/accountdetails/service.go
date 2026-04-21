package accountdetails

import (
	"context"

	wallbitaccountdetails "github.com/jeremyjsx/wallbit-go/services/accountdetails"
)

type Service struct {
	sdk *wallbitaccountdetails.Service
}

func New(sdk *wallbitaccountdetails.Service) *Service {
	return &Service{sdk: sdk}
}

func (s *Service) Get(ctx context.Context, req *wallbitaccountdetails.GetRequest) (*wallbitaccountdetails.GetResponse, error) {
	res, err := s.sdk.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}
