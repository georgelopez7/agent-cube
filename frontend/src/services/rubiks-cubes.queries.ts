import { queryOptions } from "@tanstack/react-query";
import { GetRubiksCubesFn } from "./rubiks-cubes.actions";

export const rubiksCubeKeys = {
  all: ["rubiks-cubes"] as const,
  lists: () => [...rubiksCubeKeys.all, "list"] as const,
  list: (limit: number) => [...rubiksCubeKeys.lists(), limit] as const,
};

// GetRubiksCubes - Fetches all rubiks cubes with an optional limit
export const GetRubiksCubes = (limit: number = 10) =>
  queryOptions({
    queryKey: rubiksCubeKeys.list(limit),
    queryFn: async () => {
      const { cubes, error } = await GetRubiksCubesFn({
        data: { limit },
      });

      if (error) {
        throw new Error(error);
      }

      return cubes;
    },
  });
