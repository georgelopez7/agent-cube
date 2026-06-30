import { createServerFn } from "@tanstack/react-start";
import type { XError } from "#/domain/errors";
import type { RubiksCube } from "#/domain/rubiks-cube";
import { API_BASE_URL } from "./constants";

export type GetAllRubiksCubesResult = {
  cubes: RubiksCube[];
  error: XError;
};

// getAllRubiksCubes - Fetches all rubiks cubes with an optional limit.
export const getAllRubiksCubes = createServerFn({ method: "GET" })
  .validator((data: { limit?: number }) => data)
  .handler(async ({ data }): Promise<GetAllRubiksCubesResult> => {
    const url = new URL("/api/v1/rubiks-cubes", API_BASE_URL);
    url.searchParams.set("limit", String(data.limit ?? 10));

    const response = await fetch(url.toString(), {
      headers: {
        Accept: "application/json",
      },
    });

    if (!response.ok) {
      return {
        cubes: [],
        error: `Failed to fetch rubiks cubes: ${response.status} ${response.statusText}`,
      };
    }

    type GetAllRubiksCubesResponseBody = {
      cubes: RubiksCube[] | null;
    };

    const body: GetAllRubiksCubesResponseBody = await response.json();
    return {
      cubes: body.cubes ?? [],
      error: null,
    };
  });

export type InvokeRubiksCubeAgentResult = {
  message: string;
  error: XError;
};

// invokeRubiksCubeAgent - Invokes the agent for a rubiks cube by its hex ID.
export const invokeRubiksCubeAgent = createServerFn({ method: "POST" })
  .validator((data: { id: string }) => data)
  .handler(async ({ data }): Promise<InvokeRubiksCubeAgentResult> => {
    const url = new URL(
      `/api/v1/rubiks-cubes/${data.id}/agents/invoke`,
      API_BASE_URL,
    );

    const response = await fetch(url.toString(), {
      method: "POST",
      headers: {
        Accept: "application/json",
      },
    });

    if (!response.ok) {
      return {
        message: "",
        error: `Failed to invoke rubiks cube agent: ${response.status} ${response.statusText}`,
      };
    }

    type InvokeRubiksCubeAgentResponseBody = {
      message: string;
    };

    const body: InvokeRubiksCubeAgentResponseBody = await response.json();
    return {
      message: body.message,
      error: null,
    };
  });
