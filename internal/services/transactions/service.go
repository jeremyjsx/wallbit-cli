package transactions

import (
	"context"
	"time"

	wallbitclient "github.com/jeremyjsx/wallbit-go/client"
	wallbittx "github.com/jeremyjsx/wallbit-go/services/transactions"
)

type ClientProvider func() (*wallbitclient.Client, error)

type Service struct {
	clientProvider ClientProvider
}

func New(clientProvider ClientProvider) *Service {
	return &Service{clientProvider: clientProvider}
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
	c, err := s.clientProvider()
	if err != nil {
		return nil, err
	}
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
	return c.Transactions.List(ctx, req)
}
