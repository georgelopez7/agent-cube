import { queryOptions } from "@tanstack/react-query";
import { getAIModels } from "./ai.actions";

const AI_QUERY_KEY = "ai";

export const aiKeys = {
  all: [AI_QUERY_KEY] as const,
  lists: () => [...aiKeys.all, "list"] as const,
  models: () => [...aiKeys.lists(), "models"] as const,
};

export const AI_QUERIES = {
  // getModels - Fetches all available AI models.
  getModels: () =>
    queryOptions({
      queryKey: aiKeys.models(),
      queryFn: async () => {
        const { models, error } = await getAIModels({ data: undefined });

        if (error) {
          throw new Error(error);
        }

        return models;
      },
      staleTime: 30_000,
    }),
};
