package workflow

import "testing"

func TestParseSpecDefaultsOnError(t *testing.T) {
	spec, err := ParseSpec([]byte(`
version: 1
name: test
steps:
  - id: s1
    run: balance.get_checking
`))
	if err != nil {
		t.Fatalf("ParseSpec() error = %v", err)
	}
	if spec.OnError != OnErrorFailFast {
		t.Fatalf("expected default on_error %q, got %q", OnErrorFailFast, spec.OnError)
	}
}

func TestValidateSupportedRuns(t *testing.T) {
	spec := &Spec{
		Version: 1,
		Name:    "test",
		OnError: OnErrorFailFast,
		Steps: []Step{
			{ID: "s1", Run: "balance.get_checking"},
			{ID: "s2", Run: "unknown.run"},
		},
	}
	if err := ValidateSupportedRuns(spec); err == nil {
		t.Fatalf("expected unsupported run validation error")
	}
}

func TestValidateStepInputsTradesCreateRequiresExactlyOneAmountOrShares(t *testing.T) {
	spec := &Spec{
		Version: 1,
		Name:    "test",
		OnError: OnErrorFailFast,
		Steps: []Step{
			{
				ID:  "trade",
				Run: "trades.create",
				With: map[string]any{
					"symbol":     "AAPL",
					"direction":  "BUY",
					"currency":   "USD",
					"order_type": "MARKET",
				},
			},
		},
	}
	if err := ValidateStepInputs(spec); err == nil {
		t.Fatalf("expected trade input validation error when amount/shares missing")
	}

	spec.Steps[0].With["amount"] = 10.0
	spec.Steps[0].With["shares"] = 1.0
	if err := ValidateStepInputs(spec); err == nil {
		t.Fatalf("expected trade input validation error when both amount/shares present")
	}

	delete(spec.Steps[0].With, "shares")
	if err := ValidateStepInputs(spec); err != nil {
		t.Fatalf("expected valid input, got error: %v", err)
	}
}
