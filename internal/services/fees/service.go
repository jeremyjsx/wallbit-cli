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

type GetInput struct {
	Type string
}

type GetResponse = wallbitfees.GetResponse

func (s *Service) Get(ctx context.Context, input *GetInput) (*GetResponse, error) {
	req := wallbitfees.GetRequest{}
	if input != nil {
		req.Type = input.Type
	}
	res, err := s.sdk.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}
