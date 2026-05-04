package workflow

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

const (
	SpecVersion1    = 1
	OnErrorFailFast = "fail_fast"
	OnErrorContinue = "continue"
)

type Spec struct {
	Version int    `yaml:"version"`
	Name    string `yaml:"name"`
	OnError string `yaml:"on_error"`
	Steps   []Step `yaml:"steps"`
}

type Step struct {
	ID   string         `yaml:"id"`
	Run  string         `yaml:"run"`
	With map[string]any `yaml:"with"`
}

func ParseSpec(data []byte) (*Spec, error) {
	var spec Spec
	if err := yaml.Unmarshal(data, &spec); err != nil {
		return nil, err
	}
	if err := spec.Validate(); err != nil {
		return nil, err
	}
	return &spec, nil
}

func (s *Spec) Validate() error {
	if s.Version != SpecVersion1 {
		return fmt.Errorf("unsupported workflow version %d", s.Version)
	}
	if len(s.Steps) == 0 {
		return fmt.Errorf("workflow requires at least one step")
	}
	if s.OnError == "" {
		s.OnError = OnErrorFailFast
	}
	if s.OnError != OnErrorFailFast && s.OnError != OnErrorContinue {
		return fmt.Errorf("invalid on_error %q", s.OnError)
	}
	seen := make(map[string]struct{}, len(s.Steps))
	for i, step := range s.Steps {
		if step.ID == "" {
			return fmt.Errorf("steps[%d].id is required", i)
		}
		if _, ok := seen[step.ID]; ok {
			return fmt.Errorf("duplicate step id %q", step.ID)
		}
		seen[step.ID] = struct{}{}
		if step.Run == "" {
			return fmt.Errorf("steps[%d].run is required", i)
		}
	}
	return nil
}

func ValidateSupportedRuns(s *Spec) error {
	for i, step := range s.Steps {
		if _, ok := Registry[step.Run]; !ok {
			return fmt.Errorf("steps[%d].run %q is not supported", i, step.Run)
		}
	}
	return nil
}

func ValidateStepInputs(s *Spec) error {
	for i, step := range s.Steps {
		validator, ok := InputValidators[step.Run]
		if !ok {
			continue
		}
		if err := validator(step.With); err != nil {
			return fmt.Errorf("steps[%d] (%s): %w", i, step.Run, err)
		}
	}
	return nil
}
