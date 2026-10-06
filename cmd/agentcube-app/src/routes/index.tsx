import { createFileRoute, Link } from "@tanstack/react-router";
import { ArrowUpRight } from "lucide-react";
import { useEffect, useRef } from "react";

import DitheringBackground from "#/components/(backgrounds)/dithering-background/dithering-background";
import GitHubIcon from "#/components/(icons)/github-icon/github-icon";
import XLogo from "#/components/(icons)/x-logo/x-logo";
import { PageLayout } from "#/components/(layouts)/page-layout/page-layout";
import Spacer from "#/components/(layouts)/spacer/spacer";
import RubiksCube, {
  type RubiksCubeRef,
} from "#/components/(rubiks-cube)/rubiks-cube/rubiks-cube";
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
  const cubeRef = useRef<RubiksCubeRef>(null);

  useEffect(() => {
    const moves = ["R", "U", "R'", "U'", "F", "R", "U", "F'"];
    let moveIndex = 0;

    const interval = window.setInterval(() => {
      if (!cubeRef.current?.getCube()) return;

      void cubeRef.current.rotate(moves[moveIndex % moves.length]);
      moveIndex += 1;
    }, 1200);

    return () => window.clearInterval(interval);
  }, []);

  return (
    <PageLayout
      showGrid={false}
      navbar={false}
      background={<DitheringBackground />}
    >
      <div className="flex flex-1 items-center justify-center py-6">
        <div className="relative flex min-h-104 w-full items-center overflow-hidden bg-transparent sm:min-h-120">
          <div className="relative flex min-h-96 w-full flex-col items-center justify-center gap-5 px-8 py-6 sm:min-h-112 sm:px-14 sm:py-8 md:flex-row md:gap-8">
            <div className="flex min-w-0 flex-1 flex-col justify-center">
              <h1 className="font-heading text-5xl font-extrabold uppercase tracking-tight sm:text-7xl lg:text-8xl">
                Agent Cube
              </h1>
              <Spacer size="xs" />
              <p className="max-w-lg font-mono text-sm leading-relaxed text-neutral-400 sm:text-base">
                Benchmarking LLM reasoning and problem-solving through the
                Rubik's Cube challenge
              </p>
              <Spacer size="xs" />
              <div className="flex items-center gap-2 self-start">
                <Link
                  to="/agents"
                  className={buttonVariants({
                    size: "sm",
                  })}
                >
                  View Agents
                  <ArrowUpRight className="ml-2" />
                </Link>
                <a
                  href={import.meta.env.VITE_GITHUB_URL}
                  target="_blank"
                  rel="noopener noreferrer"
                  className={buttonVariants({ size: "icon-sm" })}
                  aria-label="GitHub"
                >
                  <GitHubIcon size={20} />
                </a>
                <a
                  href={import.meta.env.VITE_X_URL}
                  target="_blank"
                  rel="noopener noreferrer"
                  className={buttonVariants({ size: "icon-sm" })}
                  aria-label="X"
                >
                  <XLogo size={18} />
                </a>
              </div>
            </div>
            <div className="hidden shrink-0 items-center justify-center border border-neutral-700 bg-black p-2 md:flex">
              <RubiksCube
                ref={cubeRef}
                algorithm="R U R' U' F R U R' U' F'"
                size={250}
                showRotationsCollapsible={false}
                showRotationButtons={false}
                showResetAlgoButton={false}
                showResetCameraButton={false}
                showBorder={false}
              />
            </div>
          </div>
        </div>
      </div>
    </PageLayout>
  );
}
