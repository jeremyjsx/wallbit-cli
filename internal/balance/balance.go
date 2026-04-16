package balance

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jeremyjsx/wallbit-cli/internal/client"
)

const pathChecking = "/api/public/v1/balance/checking"

type CheckingRow struct {
	Currency string  `json:"currency"`
	Balance  float64 `json:"balance"`
}

type CheckingResponse struct {
	Data []CheckingRow `json:"data"`
}

func GetChecking(ctx context.Context, c *client.Client) (*CheckingResponse, error) {
	body, _, err := c.Get(ctx, pathChecking)
	if err != nil {
		return nil, err
	}
	var out CheckingResponse
	if err := json.Unmarshal(body, &out); err != nil {
		return nil, fmt.Errorf("decode checking balance: %w", err)
	}
	return &out, nil
}
