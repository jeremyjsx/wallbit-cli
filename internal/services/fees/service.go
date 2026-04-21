package fees

import (
	"context"

	wallbitfees "github.com/jeremyjsx/wallbit-go/services/fees"
	"github.com/jeremyjsx/wallbit-go/wallbit"
)

type ClientProvider func() (*wallbit.Client, error)

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
	res, err := c.Fees.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}
