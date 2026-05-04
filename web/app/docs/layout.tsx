import type { Metadata } from "next";

export const metadata: Metadata = {
  title: "Documentation | Wallbit CLI",
  description:
    "Install Wallbit CLI, configure API keys, use every command, and author YAML workflows with validate and run.",
};

export default function DocsLayout({ children }: { children: React.ReactNode }) {
  return children;
}
