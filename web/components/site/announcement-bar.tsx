/**
 * Site-wide announcement strip. Slightly taller than a standard top bar
 * because the brand has a real release to announce. Uses three layered
 * effects to draw the eye without screaming:
 *   1. centered radial brand glow,
 *   2. periodic shimmer sweep (CSS-only, see globals.css),
 *   3. luminous brand bottom border.
 */
export function AnnouncementBar() {
  return (
    <div
      role="region"
      aria-label="Announcement"
      className="relative isolate overflow-hidden border-b border-wallbit-700/40 bg-[#070a10]"
    >
      <div
        aria-hidden
        className="pointer-events-none absolute inset-0"
        style={{
          backgroundImage:
            "radial-gradient(ellipse 65% 220% at 50% 50%, rgba(13,153,255,0.24), rgba(13,153,255,0.06) 45%, transparent 78%)",
        }}
      />
      <span
        aria-hidden
        className="announcement-shimmer pointer-events-none absolute inset-y-0 left-0 w-1/2"
      />
      <div
        aria-hidden
        className="pointer-events-none absolute inset-x-0 bottom-0 h-px bg-gradient-to-r from-transparent via-wallbit-300/70 to-transparent"
        style={{ boxShadow: "0 0 10px rgba(13,153,255,0.55)" }}
      />

      <a
        href="https://github.com/jeremyjsx/skills"
        target="_blank"
        rel="noopener noreferrer"
        className="group relative mx-auto flex h-11 max-w-6xl items-center justify-center gap-3 px-4 text-[12.5px] no-underline transition-colors sm:gap-4 sm:px-6"
      >
        <span className="flex min-w-0 items-center gap-2 truncate">
          <span className="font-semibold text-zinc-50">Wallbit Workflow Builder</span>
          <span className="hidden text-zinc-400 sm:inline">
            — build, edit &amp; validate workflow YAML from your AI editor
          </span>
        </span>

        <span className="inline-flex shrink-0 items-center rounded-md border border-zinc-700/80 bg-zinc-900/80 px-2 py-1 font-mono text-[10px] font-semibold uppercase tracking-[0.14em] text-zinc-200 transition-colors duration-200 group-hover:border-wallbit-400/60 group-hover:bg-wallbit-950/70 group-hover:text-wallbit-100 group-focus-visible:border-wallbit-400/60 group-focus-visible:bg-wallbit-950/70 group-focus-visible:text-wallbit-100">
          Install
        </span>
      </a>
    </div>
  );
}
