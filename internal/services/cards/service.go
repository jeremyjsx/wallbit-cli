package cards

import (
	"context"

	wallbitcards "github.com/jeremyjsx/wallbit-go/services/cards"
)

type Service struct {
	sdk *wallbitcards.Service
}

func New(sdk *wallbitcards.Service) *Service {
	return &Service{sdk: sdk}
}

func (s *Service) List(ctx context.Context) (*wallbitcards.ListResponse, error) {
	res, err := s.sdk.List(ctx)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

func (s *Service) Block(ctx context.Context, cardUUID string) (*wallbitcards.UpdateStatusResponse, error) {
	res, err := s.sdk.Block(ctx, cardUUID)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

func (s *Service) Unblock(ctx context.Context, cardUUID string) (*wallbitcards.UpdateStatusResponse, error) {
	res, err := s.sdk.Unblock(ctx, cardUUID)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}
