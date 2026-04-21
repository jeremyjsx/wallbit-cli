package transactions

import (
	"context"
	"time"

	wallbittx "github.com/jeremyjsx/wallbit-go/services/transactions"
)

type Service struct {
	sdk *wallbittx.Service
}

func New(sdk *wallbittx.Service) *Service {
	return &Service{sdk: sdk}
}

type ListInput struct {
	Page       *int
	Limit      *int
	Status     string
	Type       string
	Currency   string
	FromDate   *time.Time
	ToDate     *time.Time
	FromAmount *float64
	ToAmount   *float64
}

func (s *Service) List(ctx context.Context, input *ListInput) (*wallbittx.ListResponse, error) {
	req := &wallbittx.ListRequest{}
	if input != nil {
		req.Page = input.Page
		req.Limit = input.Limit
		req.Status = input.Status
		req.Type = input.Type
		req.Currency = input.Currency
		req.FromDate = input.FromDate
		req.ToDate = input.ToDate
		req.FromAmount = input.FromAmount
		req.ToAmount = input.ToAmount
	}
	res, err := s.sdk.List(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}
