package apikey

import (
	"context"

	wallbitapikey "github.com/jeremyjsx/wallbit-go/services/apikey"
	"github.com/jeremyjsx/wallbit-go/wallbit"
)

type ClientProvider func() (*wallbit.Client, error)

type Service struct {
	clientProvider ClientProvider
}

func New(clientProvider ClientProvider) *Service {
	return &Service{clientProvider: clientProvider}
}

func (s *Service) Revoke(ctx context.Context) (*wallbitapikey.RevokeResponse, error) {
	c, err := s.clientProvider()
	if err != nil {
		return nil, err
	}
	res, err := c.APIKey.Revoke(ctx)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}
