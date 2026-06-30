import { useQuery } from "@tanstack/react-query";
import { createFileRoute } from "@tanstack/react-router";

import { PageLayout } from "#/components/(layouts)/page-layout/page-layout";
import Spacer from "#/components/(layouts)/spacer/spacer";
import CreateAgentModal, {
  type CreateAgentFormData,
} from "#/components/(modals)/create-agent-modal/create-agent-modal";
import RubiksCubeSection from "#/components/(sections)/rubiks-cube-section/rubiks-cube-section";
import Websocket from "#/components/websocket/websocket";
import { AI_QUERIES } from "#/services/ai.queries";
import {
  RUBIKS_CUBES_QUERIES,
  useCreateRubiksCube,
} from "#/services/rubiks-cubes.queries";

const CUBES_LIMIT = 10;

export const Route = createFileRoute("/(pages)/agents")({
  component: RouteComponent,
});

function RouteComponent() {
  const { data: cubes, isLoading } = useQuery(
    RUBIKS_CUBES_QUERIES.getAll(CUBES_LIMIT),
  );

  const { data: llms } = useQuery(AI_QUERIES.getModels());
  const { mutateAsync: createRubiksCube } = useCreateRubiksCube();

  const handleCreateAgent = async (data: CreateAgentFormData) => {
    await createRubiksCube({
      llm: data.model,
      scramble: data.scramble,
    });
  };

  return (
    <PageLayout>
      <Websocket debug />
      <div className="flex flex-col">
        <CreateAgentModal llms={llms ?? []} onSubmit={handleCreateAgent} />
        <Spacer size="xs" />
        <RubiksCubeSection
          cubes={cubes}
          isLoading={isLoading}
          limit={CUBES_LIMIT}
        />
      </div>
    </PageLayout>
  );
}
