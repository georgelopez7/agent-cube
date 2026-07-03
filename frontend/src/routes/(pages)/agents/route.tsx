import { useQuery, useQueryClient } from "@tanstack/react-query";
import { createFileRoute } from "@tanstack/react-router";
import { PageLayout } from "#/components/(layouts)/page-layout/page-layout";
import Spacer from "#/components/(layouts)/spacer/spacer";
import CreateAgentModal, {
  type CreateAgentFormData,
} from "#/components/(modals)/create-agent-modal/create-agent-modal";
import RubiksCubeSection from "#/components/(sections)/rubiks-cube-section/rubiks-cube-section";
import Websocket from "#/components/websocket/websocket";
import { GetAIModels } from "#/services/ai.queries";
import { CreateRubiksCubeFn } from "#/services/rubiks-cubes.actions";
import {
  GetRubiksCubes,
  rubiksCubeKeys,
} from "#/services/rubiks-cubes.queries";

export const Route = createFileRoute("/(pages)/agents")({
  component: RouteComponent,
});

function RouteComponent() {
  const CubesLimit = 10;

  const _query = useQueryClient();

  const { data: cubes, isLoading } = useQuery(GetRubiksCubes(CubesLimit));
  const { data: llms } = useQuery(GetAIModels());

  const handleCreateCube = async (data: CreateAgentFormData) => {
    const { error } = await CreateRubiksCubeFn({
      data: {
        llm: data.model,
        scramble: data.scramble,
        maxDurationMS: data.maxDuration,
      },
    });

    if (error) {
      alert(error);
      return;
    }

    _query.invalidateQueries({ queryKey: rubiksCubeKeys.lists() });
  };

  return (
    <PageLayout>
      <Websocket debug />
      <div className="flex flex-col">
        <CreateAgentModal llms={llms ?? []} onSubmit={handleCreateCube} />
        <Spacer size="xs" />
        <RubiksCubeSection
          cubes={cubes}
          isLoading={isLoading}
          limit={CubesLimit}
        />
      </div>
    </PageLayout>
  );
}
