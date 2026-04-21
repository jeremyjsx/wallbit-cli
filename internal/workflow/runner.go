package workflow

import (
	"context"
	"fmt"
	"time"

	"github.com/jeremyjsx/wallbit-cli/internal/services"
)

func Run(ctx context.Context, spec *Spec, svc *services.Services) *RunResult {
	start := time.Now()
	out := &RunResult{
		Workflow:  spec.Name,
		OK:        true,
		StartedAt: start,
		Steps:     make([]StepResult, 0, len(spec.Steps)),
	}

	for _, step := range spec.Steps {
		stepStart := time.Now()
		handler, ok := Registry[step.Run]
		if !ok {
			r := StepResult{
				ID:         step.ID,
				Run:        step.Run,
				OK:         false,
				Error:      &StepError{Message: fmt.Sprintf("unsupported run %q", step.Run)},
				DurationMS: time.Since(stepStart).Milliseconds(),
			}
			out.Steps = append(out.Steps, r)
			out.OK = false
			out.FailedStepID = step.ID
			if spec.OnError == OnErrorFailFast {
				break
			}
			continue
		}

		data, err := handler(ctx, svc, step.With)
		if err != nil {
			r := StepResult{
				ID:         step.ID,
				Run:        step.Run,
				OK:         false,
				Error:      &StepError{Message: err.Error()},
				DurationMS: time.Since(stepStart).Milliseconds(),
			}
			out.Steps = append(out.Steps, r)
			out.OK = false
			out.FailedStepID = step.ID
			if spec.OnError == OnErrorFailFast {
				break
			}
			continue
		}

		out.Steps = append(out.Steps, StepResult{
			ID:         step.ID,
			Run:        step.Run,
			OK:         true,
			Data:       data,
			DurationMS: time.Since(stepStart).Milliseconds(),
		})
	}

	out.FinishedAt = time.Now()
	return out
}
