import { queryOptions } from "@tanstack/react-query";
import { getAllRubiksCubes } from "./rubiks-cubes.action";

const RUBIKS_CUBES_QUERY_KEY = "rubiks-cubes";

export const RUBIKS_CUBES_QUERIES = {
  // getAll: Fetches all rubiks cubes with an optional limit.
  getAll: (limit = 10) =>
    queryOptions({
      queryKey: [RUBIKS_CUBES_QUERY_KEY, limit],
      queryFn: async () => {
        const { cubes, error } = await getAllRubiksCubes({
          data: { limit },
        });

        if (error) {
          throw new Error(error);
        }

        return cubes;
      },
    }),
};
