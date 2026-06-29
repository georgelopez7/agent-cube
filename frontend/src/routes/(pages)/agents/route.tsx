import { useQuery } from "@tanstack/react-query";
import { createFileRoute } from "@tanstack/react-router";

import RubiksCubeCard from "#/components/(cards)/rubiks-cube-card/rubiks-cube-card";
import { PageLayout } from "#/components/(layouts)/page-layout/page-layout";
import CreateAgentModal from "#/components/(modals)/create-agent-modal/create-agent-modal";
import { Skeleton } from "#/components/ui/skeleton";
import { RUBIKS_CUBES_QUERIES } from "#/services/rubiks-cubes.queries";

const CUBES_LIMIT = 10;

export const Route = createFileRoute("/(pages)/agents")({
  component: RouteComponent,
});

function RouteComponent() {
  const { data: cubes, isLoading } = useQuery(
    RUBIKS_CUBES_QUERIES.getAll(CUBES_LIMIT),
  );

  return (
    <PageLayout>
      <div className="w-[80vw] flex flex-col gap-8">
        <div className="flex items-center justify-between gap-4">
          <h1 className="text-4xl font-semibold">Agents</h1>
          <CreateAgentModal />
        </div>
        {isLoading ? (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
            {Array.from({ length: CUBES_LIMIT }).map((_, index) => (
              <Skeleton key={index} className="h-96 w-full rounded-lg" /> // biome-ignore
            ))}
          </div>
        ) : (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
            {cubes?.map((cube) => (
              <RubiksCubeCard key={cube.id} cube={cube} />
            ))}
          </div>
        )}
      </div>
    </PageLayout>
  );
}
