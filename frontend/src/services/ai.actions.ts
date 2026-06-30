import { createServerFn } from "@tanstack/react-start";
import { AIProvider, type LLM } from "#/domain/ai";
import type { XError } from "#/domain/errors";
import { API_BASE_URL } from "./_constants";

export type GetAIModelsResult = {
  models: LLM[];
  error: XError;
};

// getAIModels - Fetches all available AI models.
export const getAIModels = createServerFn({ method: "GET" })
  .validator((data: undefined) => data)
  .handler(async (): Promise<GetAIModelsResult> => {
    const url = new URL("/api/v1/ai/models", API_BASE_URL);

    const response = await fetch(url.toString(), {
      headers: {
        Accept: "application/json",
      },
    });

    if (!response.ok) {
      return {
        models: [],
        error: `Failed to fetch AI models: ${response.status} ${response.statusText}`,
      };
    }

    type GetAIModelsResponseBody = {
      models: Array<{ provider: AIProvider; model: string }> | null;
    };

    const body: GetAIModelsResponseBody = await response.json();
    return {
      models: body.models ?? [],
      error: null,
    };
  });
