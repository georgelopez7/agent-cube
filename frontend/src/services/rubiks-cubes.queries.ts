import {
  queryOptions,
  useMutation,
  useQueryClient,
} from "@tanstack/react-query";
import type { LLM } from "#/domain/ai";
import {
  createRubiksCube,
  getAllRubiksCubes,
  invokeRubiksCubeAgent,
} from "./rubiks-cubes.actions";

const RUBIKS_CUBES_QUERY_KEY = "rubiks-cubes";

export const rubiksCubesKeys = {
  all: [RUBIKS_CUBES_QUERY_KEY] as const,
  lists: () => [...rubiksCubesKeys.all, "list"] as const,
  list: (limit: number) => [...rubiksCubesKeys.lists(), limit] as const,
};

export const RUBIKS_CUBES_QUERIES = {
  // getAll - Fetches all rubiks cubes with an optional limit.
  getAll: (limit = 10) =>
    queryOptions({
      queryKey: rubiksCubesKeys.list(limit),
      queryFn: async () => {
        const { cubes, error } = await getAllRubiksCubes({
          data: { limit },
        });

        if (error) {
          throw new Error(error);
        }

        return cubes;
      },
      staleTime: 30_000,
    }),
};

// useCreateRubiksCube - Creates a new rubiks cube and invalidates the cubes list.
export const useCreateRubiksCube = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: async (data: { llm: LLM; scramble: number }) => {
      const { cube, error } = await createRubiksCube({ data });

      if (error) {
        throw new Error(error);
      }

      return cube;
    },
    onSettled: () => {
      queryClient.invalidateQueries({ queryKey: rubiksCubesKeys.all });
    },
  });
};

// useInvokeRubiksCubeAgent - Invokes the agent for a rubiks cube by ID.
export const useInvokeRubiksCubeAgent = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: async (id: string) => {
      const { message, error } = await invokeRubiksCubeAgent({
        data: { id },
      });

      if (error) {
        throw new Error(error);
      }

      return message;
    },
    onSettled: () => {
      queryClient.invalidateQueries({ queryKey: rubiksCubesKeys.all });
    },
  });
};
