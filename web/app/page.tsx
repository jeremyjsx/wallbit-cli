import { ArchitectureHub } from "../components/architecture/architecture-hub";
import { FeaturesSection } from "../components/home/features-section";
import { GetStartedSection } from "../components/home/get-started-section";
import { HeroSection } from "../components/home/hero-section";
import { HomeCtaSection } from "../components/home/home-cta-section";
import { HomeHeader } from "../components/home/home-header";
import { ProductSection } from "../components/home/product-section";

const GITHUB = "https://github.com/jeremyjsx/wallbit-cli";

export default function Home() {
  return (
    <div className="flex min-h-dvh flex-col">
      <HomeHeader />

      <main className="flex-1">
        <HeroSection githubUrl={GITHUB} />
        <ProductSection />
        <GetStartedSection />
        <FeaturesSection />

        <ArchitectureHub />
        <HomeCtaSection githubUrl={GITHUB} />
      </main>
    </div>
  );
}
