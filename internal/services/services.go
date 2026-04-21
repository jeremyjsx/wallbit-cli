package services

import (
	"github.com/jeremyjsx/wallbit-cli/internal/services/apikey"
	"github.com/jeremyjsx/wallbit-cli/internal/services/balance"
	"github.com/jeremyjsx/wallbit-cli/internal/services/fees"
	"github.com/jeremyjsx/wallbit-cli/internal/services/trades"
	"github.com/jeremyjsx/wallbit-cli/internal/services/transactions"
	"github.com/jeremyjsx/wallbit-go/wallbit"
)

// Services wires every domain wrapper to the corresponding SDK sub-client.
// Construct with [New] after you have a *wallbit.Client (one client, one graph).
type Services struct {
	APIKey       *apikey.Service
	Balance      *balance.Service
	Fees         *fees.Service
	Trades       *trades.Service
	Transactions *transactions.Service
}

func New(c *wallbit.Client) *Services {
	return &Services{
		APIKey:       apikey.New(c.APIKey),
		Balance:      balance.New(c.Balance),
		Fees:         fees.New(c.Fees),
		Trades:       trades.New(c.Trades),
		Transactions: transactions.New(c.Transactions),
	}
}
