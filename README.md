# Wallbit CLI

Command-line interface for the **Wallbit API**, built on the Go SDK [wallbit-go](https://github.com/jeremyjsx/wallbit-go). It covers balances, assets, fees, trades, YAML workflows, and more.

## Installation

**Current release: [v0.1.0](https://github.com/jeremyjsx/wallbit-cli/releases/tag/v0.1.0)**

Install a fixed version (recommended for reproducible installs):

```bash
go install github.com/jeremyjsx/wallbit-cli/cmd/wallbit@v0.1.0
```

Requires Go 1.23 or newer.

Install the latest commit on `main` from the Go module proxy:

```bash
go install github.com/jeremyjsx/wallbit-cli/cmd/wallbit@latest
```

**Build from a clone:**

```bash
go build -o wallbit ./cmd/wallbit
```

**Release binaries:** inject the reported version with `-X` on `internal/cli.Version`:

```bash
go build -ldflags "-X github.com/jeremyjsx/wallbit-cli/internal/cli.Version=v0.1.0" -o wallbit ./cmd/wallbit
```

## Authentication

API key resolution order:

1. Global **`--api-key`** flag  
2. **`WALLBIT_API_KEY`** environment variable  
3. Credential stored via **`wallbit auth login`**

```bash
wallbit auth login    # interactive; stores key in user config dir
wallbit auth status   # shows whether a key is configured (never prints it)
wallbit auth logout   # removes locally stored key
```

## Global flags

| Flag | Description |
|------|-------------|
| `--api-key` | API key (optional if env or login is set) |
| `--base-url` | API base URL (default `https://api.wallbit.io`) |
| `--timeout` | HTTP client timeout (default `30s`) |
| `--version`, `-v` | Binary version |

After `go install` at tag **`v0.1.0`** (or with `-ldflags` as above), `wallbit --version` prints that version. From a raw checkout without tags, it may show **`(devel)`** based on Go build metadata.

## Quick start

```bash
wallbit balance checking
wallbit workflow validate workflow.yaml
wallbit workflow run workflow.yaml
wallbit --help
```

Top-level commands include `auth`, `balance`, `assets`, `rates`, `fees`, `trades`, `transactions`, `wallets`, `cards`, `operations`, `roboadvisor`, `account-details`, `apikey`, and `workflow`. Run `wallbit help` for the full command tree.

## Workflows (YAML)

Workflow steps use `run` keys implemented by this CLI (balances, fees, operations, etc.). Validate before running:

```bash
wallbit workflow validate my-flow.yaml
```

## Related

| Resource | Link |
|----------|------|
| Go SDK | [jeremyjsx/wallbit-go](https://github.com/jeremyjsx/wallbit-go) |
| Releases | [wallbit-cli releases](https://github.com/jeremyjsx/wallbit-cli/releases) |
