export function ProductSection() {
  return (
    <section id="product" className="border-b border-zinc-800/60 py-14 sm:py-20">
      <div className="mx-auto max-w-6xl px-4 sm:px-6">
        <p className="text-center font-mono text-[10px] font-semibold uppercase tracking-[0.2em] text-wallbit-500/80 sm:text-left">
          Product
        </p>
        <h2 className="mt-2 text-center text-2xl font-semibold tracking-tight text-zinc-100 sm:text-left sm:text-3xl">
          One CLI, two ways to win
        </h2>
        <p className="mx-auto mt-3 max-w-2xl text-center text-zinc-500 sm:mx-0 sm:text-left">
          Move fast with direct commands, then lock in repeatability with YAML workflows you and your agents can run with
          the same outcome every time.
        </p>
        <div className="mt-10 grid gap-5 lg:grid-cols-2">
          <div className="relative overflow-hidden rounded-2xl border border-zinc-800 bg-gradient-to-b from-zinc-900/80 to-zinc-950/80 p-6 sm:p-8">
            <div className="absolute right-4 top-4 font-mono text-[10px] text-zinc-600">01</div>
            <p className="font-mono text-xs font-medium uppercase tracking-wider text-wallbit-400/90">Direct commands</p>
            <h3 className="mt-2 text-lg font-semibold text-zinc-100">Ship outcomes in seconds</h3>
            <p className="mt-3 text-sm leading-relaxed text-zinc-500">
              Query balances, rates, cards, trades, and more with explicit flags and clean JSON. Perfect for quick
              decisions, scripts, and agent prompts.
            </p>
            <pre className="mt-5 overflow-x-auto rounded-lg border border-zinc-800/80 bg-black/40 p-3 font-mono text-[12px] leading-relaxed text-zinc-400">
              <code>
                <span className="text-zinc-600">$ </span>
                <span className="text-zinc-200">wallbit balance checking</span>
                {"\n"}
                <span className="text-zinc-600">$ </span>
                <span className="text-zinc-200">wallbit rates get --source USD --dest EUR</span>
              </code>
            </pre>
          </div>
          <div className="relative overflow-hidden rounded-2xl border border-wallbit-800/30 bg-gradient-to-b from-wallbit-950/40 to-zinc-950/90 p-6 sm:p-8">
            <div className="absolute right-4 top-4 font-mono text-[10px] text-wallbit-700">02</div>
            <p className="font-mono text-xs font-medium uppercase tracking-wider text-wallbit-400/90">YAML workflows</p>
            <h3 className="mt-2 text-lg font-semibold text-zinc-100">Scale execution like infrastructure</h3>
            <p className="mt-3 text-sm leading-relaxed text-zinc-500">
              Define multi-step financial flows in YAML, validate before run, and execute the same plan across environments
              with zero drift.
            </p>
            <pre className="mt-5 overflow-x-auto rounded-lg border border-wallbit-900/40 bg-black/35 p-3 font-mono text-[12px] leading-relaxed text-zinc-400">
              <code>
                <span className="text-zinc-600">$ </span>
                <span className="text-zinc-200">wallbit workflow validate ./payroll.yaml</span>
                {"\n"}
                <span className="text-zinc-600">$ </span>
                <span className="text-zinc-200">wallbit workflow run ./payroll.yaml</span>
              </code>
            </pre>
          </div>
        </div>
      </div>
    </section>
  );
}
