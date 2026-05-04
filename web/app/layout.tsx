import type { Metadata } from "next";
import { Geist, Geist_Mono } from "next/font/google";
import { AnnouncementBar } from "../components/site/announcement-bar";
import { GlobalFooter } from "../components/site/global-footer";
import "./globals.css";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: "Wallbit CLI",
  description: "YAML-first Wallbit CLI. Plan workflows, validate, and run from the terminal.",
  openGraph: {
    title: "Wallbit CLI",
    description: "Plan in YAML. Run from the terminal.",
  },
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" className={`${geistSans.variable} ${geistMono.variable} antialiased`}>
      <body className="min-h-dvh flex flex-col bg-[var(--background)] text-[var(--foreground)]">
        <a
          href="#main-content"
          className="sr-only absolute left-3 top-3 z-[100] rounded-md bg-zinc-100 px-3 py-2 text-sm font-medium text-zinc-950 no-underline focus:not-sr-only focus:outline-none focus:ring-2 focus:ring-wallbit-500"
        >
          Skip to content
        </a>
        <AnnouncementBar />
        <div id="main-content" className="contents">
          {children}
        </div>
        <GlobalFooter />
      </body>
    </html>
  );
}
