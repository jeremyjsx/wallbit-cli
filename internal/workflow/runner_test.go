package workflow

import (
	"context"
	"errors"
	"testing"

	"github.com/jeremyjsx/wallbit-cli/internal/services"
)

func TestRunFailFastStopsAfterFirstError(t *testing.T) {
	orig := Registry
	t.Cleanup(func() { Registry = orig })

	Registry = map[string]StepHandler{
		"ok.step": func(ctx context.Context, svc *services.Services, with map[string]any) (any, error) {
			return map[string]any{"ok": true}, nil
		},
		"bad.step": func(ctx context.Context, svc *services.Services, with map[string]any) (any, error) {
			return nil, errors.New("boom")
		},
	}

	spec := &Spec{
		Version: 1,
		Name:    "runner-fail-fast",
		OnError: OnErrorFailFast,
		Steps: []Step{
			{ID: "s1", Run: "ok.step"},
			{ID: "s2", Run: "bad.step"},
			{ID: "s3", Run: "ok.step"},
		},
	}

	out := Run(context.Background(), spec, nil)
	if out.OK {
		t.Fatalf("expected overall failure")
	}
	if out.FailedStepID != "s2" {
		t.Fatalf("expected failed step s2, got %q", out.FailedStepID)
	}
	if len(out.Steps) != 2 {
		t.Fatalf("expected 2 executed steps in fail-fast mode, got %d", len(out.Steps))
	}
}

func TestRunContinueExecutesRemainingSteps(t *testing.T) {
	orig := Registry
	t.Cleanup(func() { Registry = orig })

	Registry = map[string]StepHandler{
		"ok.step": func(ctx context.Context, svc *services.Services, with map[string]any) (any, error) {
			return map[string]any{"ok": true}, nil
		},
		"bad.step": func(ctx context.Context, svc *services.Services, with map[string]any) (any, error) {
			return nil, errors.New("boom")
		},
	}

	spec := &Spec{
		Version: 1,
		Name:    "runner-continue",
		OnError: OnErrorContinue,
		Steps: []Step{
			{ID: "s1", Run: "bad.step"},
			{ID: "s2", Run: "ok.step"},
		},
	}

	out := Run(context.Background(), spec, nil)
	if out.OK {
		t.Fatalf("expected overall failure due to one failed step")
	}
	if len(out.Steps) != 2 {
		t.Fatalf("expected 2 executed steps in continue mode, got %d", len(out.Steps))
	}
	if !out.Steps[1].OK {
		t.Fatalf("expected second step to run and pass")
	}
}

func TestRunResolvesStepReferences(t *testing.T) {
	orig := Registry
	t.Cleanup(func() { Registry = orig })

	Registry = map[string]StepHandler{
		"produce.step": func(ctx context.Context, svc *services.Services, with map[string]any) (any, error) {
			return map[string]any{
				"value": "AAPL",
			}, nil
		},
		"consume.step": func(ctx context.Context, svc *services.Services, with map[string]any) (any, error) {
			got, _ := with["symbol"].(string)
			return map[string]any{"symbol": got}, nil
		},
	}

	spec := &Spec{
		Version: 1,
		Name:    "runner-refs",
		OnError: OnErrorFailFast,
		Steps: []Step{
			{ID: "s1", Run: "produce.step"},
			{
				ID:  "s2",
				Run: "consume.step",
				With: map[string]any{
					"symbol": "${steps.s1.data.value}",
				},
			},
		},
	}

	out := Run(context.Background(), spec, nil)
	if !out.OK {
		t.Fatalf("expected success, got failure: %+v", out)
	}
	if len(out.Steps) != 2 {
		t.Fatalf("expected 2 steps, got %d", len(out.Steps))
	}
	data, _ := out.Steps[1].Data.(map[string]any)
	if got := data["symbol"]; got != "AAPL" {
		t.Fatalf("expected resolved symbol AAPL, got %#v", got)
	}
}

func TestRunReferenceMissingStepFails(t *testing.T) {
	orig := Registry
	t.Cleanup(func() { Registry = orig })

	Registry = map[string]StepHandler{
		"consume.step": func(ctx context.Context, svc *services.Services, with map[string]any) (any, error) {
			return map[string]any{"ok": true}, nil
		},
	}

	spec := &Spec{
		Version: 1,
		Name:    "runner-missing-ref",
		OnError: OnErrorFailFast,
		Steps: []Step{
			{
				ID:  "s1",
				Run: "consume.step",
				With: map[string]any{
					"symbol": "${steps.unknown.data.value}",
				},
			},
		},
	}

	out := Run(context.Background(), spec, nil)
	if out.OK {
		t.Fatalf("expected failure for missing reference")
	}
	if out.FailedStepID != "s1" {
		t.Fatalf("expected failed step s1, got %q", out.FailedStepID)
	}
}
