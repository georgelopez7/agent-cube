import { createFileRoute, Link } from "@tanstack/react-router";
import { ArrowUpRight } from "lucide-react";

import { PageLayout } from "#/components/(layouts)/page-layout/page-layout";
import Spacer from "#/components/(layouts)/spacer/spacer";
import CubePulse from "#/components/(rubiks-cube)/cube-pulse/cube-pulse";
import { buttonVariants } from "#/components/ui/button";
import { getOGImage, getSiteUrl } from "#/utils/seo";

const SEO = {
  url: getSiteUrl(),
  title: "Agent Cube | Benchmark LLMs on the Rubik's Cube",
  description:
    "Agent Cube is an open benchmark that tests large language models by challenging them to solve Rubik's Cube puzzles. Compare reasoning, planning, and problem-solving across top LLMs.",
  image: getOGImage(),
};

export const Route = createFileRoute("/")({
  component: Home,
  head: () => ({
    meta: [
      {
        title: SEO.title,
      },
      {
        name: "description",
        content: SEO.description,
      },
      {
        property: "og:title",
        content: SEO.title,
      },
      {
        property: "og:description",
        content: SEO.description,
      },
      {
        property: "og:type",
        content: "website",
      },
      {
        property: "og:image",
        content: SEO.image,
      },
      {
        name: "twitter:card",
        content: "summary_large_image",
      },
      {
        name: "twitter:title",
        content: SEO.title,
      },
      {
        name: "twitter:description",
        content: SEO.description,
      },
      {
        name: "twitter:image",
        content: SEO.image,
      },
    ],
  }),
});

function Home() {
  return (
    <PageLayout>
      <div className="flex flex-1 items-center justify-center -translate-y-12">
        <div className="flex items-center gap-24">
          <div>
            <h1 className="text-6xl md:text-7xl lg:text-8xl font-bold font-heading tracking-tight text-foreground">
              Agent Cube
            </h1>
            <Spacer size="xs" />
            <p className="text-lg md:text-xl text-muted-foreground max-w-lg leading-relaxed">
              Benchmarking LLM reasoning and problem-solving through the Rubik's
              Cube challenge
            </p>
            <Spacer size="xs" />
            <Link to="/agents" className={buttonVariants()}>
              View Agents
              <ArrowUpRight />
            </Link>
          </div>
          <div className="hidden md:flex items-center justify-center">
            <CubePulse randomize />
          </div>
        </div>
      </div>
    </PageLayout>
  );
}
