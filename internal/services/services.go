package services

import (
	"github.com/jeremyjsx/wallbit-cli/internal/services/accountdetails"
	"github.com/jeremyjsx/wallbit-cli/internal/services/apikey"
	"github.com/jeremyjsx/wallbit-cli/internal/services/assets"
	"github.com/jeremyjsx/wallbit-cli/internal/services/balance"
	"github.com/jeremyjsx/wallbit-cli/internal/services/cards"
	"github.com/jeremyjsx/wallbit-cli/internal/services/fees"
	"github.com/jeremyjsx/wallbit-cli/internal/services/rates"
	"github.com/jeremyjsx/wallbit-cli/internal/services/roboadvisor"
	"github.com/jeremyjsx/wallbit-cli/internal/services/trades"
	"github.com/jeremyjsx/wallbit-cli/internal/services/transactions"
	"github.com/jeremyjsx/wallbit-cli/internal/services/wallets"
	"github.com/jeremyjsx/wallbit-go/wallbit"
)

// Services wires every domain wrapper to the corresponding SDK sub-client.
// Construct with [New] after you have a *wallbit.Client (one client, one graph).
type Services struct {
	AccountDetails *accountdetails.Service
	APIKey       *apikey.Service
	Assets       *assets.Service
	Balance      *balance.Service
	Cards        *cards.Service
	Fees         *fees.Service
	Rates        *rates.Service
	RoboAdvisor  *roboadvisor.Service
	Trades       *trades.Service
	Transactions *transactions.Service
	Wallets      *wallets.Service
}

func New(c *wallbit.Client) *Services {
	return &Services{
		AccountDetails: accountdetails.New(c.AccountDetails),
		APIKey:       apikey.New(c.APIKey),
		Assets:       assets.New(c.Assets),
		Balance:      balance.New(c.Balance),
		Cards:        cards.New(c.Cards),
		Fees:         fees.New(c.Fees),
		Rates:        rates.New(c.Rates),
		RoboAdvisor:  roboadvisor.New(c.RoboAdvisor),
		Trades:       trades.New(c.Trades),
		Transactions: transactions.New(c.Transactions),
		Wallets:      wallets.New(c.Wallets),
	}
}
