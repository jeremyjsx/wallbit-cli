import Link from "next/link";

const GITHUB_REPO = "jeremyjsx/wallbit-cli";
const GITHUB_URL = `https://github.com/${GITHUB_REPO}`;

export async function GlobalFooter() {
  return (
    <footer className="border-t border-zinc-800/80 bg-black/30">
      <div className="mx-auto flex max-w-6xl flex-col gap-3 px-4 py-5 sm:px-6 md:flex-row md:items-center md:justify-between">
        <p className="text-xs leading-relaxed text-zinc-500">
          &copy; {new Date().getFullYear()} <span className="text-zinc-300">wallbit-cli</span> by Jeremy. Independent tool
          built on top of the Wallbit public API.
        </p>
        <div className="flex flex-wrap items-center gap-x-4 gap-y-1 text-xs">
          <Link href="/docs#workflow-quickstart" className="text-wallbit-300 no-underline transition-colors hover:text-wallbit-200 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-wallbit-500/90 focus-visible:ring-offset-2 focus-visible:ring-offset-black rounded-sm">
            Quickstart
          </Link>
          <Link href="/docs" className="text-zinc-500 no-underline transition-colors hover:text-wallbit-400 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-wallbit-500/90 focus-visible:ring-offset-2 focus-visible:ring-offset-black rounded-sm">
            Docs
          </Link>
          <a
            href={GITHUB_URL}
            target="_blank"
            rel="noreferrer noopener"
            className="text-zinc-500 no-underline transition-colors hover:text-wallbit-400 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-wallbit-500/90 focus-visible:ring-offset-2 focus-visible:ring-offset-black rounded-sm"
          >
            GitHub
          </a>
        </div>
      </div>
    </footer>
  );
}
