"use client";

import { useCallback, useEffect, useLayoutEffect, useRef, useState } from "react";
import { DOCS_NAV_SECTIONS, docsNavIdsInOrder } from "./nav-config";

/** Viewport line (px from top) used by scrollspy. Slightly below sticky header for stable section matching. */
const ACTIVE_LINE_PX = 120;

function computeActiveId(): string {
  const ids = docsNavIdsInOrder();
  const entries = ids
    .map((id) => {
      const el = document.getElementById(id);
      return el ? { id, top: el.getBoundingClientRect().top } : null;
    })
    .filter((e): e is { id: string; top: number } => Boolean(e));

  if (entries.length === 0) return ids[0] || "";

  // If the URL hash is already near the active line, prefer it to avoid "previous item still highlighted".
  const hashId = window.location.hash.slice(1);
  if (hashId) {
    const hashEntry = entries.find((e) => e.id === hashId);
    if (hashEntry && hashEntry.top <= ACTIVE_LINE_PX + 220 && hashEntry.top >= ACTIVE_LINE_PX - 80) {
      return hashEntry.id;
    }
  }

  // Primary rule: last heading that has crossed the active line.
  const passed = entries.filter((e) => e.top <= ACTIVE_LINE_PX + 4);
  if (passed.length > 0) {
    return passed[passed.length - 1]!.id;
  }

  // Fallback near top: first upcoming heading.
  return entries[0]!.id;
}

export function DocsSidebar() {
  const [activeId, setActiveId] = useState("");
  const linkRefs = useRef<Map<string, HTMLAnchorElement>>(new Map());

  const sync = useCallback(() => {
    setActiveId(computeActiveId());
  }, []);

  useLayoutEffect(() => {
    if (!activeId) return;
    const el = linkRefs.current.get(activeId);
    el?.scrollIntoView({ block: "nearest", behavior: "instant" });
  }, [activeId]);

  useEffect(() => {
    const t0 = window.setTimeout(() => sync(), 0);
    const onScroll = () => sync();
    window.addEventListener("scroll", onScroll, { passive: true });
    window.addEventListener("resize", onScroll);
    window.addEventListener("hashchange", sync);
    const t = window.setTimeout(sync, 200);
    return () => {
      window.clearTimeout(t0);
      window.removeEventListener("scroll", onScroll);
      window.removeEventListener("resize", onScroll);
      window.removeEventListener("hashchange", sync);
      window.clearTimeout(t);
    };
  }, [sync]);

  return (
    <aside className="shrink-0 lg:w-52">
      <nav
        className="docs-scrollbar-hidden lg:sticky lg:top-20 lg:max-h-[calc(100dvh-6rem)] lg:overflow-y-auto lg:pr-1"
        style={{ msOverflowStyle: "none", scrollbarWidth: "none" }}
        aria-label="Documentation"
      >
        <p className="mb-3 font-mono text-[10px] font-medium uppercase tracking-widest text-wallbit-600">Menu</p>
        <div className="flex flex-wrap gap-2 border-b border-zinc-800/60 pb-6 lg:flex-col lg:gap-0 lg:border-0 lg:pb-0">
          {DOCS_NAV_SECTIONS.map((section) => (
            <div key={section.title} className="min-w-[140px] flex-1 lg:mb-6 lg:flex-none">
              <p className="mb-2 font-mono text-[10px] font-semibold uppercase tracking-widest text-wallbit-500/90">
                {section.title}
              </p>
              <ul className="space-y-0.5">
                {section.items.map((item) => {
                  const id = item.href.slice(1);
                  const isActive = activeId === id;
                  return (
                    <li key={item.href}>
                      <a
                        ref={(node) => {
                          if (node) linkRefs.current.set(id, node);
                          else linkRefs.current.delete(id);
                        }}
                        href={item.href}
                        className={[
                          "block rounded-md border-l-2 py-1.5 pl-2 pr-2 text-sm no-underline transition-colors",
                          isActive
                            ? "border-wallbit-500 bg-wallbit-950/70 text-wallbit-200"
                            : "border-transparent text-zinc-400 hover:bg-wallbit-950/50 hover:text-wallbit-300",
                        ].join(" ")}
                        aria-current={isActive ? "location" : undefined}
                        onClick={() => {
                          window.requestAnimationFrame(() => sync());
                        }}
                      >
                        {item.label}
                      </a>
                    </li>
                  );
                })}
              </ul>
            </div>
          ))}
        </div>
      </nav>
    </aside>
  );
}
