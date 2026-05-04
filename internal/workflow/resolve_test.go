package workflow

import (
	"testing"

	wallbitrates "github.com/jeremyjsx/wallbit-go/services/rates"
)

func TestResolveStepWithRatesDataPath(t *testing.T) {
	prior := []StepResult{
		{
			ID: "fx_base",
			OK: true,
			Data: &wallbitrates.GetResponse{
				Data: wallbitrates.ExchangeRate{
					SourceCurrency: "USD",
					DestCurrency:   "EUR",
				},
			},
		},
	}
	with := map[string]any{
		"source": "${steps.fx_base.data.Data.DestCurrency}",
		"dest":   "${steps.fx_base.data.Data.SourceCurrency}",
	}
	resolved, err := resolveStepWith(with, prior)
	if err != nil {
		t.Fatalf("resolveStepWith: %v", err)
	}
	if resolved["source"] != "EUR" {
		t.Fatalf("source = %q want EUR", resolved["source"])
	}
	if resolved["dest"] != "USD" {
		t.Fatalf("dest = %q want USD", resolved["dest"])
	}
}

func TestResolveStepRefMissingStep(t *testing.T) {
	_, err := resolveStepWith(map[string]any{
		"x": "${steps.nope.data.Data.SourceCurrency}",
	}, []StepResult{})
	if err == nil {
		t.Fatal("expected error for missing step")
	}
}

func TestResolveStepRefFailedStep(t *testing.T) {
	prior := []StepResult{
		{ID: "fx_base", OK: false, Error: &StepError{Message: "boom"}},
	}
	_, err := resolveStepWith(map[string]any{
		"x": "${steps.fx_base.data.Data.SourceCurrency}",
	}, prior)
	if err == nil {
		t.Fatal("expected error when referenced step failed")
	}
}

func TestResolveStringPartialInterpolation(t *testing.T) {
	prior := []StepResult{
		{ID: "s1", OK: true, Data: map[string]any{"value": "X"}},
	}
	out, err := resolveAny(`prefix-${steps.s1.data.value}-suffix`, prior)
	if err != nil {
		t.Fatal(err)
	}
	if out != "prefix-X-suffix" {
		t.Fatalf("got %q", out)
	}
}
