package cli

import (
	"testing"
)

func TestResolveVersion_ldflags(t *testing.T) {
	t.Cleanup(func() { Version = "" })

	Version = "  v1.2.3  "
	got := resolveVersion()
	if got != "v1.2.3" {
		t.Fatalf("got %q", got)
	}
}

func TestFormatBuildVersion_matchesResolve(t *testing.T) {
	t.Cleanup(func() { Version = "" })
	Version = "0.9.0"
	if formatBuildVersion() != resolveVersion() {
		t.Fatal("formatBuildVersion should mirror resolveVersion")
	}
}
