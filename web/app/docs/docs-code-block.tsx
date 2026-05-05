"use client";

import { useCallback, useState } from "react";

export function DocsCodeBlock({ children }: { children: string }) {
  const [copied, setCopied] = useState(false);

  const copy = useCallback(async () => {
    try {
      await navigator.clipboard.writeText(children);
      setCopied(true);
      window.setTimeout(() => setCopied(false), 2000);
    } catch {
      setCopied(false);
    }
  }, [children]);

  return (
    <div className="group relative my-4">
      <button
        type="button"
        onClick={copy}
        aria-label={copied ? "Copied" : "Copy code"}
        className="absolute right-2 top-2 z-10 inline-flex h-8 items-center gap-1.5 rounded-lg border border-zinc-700/80 bg-zinc-900/90 px-2.5 text-xs font-medium text-zinc-400 shadow-sm backdrop-blur-sm transition-colors hover:border-wallbit-600/40 hover:bg-zinc-800/90 hover:text-wallbit-300 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-wallbit-500/70 focus-visible:ring-offset-2 focus-visible:ring-offset-[var(--panel)]"
      >
        {copied ? (
          <>
            <svg className="h-3.5 w-3.5 text-wallbit-400" viewBox="0 0 24 24" fill="none" aria-hidden>
              <path
                d="M20 6L9 17l-5-5"
                stroke="currentColor"
                strokeWidth="2"
                strokeLinecap="round"
                strokeLinejoin="round"
              />
            </svg>
            <span className="text-wallbit-300/95">Copied</span>
          </>
        ) : (
          <>
            <svg className="h-3.5 w-3.5 opacity-80" viewBox="0 0 24 24" fill="none" aria-hidden>
              <rect
                x="9"
                y="9"
                width="13"
                height="13"
                rx="2"
                stroke="currentColor"
                strokeWidth="1.5"
              />
              <path
                d="M5 15H4a2 2 0 01-2-2V4a2 2 0 012-2h9a2 2 0 012 2v1"
                stroke="currentColor"
                strokeWidth="1.5"
                strokeLinecap="round"
              />
            </svg>
            <span>Copy</span>
          </>
        )}
      </button>
      <pre className="overflow-x-auto rounded-xl border border-zinc-800 border-l-2 border-l-wallbit-500/55 bg-[var(--panel)] p-4 text-[13px] leading-relaxed text-zinc-300">
        <code>{children}</code>
      </pre>
    </div>
  );
}
