import Link from "next/link";

type HomeCtaSectionProps = {
  githubUrl: string;
};

export function HomeCtaSection({ githubUrl }: HomeCtaSectionProps) {
  return (
    <section className="border-t border-zinc-800/60 py-24 sm:py-28">
      <div className="mx-auto max-w-6xl px-4 text-center sm:px-6">
        <p className="font-mono text-[12px] text-zinc-500">
          <span className="text-wallbit-500/85">[</span>{" "}
          <span className="text-zinc-300">ready to ship?</span>{" "}
          <span className="text-wallbit-500/85">]</span>
        </p>
        <h2 className="mt-3 text-3xl font-semibold tracking-tight text-zinc-100 sm:text-4xl">
          Ready to run your first real workflow?
        </h2>
        <p className="mx-auto mt-4 max-w-3xl text-base leading-relaxed text-zinc-400 sm:text-lg">
          Start with the quickstart, adapt the YAML to your use case, and execute with deterministic output from the CLI.
        </p>
        <div className="mt-10 flex flex-col items-center justify-center gap-3 sm:flex-row">
          <Link
            href="/docs#workflow-quickstart"
            className="inline-flex h-12 min-w-[220px] items-center justify-center rounded-full bg-wallbit-500 px-7 text-base font-semibold text-zinc-950 no-underline transition-colors hover:bg-wallbit-400 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-wallbit-400/90 focus-visible:ring-offset-2 focus-visible:ring-offset-black"
          >
            Open Quickstart
          </Link>
          <a
            href={githubUrl}
            target="_blank"
            rel="noreferrer noopener"
            className="inline-flex h-12 min-w-[220px] items-center justify-center rounded-full border border-zinc-700 bg-zinc-900/40 px-7 text-base font-semibold text-zinc-200 no-underline transition-colors hover:border-zinc-500 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-wallbit-400/90 focus-visible:ring-offset-2 focus-visible:ring-offset-black"
          >
            View on GitHub
          </a>
        </div>
      </div>
    </section>
  );
}
