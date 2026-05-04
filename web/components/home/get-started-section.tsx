import type { CSSProperties } from "react";

type LineStyle = CSSProperties & { "--tw-len"?: number; "--tw-delay"?: string };

const lineStyle = (len: number, delay: number): LineStyle => ({
  ["--tw-len"]: len,
  ["--tw-delay"]: `${delay}s`,
});

export function GetStartedSection() {
  return (
    <section id="get-started" className="border-b border-zinc-800/60 py-14 sm:py-20">
      <div className="mx-auto max-w-6xl px-4 sm:px-6">
        <p className="font-mono text-[10px] font-semibold uppercase tracking-[0.2em] text-wallbit-500/80">Quickstart</p>
        <h2 className="mt-2 text-2xl font-semibold tracking-tight text-zinc-100 sm:text-3xl">Install, Validate, Execute</h2>
        <p className="mt-3 max-w-2xl text-zinc-500">
          Run your first end-to-end workflow in minutes with the exact commands below.
        </p>
        <div className="mb-4 mt-6 flex flex-wrap items-center gap-2 text-xs">
          <span className="rounded-full border border-zinc-700/80 bg-zinc-900/60 px-2.5 py-1 font-mono text-zinc-400">1. Install</span>
          <span className="rounded-full border border-zinc-700/80 bg-zinc-900/60 px-2.5 py-1 font-mono text-zinc-400">2. Auth</span>
          <span className="rounded-full border border-zinc-700/80 bg-zinc-900/60 px-2.5 py-1 font-mono text-zinc-400">3. Validate</span>
          <span className="rounded-full border border-zinc-700/80 bg-zinc-900/60 px-2.5 py-1 font-mono text-zinc-400">4. Run</span>
        </div>
        <div className="overflow-hidden rounded-2xl border border-zinc-800 bg-[var(--panel)] shadow-[0_0_0_1px_rgba(255,255,255,0.04)]">
          <div className="flex items-center gap-2 border-b border-zinc-800/80 px-4 py-2.5">
            <span className="h-2.5 w-2.5 rounded-full bg-zinc-600" />
            <span className="h-2.5 w-2.5 rounded-full bg-zinc-600" />
            <span className="h-2.5 w-2.5 rounded-full bg-zinc-600" />
            <span className="ml-2 font-mono text-[11px] text-zinc-500">~/wallbit · zsh</span>
          </div>
          <pre
            className="overflow-x-auto p-4 font-mono text-[13px] leading-relaxed text-zinc-300 sm:p-6 sm:text-sm"
            aria-label="Terminal walkthrough: install, authenticate, validate, run."
          >
            <code>
              <span className="tw-line" style={lineStyle(63, 0)}>
                <span className="text-zinc-600">$ </span>
                <span className="text-wallbit-300">go install github.com/jeremyjsx/wallbit-cli/cmd/wallbit@latest</span>
              </span>
              <span className="block h-[0.9em]" aria-hidden />
              <span className="tw-line" style={lineStyle(20, 2.1)}>
                <span className="text-zinc-600">$ </span>
                <span className="text-zinc-100">wallbit auth login</span>
              </span>
              <span className="tw-line" style={lineStyle(33, 2.85)}>
                <span className="text-zinc-500"># paste key once → stored locally</span>
              </span>
              <span className="block h-[0.9em]" aria-hidden />
              <span className="tw-line" style={lineStyle(43, 3.95)}>
                <span className="text-zinc-600">$ </span>
                <span className="text-zinc-100">wallbit workflow validate ./workflow.yaml</span>
              </span>
              <span className="tw-line" style={lineStyle(33, 5.3)}>
                <span className="text-wallbit-400">&gt; OK — steps and inputs look good</span>
              </span>
              <span className="block h-[0.9em]" aria-hidden />
              <span className="tw-line" style={lineStyle(38, 6.4)}>
                <span className="text-zinc-600">$ </span>
                <span className="text-zinc-100">wallbit workflow run ./workflow.yaml</span>
              </span>
              <span className="tw-line" style={lineStyle(27, 7.6)}>
                <span className="text-zinc-500">&gt; JSON run report on stdout</span>
                <span className="tw-caret text-wallbit-300" style={{ ["--tw-delay"]: "8.4s" } as LineStyle} aria-hidden />
              </span>
            </code>
          </pre>
        </div>
      </div>
    </section>
  );
}
