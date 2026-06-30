import RubiksCubeCard from "#/components/(cards)/rubiks-cube-card/rubiks-cube-card";
import { Skeleton } from "#/components/ui/skeleton";
import { type RubiksCube, RubiksCubeStatus } from "#/domain/rubiks-cube";
import { useInvokeRubiksCubeAgent } from "#/services/rubiks-cubes.queries";
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
  // STORE
  const setStatus = useRubiksCubeStore((state) => state.SetStatus);

  // MUTATIONS
  const { mutate: invokeAgent } = useInvokeRubiksCubeAgent();

  // HANDLERS
  const handleInvoke = async (id: string) => {
    invokeAgent(id, {
      onSuccess: () => setStatus(id, RubiksCubeStatus.InProgress),
    });
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
        <RubiksCubeCard key={cube.id} cube={cube} onInvoke={handleInvoke} />
      ))}
    </div>
  );
};

export default RubiksCubeSection;
