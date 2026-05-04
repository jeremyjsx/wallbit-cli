export type DocsNavItem = { href: string; label: string };
export type DocsNavSection = { title: string; items: DocsNavItem[] };

/** Single source of truth for docs anchors — must match `id` on headings in `page.tsx`. */
export const DOCS_NAV_SECTIONS: DocsNavSection[] = [
  {
    title: "Getting started",
    items: [
      { href: "#prerequisites", label: "Prerequisites" },
      { href: "#installation", label: "Installation" },
      { href: "#authentication", label: "Authentication" },
      { href: "#global-flags", label: "Global flags" },
    ],
  },
  {
    title: "Command reference",
    items: [
      { href: "#cmd-auth", label: "auth" },
      { href: "#cmd-overview", label: "Command overview" },
      { href: "#cmd-trades", label: "trades (detailed)" },
      { href: "#environment", label: "Environment" },
    ],
  },
  {
    title: "Workflows",
    items: [
      { href: "#workflow-quickstart", label: "First workflow in 5 min" },
      { href: "#workflow-language", label: "Workflow language" },
      { href: "#workflow-patterns", label: "Patterns cookbook" },
      { href: "#workflow-ai-prompts", label: "AI prompt templates" },
      { href: "#workflow-spec", label: "YAML specification" },
      { href: "#workflow-validate-run", label: "Validate, Run & Output" },
      { href: "#cmd-workflow", label: "workflow" },
      { href: "#output", label: "Output format" },
      { href: "#workflow-step-refs", label: "Step output references" },
      { href: "#workflow-catalog", label: "Run catalog" },
    ],
  },
];

export function docsNavIdsInOrder(): string[] {
  return DOCS_NAV_SECTIONS.flatMap((s) => s.items.map((i) => i.href.slice(1)));
}
