"use client";

import { useState } from "react";

const pillars = [
  {
    id: "define",
    title: "DEFINE STEPS",
    subtitle: "YAML engine",
    description: "Shape rebalance intent in ordered YAML steps.",
    accent: "from-wallbit-400/70 via-wallbit-500/25 to-transparent",
    Icon: IconBlueprint,
    featured: false,
  },
  {
    id: "validate",
    title: "CHECK INPUTS",
    subtitle: "Validation layer",
    description: "Verify structure, handlers, and step references before run.",
    accent: "from-wallbit-300/60 via-wallbit-500/20 to-transparent",
    Icon: IconShieldCheck,
    featured: false,
  },
  {
    id: "execute",
    title: "EXECUTE FLOW",
    subtitle: "Command runtime",
    description: "Run the validated plan and consume deterministic JSON output.",
    accent: "from-wallbit-500/70 via-wallbit-600/20 to-transparent",
    Icon: IconPlayResults,
    featured: false,
  },
  {
    id: "automate",
    title: "AGENT AUTOMATION",
    subtitle: "Autonomous orchestration",
    description: "Let agents trigger validated workflows with guardrails and predictable outputs.",
    accent: "from-wallbit-200/80 via-wallbit-500/30 to-transparent",
    Icon: IconAgentSpark,
    featured: true,
  },
] as const;

const flowSteps = [
  {
    id: "define",
    button: "DEFINE STEPS",
    tag: "STEP 1",
    title: "Model the rebalance intent in YAML",
    description:
      "Define the portfolio check as ordered steps so you and your agents share the same plan before execution.",
    example: `version: 1
name: rebalance-check
steps:
  - id: checking
    run: balance.get_checking
  - id: stocks
    run: balance.get_stocks
  - id: fx
    run: rates.get
    with:
      source: USD
      dest: EUR`,
    chips: ["Human-readable plan", "Composable steps", "Reusable file"],
  },
  {
    id: "validate",
    button: "CHECK INPUTS",
    tag: "STEP 2",
    title: "Validate structure and inputs before money moves",
    description:
      "Catch unsupported runs, missing fields, and bad references early so runs stay reliable across environments.",
    example: `$ wallbit workflow validate ./rebalance.yaml
{
  "ok": true,
  "workflow": "rebalance-check",
  "steps": 3
}`,
    chips: ["Contract checks", "Reference checks", "Fast feedback"],
  },
  {
    id: "execute",
    button: "EXECUTE FLOW",
    tag: "STEP 3",
    title: "Run and review deterministic output",
    description:
      "Execute the exact validated flow and consume RunResult JSON for automation, monitoring, or decision layers.",
    example: `$ wallbit workflow run ./rebalance.yaml
{
  "ok": true,
  "started_at": "...",
  "finished_at": "...",
  "steps": [{ "id": "fx", "ok": true }]
}`,
    chips: ["Stable exits", "Structured JSON", "Agent-ready output"],
  },
  {
    id: "automate",
    button: "AGENT AUTOMATION",
    tag: "STEP 4",
    title: "Automate rebalance decisions with your agent",
    description:
      "Once the workflow is validated, your copilot can trigger runs on schedule or based on signals, while preserving the same deterministic contract.",
    example: `# agent loop (pseudo)
signal: "stocks_weight > 60%"
action:
  run: wallbit workflow run ./rebalance.yaml
expect:
  ok: true
  output: RunResult`,
    chips: ["Scheduled or event-driven", "Same validated YAML", "Safe-by-default automation"],
  },
] as const;

function IconBlueprint() {
  return (
    <svg viewBox="0 0 24 24" className="h-6 w-6 sm:h-7 sm:w-7" fill="none" stroke="currentColor" strokeWidth="1.25" aria-hidden>
      <rect x="3.5" y="3.5" width="17" height="17" rx="2.5" />
      <path d="M8 8h8M8 12h8M8 16h5" strokeLinecap="round" />
      <path d="M6.5 6.5v11M17.5 6.5v11" opacity="0.4" />
    </svg>
  );
}

function IconShieldCheck() {
  return (
    <svg viewBox="0 0 24 24" className="h-6 w-6 sm:h-7 sm:w-7" fill="none" stroke="currentColor" strokeWidth="1.25" aria-hidden>
      <path d="M12 3 5 6v5c0 4.2 2.6 7.8 7 10 4.4-2.2 7-5.8 7-10V6l-7-3Z" />
      <path d="m9.2 12.2 1.8 1.8 3.8-3.8" strokeLinecap="round" strokeLinejoin="round" />
    </svg>
  );
}

function IconPlayResults() {
  return (
    <svg viewBox="0 0 24 24" className="h-6 w-6 sm:h-7 sm:w-7" fill="none" stroke="currentColor" strokeWidth="1.25" aria-hidden>
      <rect x="3" y="4" width="18" height="16" rx="2.5" />
      <path d="m9 9 5 3-5 3V9Z" strokeLinejoin="round" />
      <path d="M16.5 9.5h2M16.5 12h2M16.5 14.5h2" strokeLinecap="round" />
    </svg>
  );
}

function IconAgentSpark() {
  return (
    <svg viewBox="0 0 24 24" className="h-6 w-6 sm:h-7 sm:w-7" fill="none" stroke="currentColor" strokeWidth="1.25" aria-hidden>
      <rect x="5" y="7" width="14" height="10" rx="3" />
      <circle cx="10" cy="12" r="1" />
      <circle cx="14" cy="12" r="1" />
      <path d="M9 16h6M12 7V4M8.5 5.5 7 4M15.5 5.5 17 4" strokeLinecap="round" />
      <path d="m19.5 8.5.6 1.3 1.4.2-1 .9.2 1.4-1.2-.7-1.2.7.2-1.4-1-.9 1.4-.2.6-1.3Z" />
    </svg>
  );
}

