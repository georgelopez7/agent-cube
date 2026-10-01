import { useQuery, useQueryClient } from "@tanstack/react-query";
import { createFileRoute } from "@tanstack/react-router";
import { toast } from "sonner";
import { PageLayout } from "#/components/(layouts)/page-layout/page-layout";
import Spacer from "#/components/(layouts)/spacer/spacer";
import NewAgentDialog, {
  type NewAgentFormData,
} from "#/components/(modals)/new-agent-dialog/new-agent-dialog";
import RubiksCubeSection from "#/components/(sections)/rubiks-cube-section/rubiks-cube-section";
import Websocket from "#/components/websocket/websocket";
import { GetAIModels } from "#/services/ai.queries";
import { CreateRubiksCubeFn } from "#/services/rubiks-cubes.actions";
import {
  GetRubiksCubes,
  rubiksCubeKeys,
} from "#/services/rubiks-cubes.queries";
import { getOGImage, getSiteUrl } from "#/utils/seo";

const SEO = {
  url: `${getSiteUrl()}/agents`,
  title: "Agents | Agent Cube",
  description:
    "Create and run AI agents that attempt to solve Rubik's Cube scrambles. Benchmark LLM reasoning, track progress, and compare how different models approach the challenge.",
  image: getOGImage(),
};

export const Route = createFileRoute("/(pages)/agents")({
  component: RouteComponent,
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

function RouteComponent() {
  const CubesLimit = 10;

  const _query = useQueryClient();

  const { data: cubes, isLoading } = useQuery(GetRubiksCubes(CubesLimit));
  const { data: llms } = useQuery(GetAIModels());

  const handleCreateCube = async (data: NewAgentFormData) => {
    const { error } = await CreateRubiksCubeFn({
      data: {
        llm: data.model,
        scramble: data.scramble,
        maxDurationMS: data.maxDuration,
      },
    });

    if (error) {
      toast.error(error);
      return;
    }

    _query.invalidateQueries({ queryKey: rubiksCubeKeys.lists() });
  };

  return (
    <PageLayout>
      <Websocket />
      <div className="flex flex-col">
        <NewAgentDialog llms={llms ?? []} onSubmit={handleCreateCube} />
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
