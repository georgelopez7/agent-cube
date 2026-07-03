import { createServerFn } from "@tanstack/react-start";
import type { LLM } from "#/domain/ai";
import type { XError } from "#/domain/errors";
import type { RubiksCube } from "#/domain/rubiks-cube";
import { API_BASE_URL } from "./_constants";

export type GetAllRubiksCubesResult = {
  cubes: RubiksCube[];
  error: XError;
};

// GetRubiksCubesFn - Fetches all rubiks cubes with an optional limit.
export const GetRubiksCubesFn = createServerFn({ method: "GET" })
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

export type CreateRubiksCubeResult = {
  cube: RubiksCube;
  error: XError;
};

// CreateRubiksCubeFn - Creates a new rubiks cube configured for a specific LLM.
export const CreateRubiksCubeFn = createServerFn({ method: "POST" })
  .validator(
    (data: { llm: LLM; scramble: number; maxDurationMS: number }) => data,
  )
  .handler(async ({ data }): Promise<CreateRubiksCubeResult> => {
    const url = new URL("/api/v1/rubiks-cubes", API_BASE_URL);

    const response = await fetch(url.toString(), {
      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        llm: data.llm,
        scramble: data.scramble,
        max_duration_ms: data.maxDurationMS,
      }),
    });

    if (!response.ok) {
      return {
        cube: null as unknown as RubiksCube,
        error: `Failed to create rubiks cube: ${response.status} ${response.statusText}`,
      };
    }

    type CreateRubiksCubeResponseBody = {
      cube: RubiksCube | null;
    };

    const body: CreateRubiksCubeResponseBody = await response.json();
    return {
      cube: body.cube ?? (null as unknown as RubiksCube),
      error: null,
    };
  });

export type DeleteRubiksCubeResult = {
  success: boolean;
  error: XError;
};

// DeleteRubiksCubeFn - Deletes a rubiks cube by its hex ID.
export const DeleteRubiksCubeFn = createServerFn({ method: "POST" })
  .validator((data: { id: string }) => data)
  .handler(async ({ data }): Promise<DeleteRubiksCubeResult> => {
    const url = new URL(`/api/v1/rubiks-cubes/${data.id}`, API_BASE_URL);

    const response = await fetch(url.toString(), {
      method: "DELETE",
      headers: {
        Accept: "application/json",
      },
    });

    if (!response.ok) {
      return {
        success: false,
        error: `Failed to delete rubiks cube: ${response.status} ${response.statusText}`,
      };
    }

    return {
      success: true,
      error: null,
    };
  });

export type InvokeRubiksCubeAgentResult = {
  message: string;
  error: XError;
};

// InvokeRubiksCubeAgentFn - Invokes the agent for a rubiks cube by its hex ID.
export const InvokeRubiksCubeAgentFn = createServerFn({ method: "POST" })
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

export type StopRubiksCubeAgentResult = {
  message: string;
  error: XError;
};

// StopRubiksCubeAgentFn - Stops the running agent for a rubiks cube by its hex ID.
export const StopRubiksCubeAgentFn = createServerFn({ method: "POST" })
  .validator((data: { id: string }) => data)
  .handler(async ({ data }): Promise<StopRubiksCubeAgentResult> => {
    const url = new URL(
      `/api/v1/rubiks-cubes/${data.id}/agents/stop`,
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
        error: `Failed to stop rubiks cube agent: ${response.status} ${response.statusText}`,
      };
    }

    type StopRubiksCubeAgentResponseBody = {
      message: string;
    };

    const body: StopRubiksCubeAgentResponseBody = await response.json();
    return {
      message: body.message,
      error: null,
    };
  });
