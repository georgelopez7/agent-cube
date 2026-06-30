import { useQuery } from "@tanstack/react-query";
import { createFileRoute } from "@tanstack/react-router";

import { PageLayout } from "#/components/(layouts)/page-layout/page-layout";
import Spacer from "#/components/(layouts)/spacer/spacer";
import CreateAgentModal, {
  type CreateAgentFormData,
} from "#/components/(modals)/create-agent-modal/create-agent-modal";
import RubiksCubeSection from "#/components/(sections)/rubiks-cube-section/rubiks-cube-section";
import Websocket from "#/components/websocket/websocket";
import { AIProvider, type LLM } from "#/domain/ai";
import { AI_QUERIES } from "#/services/ai.queries";
import { RUBIKS_CUBES_QUERIES } from "#/services/rubiks-cubes.queries";

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
  const { data: llms = SAMPLE_LLMS } = useQuery(AI_QUERIES.getModels());

  return (
    <PageLayout>
      <Websocket debug />
      <div className="flex flex-col">
        <CreateAgentModal
          llms={llms}
          onSubmit={(data: CreateAgentFormData) => console.log(data)}
        />
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
