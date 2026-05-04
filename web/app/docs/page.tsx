import Link from "next/link";
import { Fragment } from "react";
import { WallbitMark } from "../../components/site/wallbit-mark";
import { DocsSidebar } from "./docs-sidebar";

const GITHUB = "https://github.com/jeremyjsx/wallbit-cli";

function CodeBlock({ children }: { children: string }) {
  return (
    <pre className="my-4 overflow-x-auto rounded-xl border border-zinc-800 border-l-2 border-l-wallbit-500/55 bg-[var(--panel)] p-4 text-[13px] leading-relaxed text-zinc-300">
      <code>{children}</code>
    </pre>
  );
}

function H2({ id, children }: { id: string; children: React.ReactNode }) {
  return (
    <h2
      id={id}
      className="scroll-mt-24 border-b border-wallbit-500/35 pb-2 text-xl font-semibold tracking-tight text-wallbit-400"
    >
      {children}
    </h2>
  );
}

function H3({ id, children }: { id?: string; children: React.ReactNode }) {
  return (
    <h3 id={id} className="mt-8 scroll-mt-24 text-base font-semibold text-wallbit-300">
      {children}
    </h3>
  );
}

function Table({
  headers,
  rows,
}: {
  headers: string[];
  rows: React.ReactNode[][];
}) {
  return (
    <div className="my-4 overflow-x-auto rounded-xl border border-zinc-800">
      <table className="w-full min-w-[28rem] border-collapse text-left text-sm">
        <thead>
          <tr className="border-b border-zinc-800 bg-zinc-900/50">
            {headers.map((h) => (
              <th key={h} className="px-3 py-2 font-medium text-wallbit-400">
                {h}
              </th>
            ))}
          </tr>
        </thead>
        <tbody className="text-zinc-400">
          {rows.map((cells, i) => (
            <tr key={i} className="border-b border-zinc-800/80 last:border-0">
              {cells.map((c, j) => (
                <td key={j} className="px-3 py-2 align-top">
                  {c}
                </td>
              ))}
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

export default function DocsPage() {
  return (
    <div className="min-h-dvh text-zinc-300">
      <header className="sticky top-0 z-40 border-b border-zinc-800/80 bg-[var(--background)]/90 backdrop-blur-md">
        <div className="mx-auto flex h-14 max-w-6xl items-center justify-between gap-4 px-4 sm:px-6">
          <div className="flex items-center gap-3 sm:gap-4">
            <Link
              href="/"
              aria-label="Wallbit CLI — home"
              className="inline-flex items-center gap-2 rounded-md no-underline focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-wallbit-500/90 focus-visible:ring-offset-2 focus-visible:ring-offset-black"
            >
              <WallbitMark className="h-[18px] w-auto text-wallbit-500" />
              <span className="text-sm font-semibold tracking-tight text-zinc-50">wallbit-cli</span>
            </Link>
            <span aria-hidden className="hidden text-zinc-700 sm:inline">/</span>
            <span className="hidden text-sm text-wallbit-500/85 sm:inline">Documentation</span>
          </div>
          <div className="flex items-center gap-3 text-sm">
            <Link href="/#features" className="hidden text-zinc-500 no-underline hover:text-zinc-200 sm:inline">
              Home
            </Link>
            <a href={GITHUB} target="_blank" rel="noreferrer noopener" className="text-zinc-500 no-underline hover:text-zinc-200">
              Source
            </a>
          </div>
        </div>
      </header>

      <div className="mx-auto flex max-w-6xl flex-col gap-8 px-4 py-8 lg:flex-row lg:gap-12 lg:px-6 lg:py-10">
        <DocsSidebar />

        <article className="min-w-0 flex-1 pb-24">
          <h1 className="mb-2 bg-gradient-to-r from-wallbit-300 via-wallbit-400 to-wallbit-500 bg-clip-text text-3xl font-semibold tracking-tight text-transparent">
            Documentation
          </h1>
          <p className="mb-10 max-w-2xl text-zinc-500">
            wallbit-cli is built workflow-first: define flows in YAML, validate before execution, and run repeatable
            financial automations from your terminal.
          </p>

          <section className="space-y-6">
            <H2 id="prerequisites">Prerequisites</H2>
            <ul className="list-disc space-y-1 pl-5 text-zinc-400">
              <li>Go toolchain (to install with <code className="text-wallbit-300/90">go install</code>)</li>
              <li>A Wallbit API key</li>
              <li>Network access to your Wallbit API host (default in this build is the dev API base URL unless you override it)</li>
            </ul>

            <H2 id="installation">Installation</H2>
            <p className="text-zinc-400">
              Install the <code className="text-zinc-300">wallbit</code> binary from the repository module path:
            </p>
            <CodeBlock>{`go install github.com/jeremyjsx/wallbit-cli/cmd/wallbit@latest`}</CodeBlock>
            <p className="text-zinc-400">
              Ensure <code className="text-zinc-300">$(go env GOPATH)/bin</code> (or your chosen install location) is on your{" "}
              <code className="text-zinc-300">PATH</code>.
            </p>

            <H2 id="authentication">Authentication</H2>
            <p className="text-zinc-400">
              The CLI resolves an API key in this order: <strong className="text-zinc-300">--api-key</strong> flag, then{" "}
              <strong className="text-zinc-300">WALLBIT_API_KEY</strong>, then the local credentials file written by{" "}
              <code className="text-zinc-300">wallbit auth login</code>.
            </p>
            <Table
              headers={["Command", "Description"]}
              rows={[
                [
                  <code key="a" className="text-wallbit-300/90">wallbit auth login</code>,
                  "Prompts for an API key (unless --api-key is set) and saves it to the local store.",
                ],
                [
                  <code key="b" className="text-wallbit-300/90">wallbit auth status</code>,
                  "Shows whether env/file/flag is configured. Never prints the key.",
                ],
                [
                  <code key="c" className="text-wallbit-300/90">wallbit auth logout</code>,
                  "Removes the locally stored API key.",
                ],
              ]}
            />

            <H2 id="global-flags">Global flags</H2>
            <p className="text-zinc-400">These flags are available on the root command and inherited by subcommands.</p>
            <Table
              headers={["Flag", "Description"]}
              rows={[
                [
                  <code key="1" className="text-wallbit-300/90">--api-key</code>,
                  "Wallbit API key (overrides env and stored credentials).",
                ],
                [
                  <code key="2" className="text-wallbit-300/90">--base-url</code>,
                  "Wallbit API base URL. Default matches the CLI build (currently the dev API URL unless you change the code).",
                ],
                [
                  <code key="3" className="text-wallbit-300/90">--timeout</code>,
                  "HTTP client timeout (default 30s).",
                ],
              ]}
            />

            <H2 id="cli-reference">CLI reference</H2>
            <p className="text-zinc-400">
              All successful commands write JSON to stdout unless noted. Errors are printed to stderr and exit non-zero.
            </p>

            <H3 id="cmd-auth">auth</H3>
            <p className="text-zinc-400">See Authentication above. Subcommands: <code className="text-zinc-300">login</code>,{" "}
              <code className="text-zinc-300">status</code>, <code className="text-zinc-300">logout</code>.</p>

            <H3 id="cmd-overview">Command overview</H3>
            <p className="text-zinc-400">
              Most commands are thin wrappers over one API domain. Use this table as a fast map, then jump to workflows
              when you need multi-step execution.
            </p>
            <Table
              headers={["Group", "What it covers", "Typical command"]}
              rows={[
                ["balance", "Checking and stock balances", <code key="bal" className="text-wallbit-300/90">wallbit balance checking</code>],
                ["rates", "Currency exchange quotes", <code key="rat" className="text-wallbit-300/90">wallbit rates get --source USD --dest EUR</code>],
                ["wallets", "Wallet discovery/filtering", <code key="wal" className="text-wallbit-300/90">wallbit wallets get --currency USDC</code>],
                ["assets", "Asset listing and details", <code key="ass" className="text-wallbit-300/90">wallbit assets list --limit 5</code>],
                ["transactions", "Transaction feeds and filters", <code key="tx" className="text-wallbit-300/90">wallbit transactions list --limit 10</code>],
                ["cards", "List, block, unblock card operations", <code key="car" className="text-wallbit-300/90">wallbit cards block &lt;card-uuid&gt;</code>],
                ["roboadvisor", "Portfolio balance, deposit, withdraw", <code key="rob" className="text-wallbit-300/90">wallbit roboadvisor balance</code>],
                ["fees / apikey", "Fee lookup and key revocation", <code key="fee" className="text-wallbit-300/90">wallbit fees get --type TRADE</code>],
              ]}
            />

            <H3 id="cmd-trades">trades</H3>
            <p className="text-zinc-400 mb-2">
              <code className="text-wallbit-300/90">wallbit trades create</code> — Required:{" "}
              <code className="text-zinc-300">--symbol</code>, <code className="text-zinc-300">--direction</code> (BUY or SELL),{" "}
              <code className="text-zinc-300">--order-type</code> (MARKET, LIMIT, STOP, STOP_LIMIT). Exactly one of{" "}
              <code className="text-zinc-300">--amount</code> or <code className="text-zinc-300">--shares</code>. Optional:{" "}
              <code className="text-zinc-300">--currency</code> (default USD), <code className="text-zinc-300">--stop-price</code>,{" "}
              <code className="text-zinc-300">--limit-price</code>, <code className="text-zinc-300">--time-in-force</code> (DAY or GTC for LIMIT).
            </p>
            <p className="text-sm text-zinc-600">
              LIMIT requires <code className="text-zinc-500">--limit-price</code> and <code className="text-zinc-500">--time-in-force</code>. STOP requires{" "}
              <code className="text-zinc-500">--stop-price</code>. STOP_LIMIT requires both stop and limit prices.
            </p>

            <p className="text-zinc-400">
              For the remaining groups (<code className="text-zinc-300">roboadvisor</code>, <code className="text-zinc-300">fees</code>,{" "}
              <code className="text-zinc-300">apikey</code>), prefer workflow handlers when available and use direct
              commands for operational one-offs.
            </p>

            <H3 id="environment">Environment</H3>
            <Table
              headers={["Variable", "Description"]}
              rows={[
                [
                  <code key="e" className="text-wallbit-300/90">WALLBIT_API_KEY</code>,
                  "API key used when no --api-key flag and no saved credentials (or as configured by the credentials package).",
                ],
              ]}
            />

            <H2 id="workflow-quickstart">First workflow in 5 minutes</H2>
            <p className="text-zinc-400">
              Your fastest path is: create a YAML file, validate it, then run it. Use this as a copy-paste baseline:
            </p>
            <CodeBlock>{`version: 1
name: quickstart-fx
steps:
  - id: fx
    run: rates.get
    with:
      source: USD
      dest: EUR
  - id: wallets
    run: wallets.get
    with:
      currency: USDC
      network: polygon`}</CodeBlock>
            <CodeBlock>{`wallbit workflow validate ./quickstart.yaml
wallbit workflow run ./quickstart.yaml`}</CodeBlock>

            <H2 id="workflow-language">Workflow language</H2>
            <p className="text-zinc-400">
              Workflows are plain YAML with ordered <code className="text-zinc-300">steps</code>. Each step has an{" "}
              <code className="text-zinc-300">id</code>, a <code className="text-zinc-300">run</code> key, and optional{" "}
              <code className="text-zinc-300">with</code> inputs. Use step references to build multi-step logic:
              <code className="text-wallbit-300/90"> {"${steps.<step_id>.<dot.path>}"} </code>.
            </p>
            <Table
              headers={["Field", "Required", "Meaning"]}
              rows={[
                [<code key="v" className="text-wallbit-300/90">version</code>, "yes", "Workflow schema version (currently 1)."],
                [<code key="n" className="text-wallbit-300/90">name</code>, "optional", "Human-readable run name."],
                [<code key="oe" className="text-wallbit-300/90">on_error</code>, "optional", "fail_fast (default) or continue."],
                [<code key="s" className="text-wallbit-300/90">steps</code>, "yes", "Non-empty list of executable steps."],
              ]}
            />

            <H2 id="workflow-patterns">Patterns cookbook</H2>
            <H3>Pattern 1 - FX check plus reverse confirmation</H3>
            <CodeBlock>{`version: 1
name: fx-roundtrip
steps:
  - id: base_fx
    run: rates.get
    with:
      source: USD
      dest: EUR
  - id: reverse_fx
    run: rates.get
    with:
      source: \${steps.base_fx.data.Data.DestCurrency}
      dest: \${steps.base_fx.data.Data.SourceCurrency}`}</CodeBlock>

            <H3>Pattern 2 - Portfolio context plus transaction scan</H3>
            <CodeBlock>{`version: 1
name: portfolio-scan
steps:
  - id: wallets
    run: wallets.get
    with:
      currency: USDC
      network: polygon
  - id: tx
    run: transactions.list
    with:
      page: 1
      limit: 10
      currency: USD`}</CodeBlock>

            <H3>Pattern 3 - Multi-domain snapshot</H3>
            <CodeBlock>{`version: 1
name: morning-snapshot
steps:
  - id: assets
    run: assets.list
    with:
      page: 1
      limit: 5
      category: TECHNOLOGY
  - id: cards
    run: cards.list`}</CodeBlock>

            <H2 id="workflow-ai-prompts">AI prompt templates</H2>
            <p className="text-zinc-400">
              Give your copilot strict prompts so it emits runnable YAML instead of vague prose.
            </p>
            <CodeBlock>{`Generate a wallbit-cli workflow YAML (version 1) with:
- 3 steps
- one rates.get step (USD -> EUR)
- one wallets.get step (USDC on polygon)
- one transactions.list step (limit 10)
Return only valid YAML, no markdown, no explanations.`}</CodeBlock>

            <CodeBlock>{`Refactor this workflow to use step references where possible.
Keep run keys valid for wallbit-cli.
Return only YAML.

<paste-workflow-here>`}</CodeBlock>

            <H2 id="workflow-spec">YAML workflow specification</H2>
            <p className="text-zinc-400">Top-level fields:</p>
            <Table
              headers={["Field", "Required", "Description"]}
              rows={[
                [
                  <code key="v" className="text-wallbit-300/90">version</code>,
                  "yes",
                  "Must be 1.",
                ],
                [
                  <code key="n" className="text-wallbit-300/90">name</code>,
                  "optional",
                  "Workflow display name (included in validate JSON).",
                ],
                [
                  <code key="e" className="text-wallbit-300/90">on_error</code>,
                  "optional",
                  <>
                    <code className="text-zinc-300">fail_fast</code> (default) stops after the first failed step;{" "}
                    <code className="text-zinc-300">continue</code> runs subsequent steps.
                  </>,
                ],
                [
                  <code key="s" className="text-wallbit-300/90">steps</code>,
                  "yes",
                  "Non-empty array of steps.",
                ],
              ]}
            />
            <p className="text-zinc-400">Each step:</p>
            <Table
              headers={["Field", "Required", "Description"]}
              rows={[
                [
                  <code key="i" className="text-wallbit-300/90">id</code>,
                  "yes",
                  "Unique identifier within the workflow. Used for output references.",
                ],
                [
                  <code key="r" className="text-wallbit-300/90">run</code>,
                  "yes",
                  "Handler key (see Step catalog).",
                ],
                [
                  <code key="w" className="text-wallbit-300/90">with</code>,
                  "optional",
                  "Map of inputs for the handler (snake_case keys as documented per step).",
                ],
              ]}
            />
            <CodeBlock>{`version: 1
name: example
on_error: fail_fast
steps:
  - id: usd_eur
    run: rates.get
    with:
      source: USD
      dest: EUR
  - id: list_tx
    run: transactions.list
    with:
      limit: 20`}</CodeBlock>

            <H2 id="workflow-validate-run">Validate, Run & Output</H2>
            <p className="text-zinc-400">
              Both commands parse YAML, run structural validation (<code className="text-zinc-300">version</code>, unique{" "}
              <code className="text-zinc-300">id</code>, non-empty <code className="text-zinc-300">run</code>, valid{" "}
              <code className="text-zinc-300">on_error</code>), then <code className="text-zinc-300">ValidateSupportedRuns</code> and{" "}
              <code className="text-zinc-300">ValidateStepInputs</code>. <code className="text-zinc-300">run</code> additionally executes each step in order.
            </p>

            <H3 id="cmd-workflow">workflow</H3>
            <p className="text-zinc-400 mb-2">YAML workflow runner.</p>
            <ul className="list-disc space-y-1 pl-5 text-zinc-400">
              <li>
                <code className="text-wallbit-300/90">wallbit workflow validate &lt;file.yaml&gt;</code> — Parse YAML, ensure{" "}
                <code className="text-zinc-300">version</code> and structure, check every <code className="text-zinc-300">run</code>{" "}
                is supported, validate <code className="text-zinc-300">with</code> for steps that have validators. Prints a small JSON summary on success.
              </li>
              <li>
                <code className="text-wallbit-300/90">wallbit workflow run &lt;file.yaml&gt;</code> — Same validation as above, then executes steps in order and prints JSON{" "}
                <code className="text-zinc-300">RunResult</code> (see{" "}
                <a href="#output" className="text-wallbit-400/90 hover:text-wallbit-400">
                  Output format
                </a>
                ).
              </li>
            </ul>

            <H3 id="output">Output format</H3>
            <p className="text-zinc-400">
              Commands that call the API print indented JSON. <code className="text-zinc-300">workflow run</code> prints a{" "}
              <code className="text-zinc-300">RunResult</code>: <code className="text-zinc-300">workflow</code>, <code className="text-zinc-300">ok</code>,{" "}
              <code className="text-zinc-300">started_at</code>, <code className="text-zinc-300">finished_at</code>, optional{" "}
              <code className="text-zinc-300">failed_step_id</code>, and <code className="text-zinc-300">steps</code>. Each step has{" "}
              <code className="text-zinc-300">id</code>, <code className="text-zinc-300">run</code>, <code className="text-zinc-300">ok</code>, optional{" "}
              <code className="text-zinc-300">data</code> (API payload), optional <code className="text-zinc-300">error</code> with <code className="text-zinc-300">message</code>, and{" "}
              <code className="text-zinc-300">duration_ms</code>.
            </p>

            <H2 id="workflow-step-refs">Step output references</H2>
            <p className="text-zinc-400">
              In <code className="text-zinc-300">with</code> values you can interpolate prior step results using{" "}
              <code className="text-wallbit-300/90">{"${steps.<step_id>.<dot.path>}"}</code>. The path walks struct fields (case-insensitive) or map keys. If the entire{" "}
              <code className="text-zinc-300">with</code> value is a single reference, the resolved type is preserved; otherwise references are stringified into the template.
            </p>
            <p className="text-sm text-zinc-600">
              Only steps that completed successfully before the current step are visible; references must match a prior <code className="text-zinc-500">id</code>.
            </p>

            <H2 id="workflow-catalog">Workflow step catalog</H2>
            <p className="text-zinc-400 mb-2">
              Each row is a <code className="text-zinc-300">run</code> key. Fields listed under <strong className="text-zinc-300">with</strong> use workflow YAML naming (snake_case).
            </p>
            <Table
              headers={["run", "with", "Notes"]}
              rows={[
                [
                  <code key="rg" className="text-wallbit-300/90">rates.get</code>,
                  <Fragment key="rg-with">
                    <code className="text-zinc-500">source</code>, <code className="text-zinc-500">dest</code> (required)
                  </Fragment>,
                  "Both strings (currency codes).",
                ],
                [
                  <code key="bc" className="text-wallbit-300/90">balance.get_checking</code>,
                  "—",
                  "No inputs.",
                ],
                [
                  <code key="bs" className="text-wallbit-300/90">balance.get_stocks</code>,
                  "—",
                  "No inputs.",
                ],
                [
                  <code key="wg" className="text-wallbit-300/90">wallets.get</code>,
                  <Fragment key="wg-with">
                    optional <code className="text-zinc-500">currency</code>, <code className="text-zinc-500">network</code>
                  </Fragment>,
                  "Workflow runner lowercases network for the API request.",
                ],
                [
                  <code key="al" className="text-wallbit-300/90">assets.list</code>,
                  <Fragment key="al-with">
                    optional <code className="text-zinc-500">category</code>, <code className="text-zinc-500">search</code>, <code className="text-zinc-500">page</code>,{" "}
                    <code className="text-zinc-500">limit</code>
                  </Fragment>,
                  "Use category values accepted by your Wallbit environment.",
                ],
                [
                  <code key="ag" className="text-wallbit-300/90">assets.get</code>,
                  <code key="ag-sym" className="text-zinc-500">symbol</code>,
                  "Required.",
                ],
                [
                  <code key="ad" className="text-wallbit-300/90">account_details.get</code>,
                  <Fragment key="ad-with">
                    optional <code className="text-zinc-500">country</code>, <code className="text-zinc-500">currency</code>
                  </Fragment>,
                  "—",
                ],
                [
                  <code key="tl" className="text-wallbit-300/90">transactions.list</code>,
                  <Fragment key="tl-with">
                    optional <code className="text-zinc-500">status</code>, <code className="text-zinc-500">type</code>, <code className="text-zinc-500">currency</code>,{" "}
                    <code className="text-zinc-500">page</code>, <code className="text-zinc-500">limit</code>
                  </Fragment>,
                  "API may restrict allowed limit values.",
                ],
                [
                  <code key="cl" className="text-wallbit-300/90">cards.list</code>,
                  "—",
                  "No inputs.",
                ],
                [
                  <code key="cb" className="text-wallbit-300/90">cards.block</code>,
                  <code key="cb-id" className="text-zinc-500">card_uuid</code>,
                  "Required.",
                ],
                [
                  <code key="cu" className="text-wallbit-300/90">cards.unblock</code>,
                  <code key="cu-id" className="text-zinc-500">card_uuid</code>,
                  "Required.",
                ],
                [
                  <code key="tc" className="text-wallbit-300/90">trades.create</code>,
                  <Fragment key="tc-with">
                    <code className="text-zinc-500">symbol</code>, <code className="text-zinc-500">direction</code>, <code className="text-zinc-500">currency</code>,{" "}
                    <code className="text-zinc-500">order_type</code>; exactly one of <code className="text-zinc-500">amount</code> or <code className="text-zinc-500">shares</code>; optional{" "}
                    <code className="text-zinc-500">stop_price</code>, <code className="text-zinc-500">limit_price</code>, <code className="text-zinc-500">time_in_force</code>
                  </Fragment>,
                  "Same rules as the CLI trade validator.",
                ],
                [
                  <code key="rd" className="text-wallbit-300/90">roboadvisor.deposit</code>,
                  <Fragment key="rd-with">
                    <code className="text-zinc-500">robo_advisor_id</code>, <code className="text-zinc-500">amount</code>, <code className="text-zinc-500">from</code> (DEFAULT or INVESTMENT)
                  </Fragment>,
                  "Mutation.",
                ],
                [
                  <code key="rw" className="text-wallbit-300/90">roboadvisor.withdraw</code>,
                  <Fragment key="rw-with">
                    <code className="text-zinc-500">robo_advisor_id</code>, <code className="text-zinc-500">amount</code>, <code className="text-zinc-500">to</code> (DEFAULT or INVESTMENT)
                  </Fragment>,
                  "Mutation.",
                ],
                [
                  <code key="ar" className="text-wallbit-300/90">apikey.revoke</code>,
                  "—",
                  "Revokes the current API key. Dangerous in automation.",
                ],
              ]}
            />

            <p className="mt-12 text-center text-sm leading-relaxed text-zinc-600">
              Want more workflow recipes, integrations, or step patterns? Open an issue and share your use case so we can prioritize it in the docs and examples. You can also explore the{" "}
              <a href={GITHUB} target="_blank" rel="noreferrer noopener" className="text-zinc-500 hover:text-wallbit-400">
                repository
              </a>{" "}
              and create a request directly in{" "}
              <a href={`${GITHUB}/issues`} target="_blank" rel="noreferrer noopener" className="text-zinc-500 hover:text-wallbit-400">
                issues
              </a>
              .
            </p>
          </section>
        </article>
      </div>
    </div>
  );
}
