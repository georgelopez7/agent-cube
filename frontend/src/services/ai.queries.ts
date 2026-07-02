import { queryOptions } from "@tanstack/react-query";
import { GetAIModelsFn } from "./ai.actions";

export const aiKeys = {
  all: ["ai"] as const,
  models: () => [...aiKeys.all, "models"] as const,
};

// GetAIModels - Fetches all available AI models
export const GetAIModels = () =>
  queryOptions({
    queryKey: aiKeys.models(),
    queryFn: async () => {
      const { models, error } = await GetAIModelsFn({ data: undefined });

      if (error) {
        throw new Error(error);
      }

      return models;
    },
  });
