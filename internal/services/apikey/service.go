package apikey

import (
	"context"

	wallbitapikey "github.com/jeremyjsx/wallbit-go/services/apikey"
)

type Service struct {
	sdk *wallbitapikey.Service
}

func New(sdk *wallbitapikey.Service) *Service {
	return &Service{sdk: sdk}
}

func (s *Service) Revoke(ctx context.Context) (*wallbitapikey.RevokeResponse, error) {
	res, err := s.sdk.Revoke(ctx)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}
