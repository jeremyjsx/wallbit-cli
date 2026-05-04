import Link from "next/link";
import { WallbitMark } from "../site/wallbit-mark";

function Logo() {
  return (
    <Link
      href="/"
      aria-label="Wallbit CLI — home"
      className="inline-flex items-center gap-2.5 rounded-md no-underline focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-wallbit-500/90 focus-visible:ring-offset-2 focus-visible:ring-offset-black"
    >
      <WallbitMark className="h-[22px] w-auto text-wallbit-500" />
        <span className="text-[15px] font-semibold tracking-tight text-zinc-50">wallbit-cli</span>
    </Link>
  );
}

export function HomeHeader() {
  const navLinkClass =
    "transition-colors hover:text-wallbit-400 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-wallbit-500/90 focus-visible:ring-offset-2 focus-visible:ring-offset-black rounded-sm";

  return (
    <header className="sticky top-0 z-50 border-b border-zinc-800/80 bg-[var(--background)]/80 backdrop-blur-md">
      <div className="mx-auto flex h-14 max-w-6xl items-center justify-between gap-3 px-4 sm:px-6">
        <Logo />
        <nav className="hidden items-center gap-8 text-sm text-zinc-400 sm:flex" aria-label="Main">
          <a href="#product" className={navLinkClass}>
            Product
          </a>
          <a href="#features" className={navLinkClass}>
            Features
          </a>
          <a href="#architecture" className={navLinkClass}>
            Workflow
          </a>
        </nav>
        <div className="hidden items-center gap-2 sm:flex">
          <Link
            href="/docs"
            className="rounded-full border border-zinc-700 bg-zinc-900/50 px-3.5 py-1.5 text-sm font-medium text-zinc-200 no-underline transition-colors hover:border-zinc-500 hover:text-zinc-100 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-wallbit-400/90 focus-visible:ring-offset-2 focus-visible:ring-offset-black"
          >
            Docs
          </Link>
          <Link
            href="/docs#workflow-quickstart"
            className="shrink-0 rounded-full border border-zinc-600 bg-zinc-100 px-4 py-1.5 text-sm font-medium text-zinc-950 no-underline transition-colors hover:bg-white focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-wallbit-400/90 focus-visible:ring-offset-2 focus-visible:ring-offset-black"
          >
            Quickstart
          </Link>
        </div>
      </div>
      <nav className="border-t border-zinc-900 px-4 py-2 sm:hidden" aria-label="Mobile main">
        <div className="flex items-center gap-4 overflow-x-auto whitespace-nowrap text-sm text-zinc-400 docs-scrollbar-hidden">
          <a href="#product" className={navLinkClass}>
            Product
          </a>
          <a href="#features" className={navLinkClass}>
            Features
          </a>
          <a href="#architecture" className={navLinkClass}>
            Workflow
          </a>
          <Link href="/docs" className={navLinkClass}>
            Docs
          </Link>
          <Link href="/docs#workflow-quickstart" className={navLinkClass}>Quickstart</Link>
        </div>
      </nav>
    </header>
  );
}