const focusRing =
  "focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-wallbit-500/90 focus-visible:ring-offset-2 focus-visible:ring-offset-black";

export function ArchitectureHub() {
  const [activeStepId, setActiveStepId] = useState<(typeof flowSteps)[number]["id"]>("define");
  const [progressCycle, setProgressCycle] = useState(0);
  const activeStep = flowSteps.find((step) => step.id === activeStepId) ?? flowSteps[0];
  const handleSelectStep = (stepId: (typeof flowSteps)[number]["id"]) => {
    setActiveStepId(stepId);
    setProgressCycle((prev) => prev + 1);
  };
  const handleAdvanceStep = () => {
    window.setTimeout(() => {
      const currentIndex = pillars.findIndex((p) => p.id === activeStepId);
      const nextIndex = currentIndex === -1 ? 0 : (currentIndex + 1) % pillars.length;
      setActiveStepId(pillars[nextIndex].id);
      setProgressCycle((prev) => prev + 1);
    }, 160);
  };

  return (
    <section
      id="architecture"
      className="border-t border-white/[0.06] bg-black py-20 sm:py-28"
      aria-labelledby="architecture-heading"
    >
      <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        <h2
          id="architecture-heading"
          className="mx-auto max-w-4xl text-center text-[1.65rem] font-medium leading-snug tracking-tight text-pretty text-white sm:text-4xl sm:leading-[1.15]"
        >
          Built on a foundation of fast,{" "}
          <span className="block text-zinc-500 sm:inline">production-grade tooling</span>
        </h2>
        <div className="mx-auto mt-10 grid max-w-6xl gap-6 lg:grid-cols-[1.1fr_1.4fr]">
          <div className="rounded-2xl border border-white/[0.08] bg-gradient-to-b from-zinc-900/80 to-zinc-950/90 p-6 sm:min-h-[36rem] sm:p-8">
            <p className="font-mono text-[10px] font-semibold uppercase tracking-[0.2em] text-wallbit-500/90">Common Portfolio Flow</p>

            <div id={`flow-panel-${activeStep.id}`} role="tabpanel" className="mt-6">
              <p className="font-mono text-[10px] font-semibold uppercase tracking-[0.2em] text-wallbit-500/90">{activeStep.tag}</p>
              <h3 className="mt-2 text-2xl font-semibold tracking-tight text-zinc-100">{activeStep.title}</h3>
              <p className="mt-3 text-sm leading-relaxed text-zinc-400">{activeStep.description}</p>

              <pre className="mt-5 overflow-x-auto rounded-xl border border-zinc-800 bg-black/35 p-4 font-mono text-xs leading-relaxed text-zinc-300 sm:text-[13px]">
                <code>{activeStep.example}</code>
              </pre>

              <div className="mt-4 flex flex-wrap gap-2">
                {activeStep.chips.map((chip) => (
                  <span key={chip} className="rounded-full border border-zinc-700/80 bg-zinc-900/60 px-2.5 py-1 text-[11px] text-zinc-400">
                    {chip}
                  </span>
                ))}
              </div>
            </div>
          </div>

          <div className="grid gap-3 sm:min-h-[36rem] sm:gap-4">
            <p className="font-mono text-[10px] font-semibold uppercase tracking-[0.2em] text-wallbit-500/80">Select a stage</p>
            <ul className="grid gap-3 sm:gap-4" role="tablist" aria-label="Portfolio flow stages">
              {pillars.map((p) => {
                const IconCmp = p.Icon;
                const isActive = p.id === activeStep.id;
                return (
                  <li key={p.id} className="min-w-0">
                    <div className="flow-stage-card relative overflow-hidden rounded-2xl p-px">
                      <div className="flow-stage-borderbox" aria-hidden>
                        {isActive ? (
                          <span
                            key={`${p.id}-${progressCycle}`}
                            className="flow-stage-grow block h-full w-full"
                            onAnimationEnd={handleAdvanceStep}
                          />
                        ) : null}
                      </div>
                      <button
                        type="button"
                        role="tab"
                        aria-selected={isActive}
                        aria-controls={`flow-panel-${p.id}`}
                        onClick={() => handleSelectStep(p.id)}
                        className={`group relative flex w-full touch-manipulation items-start gap-4 overflow-hidden rounded-[calc(1rem-1px)] border bg-[#0a0a0a] p-5 text-left transition-[border-color,background-color] duration-200 sm:gap-5 sm:p-6 ${isActive ? "border-white/[0.08] bg-[#0a0a0a]" : "border-white/[0.08]"
                          } ${focusRing}`}
                      >
                        <div className="relative shrink-0 rounded-md border border-white/[0.06] bg-black/60 p-2.5 text-zinc-300 transition-colors">
                          <IconCmp />
                        </div>
                        <div className="relative min-w-0 flex-1 pt-0.5">
                          <div className="flex items-center gap-2">
                            <h3 className="font-mono text-[11px] font-semibold tracking-wide text-wallbit-300" translate="no">
                              {p.title}
                            </h3>
                          </div>
                          <p className="mt-1 text-base font-medium tracking-tight text-white">{p.subtitle}</p>
                          <p className="mt-1.5 text-sm leading-relaxed text-zinc-400">{p.description}</p>
                        </div>
                      </button>
                    </div>
                  </li>
                );
              })}
            </ul>
          </div>
        </div>
      </div>
    </section>
  );
}
