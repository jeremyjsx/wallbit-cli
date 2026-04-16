package trades

import (
	"context"

	wallbitclient "github.com/jeremyjsx/wallbit-go/client"
	wallbittrades "github.com/jeremyjsx/wallbit-go/services/trades"
)

type ClientProvider func() (*wallbitclient.Client, error)

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
		req.TimeInForce = input.TimeInForce
	}
	return c.Trades.Create(ctx, req)
}
