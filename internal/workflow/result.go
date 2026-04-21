package workflow

import "time"

type RunResult struct {
	Workflow     string       `json:"workflow"`
	OK           bool         `json:"ok"`
	StartedAt    time.Time    `json:"started_at"`
	FinishedAt   time.Time    `json:"finished_at"`
	FailedStepID string       `json:"failed_step_id,omitempty"`
	Steps        []StepResult `json:"steps"`
}

type StepResult struct {
	ID         string     `json:"id"`
	Run        string     `json:"run"`
	OK         bool       `json:"ok"`
	Data       any        `json:"data,omitempty"`
	Error      *StepError `json:"error,omitempty"`
	DurationMS int64      `json:"duration_ms"`
}

type StepError struct {
	Message string `json:"message"`
}
