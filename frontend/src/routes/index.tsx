import { createFileRoute, Link } from "@tanstack/react-router";
import { ArrowUpRight } from "lucide-react";

import { PageLayout } from "#/components/(layouts)/page-layout/page-layout";
import Spacer from "#/components/(layouts)/spacer/spacer";
import CubePulse from "#/components/(rubiks-cube)/cube-pulse/cube-pulse";
import { buttonVariants } from "#/components/ui/button";

export const Route = createFileRoute("/")({ component: Home });

function Home() {
  return (
    <PageLayout>
      <div className="flex flex-1 items-center justify-center">
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
