package cli

import (
	"runtime/debug"
	"strings"
)

const modulePath = "github.com/jeremyjsx/wallbit-cli"

// Version: set at link time on release bundles, e.g.
//
//	go build -ldflags "-X github.com/jeremyjsx/wallbit-cli/internal/cli.Version=v0.2.0" ./cmd/wallbit
//
// When unset, formatBuildVersion uses debug.ReadBuildInfo (tags / pseudo-versions, or "(devel)" from the toolchain).
var Version = ""

func resolveVersion() string {
	if v := strings.TrimSpace(Version); v != "" {
		return v
	}
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "(unknown)"
	}
	if info.Main.Path == modulePath && info.Main.Version != "" {
		return info.Main.Version
	}
	return "(unknown)"
}

func formatBuildVersion() string {
	return resolveVersion()
}
