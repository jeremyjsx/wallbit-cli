package trades

import (
	"context"

	wallbittrades "github.com/jeremyjsx/wallbit-go/services/trades"
	"github.com/jeremyjsx/wallbit-go/wallbit"
)

type ClientProvider func() (*wallbit.Client, error)

type Service struct {
	clientProvider ClientProvider
}

func New(clientProvider ClientProvider) *Service {
	return &Service{clientProvider: clientProvider}
}

type CreateInput struct {
	Symbol      string
	Direction   string
	Currency    string
	OrderType   string
	Amount      *float64
	Shares      *float64
	StopPrice   *float64
	LimitPrice  *float64
	TimeInForce string
}

type CreateResponse = wallbittrades.CreateResponse

func (s *Service) Create(ctx context.Context, input *CreateInput) (*CreateResponse, error) {
	c, err := s.clientProvider()
	if err != nil {
		return nil, err
	}
	req := wallbittrades.CreateRequest{}
	if input != nil {
		req.Symbol = input.Symbol
		req.Direction = input.Direction
		req.Currency = input.Currency
		req.OrderType = input.OrderType
		req.Amount = input.Amount
		req.Shares = input.Shares
		req.StopPrice = input.StopPrice
		req.LimitPrice = input.LimitPrice
		if input.TimeInForce != "" {
			tif := input.TimeInForce
			req.TimeInForce = &tif
		}
	}
	res, err := c.Trades.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}
