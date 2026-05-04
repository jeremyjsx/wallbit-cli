import Link from "next/link";

type HeroSectionProps = {
  githubUrl: string;
};

export function HeroSection({ githubUrl }: HeroSectionProps) {
  return (
    <section className="relative overflow-hidden border-b border-zinc-800/60">
      <div className="hero-grid pointer-events-none absolute inset-0" aria-hidden />
      <div className="relative mx-auto max-w-6xl px-4 pb-16 pt-14 sm:px-6 sm:pb-24 sm:pt-20">
        <div className="mb-5 flex flex-wrap items-center justify-center gap-2 sm:justify-start">
          <span className="rounded-full border border-zinc-700/80 bg-zinc-900/60 px-3 py-1 font-mono text-[11px] text-zinc-400">
            open source
          </span>
          <a
            href="https://github.com/jeremyjsx/wallbit-go"
            target="_blank"
            rel="noreferrer noopener"
            className="rounded-full border border-wallbit-800/40 bg-wallbit-950/40 px-3 py-1 font-mono text-[11px] text-wallbit-400/90 no-underline transition-colors hover:border-wallbit-600/50 hover:text-wallbit-300 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-wallbit-500/90 focus-visible:ring-offset-2 focus-visible:ring-offset-black"
          >
            powered by wallbit-go
          </a>
        </div>
        <h1 className="mx-auto max-w-4xl text-center text-[clamp(2.25rem,7vw,3.75rem)] font-semibold leading-[1.05] tracking-[-0.035em] text-zinc-50 sm:mx-0 sm:text-left">
          <span className="block text-zinc-300">Automate financial workflows</span>
          <span className="block bg-gradient-to-br from-wallbit-200 via-wallbit-400 to-wallbit-600 bg-clip-text text-transparent">
            from your terminal.
          </span>
        </h1>
        <p className="mx-auto mt-6 max-w-2xl text-center text-lg leading-relaxed text-zinc-400 sm:mx-0 sm:text-left">
          <strong className="font-semibold text-zinc-200">wallbit-cli</strong> turns AI plans into reliable execution:
          define flows in YAML, run them with confidence, and keep every step observable.
        </p>
        <p className="mx-auto mt-4 max-w-2xl text-center text-sm leading-relaxed text-zinc-600 sm:mx-0 sm:text-left">
          Built for builders and agents who want speed without chaos: deterministic outputs, validation before execution,
          and composable workflows that scale from one-off runs to production automations.
        </p>
        <div className="mt-10 flex flex-col items-center justify-center gap-3 sm:flex-row sm:justify-start">
          <Link
            href="/docs#workflow-quickstart"
            className="inline-flex h-11 w-full items-center justify-center rounded-full bg-wallbit-500 px-6 text-sm font-medium text-zinc-950 no-underline transition-colors hover:bg-wallbit-400 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-wallbit-400/90 focus-visible:ring-offset-2 focus-visible:ring-offset-black sm:w-auto"
          >
            Open Quickstart
          </Link>
          <a
            href={githubUrl}
            target="_blank"
            rel="noreferrer noopener"
            className="inline-flex h-11 w-full items-center justify-center rounded-full border border-zinc-700 bg-transparent px-6 text-sm font-medium text-zinc-200 no-underline transition-colors hover:border-wallbit-600/50 hover:bg-zinc-900/50 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-wallbit-400/90 focus-visible:ring-offset-2 focus-visible:ring-offset-black sm:w-auto"
          >
            View on GitHub
          </a>
        </div>
      </div>
    </section>
  );
}
