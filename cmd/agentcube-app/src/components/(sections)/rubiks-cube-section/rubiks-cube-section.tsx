import { useQueryClient } from "@tanstack/react-query";
import { toast } from "sonner";
import RubiksCubeCard from "#/components/(cards)/rubiks-cube-card/rubiks-cube-card";
import { Skeleton } from "#/components/ui/skeleton";
import type { RubiksCube } from "#/domain/rubiks-cube";
import {
  DeleteRubiksCubeFn,
  InvokeRubiksCubeAgentFn,
  StopRubiksCubeAgentFn,
} from "#/services/rubiks-cubes.actions";
import { rubiksCubeKeys } from "#/services/rubiks-cubes.queries";
import { useRubiksCubeStore } from "#/stores/rubiks-cube-store";

interface RubiksCubeSectionProps {
  cubes?: RubiksCube[];
  isLoading: boolean;
  limit?: number;
}

const RubiksCubeSection = ({
  cubes,
  isLoading,
  limit = 10,
}: RubiksCubeSectionProps) => {
  const _query = useQueryClient();

  const handleInvoke = async (id: string) => {
    useRubiksCubeStore.getState().ClearReasoning(id);

    // invoked_at is written solely by RunAgent on the backend and returned on
    // the /invoke response; use only that value for time elapsed (on refresh
    // the card reads invoked_at from the cube record via the list query).
    const { cube, error } = await InvokeRubiksCubeAgentFn({ data: { id } });
    if (error) {
      toast.error(error);
    }

    if (cube?.invoked_at) {
      useRubiksCubeStore.getState().SetInvokedAt(id, cube.invoked_at);
    }

    _query.invalidateQueries({ queryKey: rubiksCubeKeys.lists() });
  };

  const handleStop = async (id: string) => {
    const { error } = await StopRubiksCubeAgentFn({ data: { id } });
    if (error) {
      toast.error(error);
    }

    _query.invalidateQueries({ queryKey: rubiksCubeKeys.lists() });
  };

  const handleDelete = async (id: string) => {
    const { error } = await DeleteRubiksCubeFn({ data: { id } });
    if (error) {
      toast.error(error);
      return;
    }

    _query.invalidateQueries({ queryKey: rubiksCubeKeys.lists() });
  };

  if (isLoading) {
    return (
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
        {Array.from({ length: limit }).map((_, index) => (
          <Skeleton key={`skeleton-cube-${index}`} className="h-96 w-full" />
        ))}
      </div>
    );
  }

  return (
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
      {cubes?.map((cube) => (
        <RubiksCubeCard
          key={cube.id}
          cube={cube}
          onInvoke={handleInvoke}
          onStop={handleStop}
          onDelete={handleDelete}
        />
      ))}
    </div>
  );
};

export default RubiksCubeSection;
