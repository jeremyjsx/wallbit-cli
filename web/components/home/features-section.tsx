import Link from "next/link";

const featureBlocks = [
  {
    tag: "BINARY",
    title: "Built for humans and agents",
    body:
      "One interface for exploratory commands and repeatable workflows. You and your agent operate on the same contract.",
    href: "/docs#cmd-overview",
    cta: "See Command Overview",
  },
  {
    tag: "WORKFLOWS",
    title: "YAML that executes, not just documents",
    body:
      "Turn intent into executable plans with `${steps.<id>.<path>}`. Compose once, run many times, and keep shell glue out of mission-critical flows.",
    href: "/docs#workflow-language",
    cta: "See Workflow Language",
  },
  {
    tag: "VALIDATE",
    title: "Validate before money moves",
    body:
      "`wallbit workflow validate` checks run keys and payload structure up front, so agents catch broken plans before execution.",
    href: "/docs#workflow-validate-run",
    cta: "See Validate & Run",
  },
  {
    tag: "JSON",
    title: "Automation-ready output",
    body:
      "Structured JSON on success, strict non-zero exits on failure. Easy to pipe, test, monitor, and chain in automation pipelines.",
    href: "/docs#output",
    cta: "See Output Format",
  },
  {
    tag: "AUTH",
    title: "Predictable auth resolution",
    body:
      "Clear credential precedence keeps runs deterministic across environments. See effective auth state instantly with `auth status`.",
    href: "/docs#authentication",
    cta: "See Authentication",
  },
  {
    tag: "SDK",
    title: "Aligned with the public API",
    body:
      "Typed requests and domain-aware commands keep execution close to API behavior, with knobs for base URL and timeout when you need control.",
    href: "/docs#cmd-overview",
    cta: "See API Mapping",
  },
] as const;

export function FeaturesSection() {
  return (
    <section id="features" className="py-16 sm:py-24">
      <div className="mx-auto max-w-6xl px-4 sm:px-6">
        <div className="mx-auto max-w-2xl text-center">
          <p className="font-mono text-[10px] font-semibold uppercase tracking-[0.2em] text-wallbit-500/80">Features</p>
          <h2 className="mt-2 text-2xl font-semibold tracking-tight text-zinc-100 sm:text-3xl">
            Why teams trust wallbit-cli
          </h2>
          <p className="mt-3 text-zinc-500">
            From prompt to production run: design workflows with your favorite agent, validate them with the CLI, and execute with
            confidence.
          </p>
        </div>
        <ul className="mt-12 grid gap-5 sm:grid-cols-2 lg:grid-cols-3">
          {featureBlocks.map((f) => (
            <li key={f.tag} className="feature-card group relative overflow-hidden rounded-2xl p-px">
              <div className="feature-card-borderbox" aria-hidden>
                <span className="feature-card-ring" />
              </div>
              <Link
                href={f.href}
                className="feature-card-inner relative z-[1] flex h-full flex-col border border-zinc-800/90 bg-zinc-950 p-6 no-underline transition-colors hover:border-zinc-700/90 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-wallbit-500/90 focus-visible:ring-offset-2 focus-visible:ring-offset-black sm:p-7"
              >
                <span className="font-mono text-[10px] font-semibold tracking-widest text-wallbit-500/90">{f.tag}</span>
                <h3 className="mt-3 text-base font-semibold text-zinc-100">{f.title}</h3>
                <p className="mt-2 flex-1 text-sm leading-relaxed text-zinc-500">{f.body}</p>
                <span className="mt-4 text-xs text-wallbit-400/90">{f.cta} →</span>
              </Link>
            </li>
          ))}
        </ul>
      </div>
    </section>
  );
}
