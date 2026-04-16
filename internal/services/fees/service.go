package fees

import (
	"context"

	wallbitclient "github.com/jeremyjsx/wallbit-go/client"
	wallbitfees "github.com/jeremyjsx/wallbit-go/services/fees"
)

type ClientProvider func() (*wallbitclient.Client, error)

type Service struct {
	clientProvider ClientProvider
}

func New(clientProvider ClientProvider) *Service {
	return &Service{clientProvider: clientProvider}
}

type GetInput struct {
	Type string
}

type GetResponse = wallbitfees.GetResponse

func (s *Service) Get(ctx context.Context, input *GetInput) (*GetResponse, error) {
	c, err := s.clientProvider()
	if err != nil {
		return nil, err
	}
	req := wallbitfees.GetRequest{}
	if input != nil {
		req.Type = input.Type
	}
	return c.Fees.Get(ctx, req)
}
