import { createServerFn } from "@tanstack/react-start";
import type { XError } from "#/domain/errors";
import type { RubiksCube } from "#/domain/rubiks-cube";
import { API_BASE_URL } from "./constants";

type GetAllRubiksCubesResponseBody = {
  cubes: RubiksCube[] | null;
};

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

    const body: GetAllRubiksCubesResponseBody = await response.json();
    return {
      cubes: body.cubes ?? [],
      error: null,
    };
  });
