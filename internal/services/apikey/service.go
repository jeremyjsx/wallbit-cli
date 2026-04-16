package apikey

import (
	"context"

	wallbitclient "github.com/jeremyjsx/wallbit-go/client"
	wallbitapikey "github.com/jeremyjsx/wallbit-go/services/apikey"
)

type ClientProvider func() (*wallbitclient.Client, error)

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
	return c.APIKey.Revoke(ctx)
}
