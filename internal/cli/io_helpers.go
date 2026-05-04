package cli

import (
	"encoding/json"

	"github.com/spf13/cobra"
)

func writeJSON(out any, cmd *cobra.Command) error {
	enc := json.NewEncoder(cmd.OutOrStdout())
	enc.SetIndent("", "  ")
	return enc.Encode(out)
}
