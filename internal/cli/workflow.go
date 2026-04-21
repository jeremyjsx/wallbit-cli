package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/jeremyjsx/wallbit-cli/internal/workflow"
	"github.com/spf13/cobra"
)

var workflowCmd = &cobra.Command{
	Use:   "workflow",
	Short: "Run AI-first YAML workflows",
}

var workflowRunCmd = &cobra.Command{
	Use:   "run <file.yaml>",
	Short: "Run a workflow YAML file",
	Args:  cobra.ExactArgs(1),
	RunE:  runWorkflowRun,
}

var workflowValidateCmd = &cobra.Command{
	Use:   "validate <file.yaml>",
	Short: "Validate a workflow YAML file without running it",
	Args:  cobra.ExactArgs(1),
	RunE:  runWorkflowValidate,
}

func init() {
	workflowCmd.AddCommand(workflowRunCmd, workflowValidateCmd)
	rootCmd.AddCommand(workflowCmd)
}

func runWorkflowRun(cmd *cobra.Command, args []string) error {
	data, err := os.ReadFile(args[0])
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	spec, err := workflow.ParseSpec(data)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	if err := workflow.ValidateSupportedRuns(spec); err != nil {
		return fmt.Errorf("%w", err)
	}

	svc, err := app.Services()
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	ctx, cancel := context.WithTimeout(cmd.Context(), app.Timeout())
	defer cancel()

	out := workflow.Run(ctx, spec, svc)

	enc := json.NewEncoder(cmd.OutOrStdout())
	enc.SetIndent("", "  ")
	return enc.Encode(out)
}

func runWorkflowValidate(cmd *cobra.Command, args []string) error {
	data, err := os.ReadFile(args[0])
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	spec, err := workflow.ParseSpec(data)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	if err := workflow.ValidateSupportedRuns(spec); err != nil {
		return fmt.Errorf("%w", err)
	}

	enc := json.NewEncoder(cmd.OutOrStdout())
	enc.SetIndent("", "  ")
	return enc.Encode(map[string]any{
		"ok":      true,
		"name":    spec.Name,
		"steps":   len(spec.Steps),
		"version": spec.Version,
	})
}
