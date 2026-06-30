import { useQuery } from "@tanstack/react-query";
import { createFileRoute } from "@tanstack/react-router";

import RubiksCubeCard from "#/components/(cards)/rubiks-cube-card/rubiks-cube-card";
import { PageLayout } from "#/components/(layouts)/page-layout/page-layout";
import Spacer from "#/components/(layouts)/spacer/spacer";
import CreateAgentModal, {
  type CreateAgentFormData,
} from "#/components/(modals)/create-agent-modal/create-agent-modal";
import { Skeleton } from "#/components/ui/skeleton";
import { AIProvider, type LLM } from "#/domain/ai";
import { RUBIKS_CUBES_QUERIES } from "#/services/rubiks-cubes.queries";
import Websocket from "#/components/websocket/websocket";

const SAMPLE_LLMS: LLM[] = [
  { provider: AIProvider.OpenAI, model: "openai/gpt-4o" },
  { provider: AIProvider.Anthropic, model: "anthropic/claude-3-5-sonnet" },
  { provider: AIProvider.Google, model: "google/gemini-1.5-pro" },
];

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
      <Websocket debug />
      <div className="flex flex-col">
        <CreateAgentModal
          llms={SAMPLE_LLMS}
          onSubmit={(data: CreateAgentFormData) => console.log(data)}
        />
        <Spacer size="xs" />
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
