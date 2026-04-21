package workflow

import (
	"context"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/jeremyjsx/wallbit-cli/internal/services"
)

var stepRefPattern = regexp.MustCompile(`\$\{steps\.([a-zA-Z0-9_\-]+)\.([^\}]+)\}`)

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

		withResolved, err := resolveStepWith(step.With, out.Steps)
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

		data, err := handler(ctx, svc, withResolved)
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

func resolveStepWith(with map[string]any, prior []StepResult) (map[string]any, error) {
	if with == nil {
		return nil, nil
	}
	resolved := make(map[string]any, len(with))
	for k, v := range with {
		out, err := resolveAny(v, prior)
		if err != nil {
			return nil, err
		}
		resolved[k] = out
	}
	return resolved, nil
}

func resolveAny(v any, prior []StepResult) (any, error) {
	switch t := v.(type) {
	case string:
		return resolveStringRefs(t, prior)
	case map[string]any:
		out := make(map[string]any, len(t))
		for k, inner := range t {
			r, err := resolveAny(inner, prior)
			if err != nil {
				return nil, err
			}
			out[k] = r
		}
		return out, nil
	case []any:
		out := make([]any, 0, len(t))
		for _, inner := range t {
			r, err := resolveAny(inner, prior)
			if err != nil {
				return nil, err
			}
			out = append(out, r)
		}
		return out, nil
	default:
		return v, nil
	}
}

func resolveStringRefs(s string, prior []StepResult) (any, error) {
	matches := stepRefPattern.FindAllStringSubmatchIndex(s, -1)
	if len(matches) == 0 {
		return s, nil
	}
	if len(matches) == 1 && matches[0][0] == 0 && matches[0][1] == len(s) {
		stepID := s[matches[0][2]:matches[0][3]]
		path := s[matches[0][4]:matches[0][5]]
		return lookupStepRef(prior, stepID, path)
	}

	var b strings.Builder
	last := 0
	for _, m := range matches {
		b.WriteString(s[last:m[0]])
		stepID := s[m[2]:m[3]]
		path := s[m[4]:m[5]]
		val, err := lookupStepRef(prior, stepID, path)
		if err != nil {
			return nil, err
		}
		b.WriteString(fmt.Sprint(val))
		last = m[1]
	}
	b.WriteString(s[last:])
	return b.String(), nil
}

func lookupStepRef(prior []StepResult, stepID string, path string) (any, error) {
	var step *StepResult
	for i := range prior {
		if prior[i].ID == stepID {
			step = &prior[i]
			break
		}
	}
	if step == nil {
		return nil, fmt.Errorf("reference step %q not found", stepID)
	}
	parts := strings.Split(path, ".")
	var cur any = step
	for _, p := range parts {
		next, ok := descend(cur, p)
		if !ok {
			return nil, fmt.Errorf("reference steps.%s.%s not found", stepID, path)
		}
		cur = next
	}
	return cur, nil
}

func descend(v any, key string) (any, bool) {
	if v == nil {
		return nil, false
	}
	rv := reflect.ValueOf(v)
	for rv.Kind() == reflect.Pointer {
		if rv.IsNil() {
			return nil, false
		}
		rv = rv.Elem()
	}
	switch rv.Kind() {
	case reflect.Map:
		if rv.Type().Key().Kind() != reflect.String {
			return nil, false
		}
		mv := rv.MapIndex(reflect.ValueOf(key))
		if !mv.IsValid() {
			return nil, false
		}
		return mv.Interface(), true
	case reflect.Struct:
		fv := rv.FieldByNameFunc(func(name string) bool { return strings.EqualFold(name, key) })
		if !fv.IsValid() || !fv.CanInterface() {
			return nil, false
		}
		return fv.Interface(), true
	default:
		return nil, false
	}
}
