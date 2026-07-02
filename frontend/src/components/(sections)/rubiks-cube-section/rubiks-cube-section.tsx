import { useQueryClient } from "@tanstack/react-query";
import RubiksCubeCard from "#/components/(cards)/rubiks-cube-card/rubiks-cube-card";
import { Skeleton } from "#/components/ui/skeleton";
import type { RubiksCube } from "#/domain/rubiks-cube";
import {
  DeleteRubiksCubeFn,
  InvokeRubiksCubeAgentFn,
} from "#/services/rubiks-cubes.actions";
import { rubiksCubeKeys } from "#/services/rubiks-cubes.queries";

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
    const { error } = await InvokeRubiksCubeAgentFn({ data: { id } });
    if (error) {
      alert(error);
    }

    _query.invalidateQueries({ queryKey: rubiksCubeKeys.lists() });
  };

  const handleDelete = async (id: string) => {
    const { error } = await DeleteRubiksCubeFn({ data: { id } });
    if (error) {
      alert(error);
      return;
    }

    _query.invalidateQueries({ queryKey: rubiksCubeKeys.lists() });
  };

  if (isLoading) {
    return (
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
        {Array.from({ length: limit }).map((_, index) => (
          <Skeleton
            key={`skeleton-cube-${index}`}
            className="h-96 w-full rounded-lg"
          />
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
          onDelete={handleDelete}
        />
      ))}
    </div>
  );
};

export default RubiksCubeSection;
