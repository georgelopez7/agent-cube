import { cn } from "cnfast";
import { format } from "date-fns";
import { ArrowRight, CheckCircle, Circle, LoaderCircle } from "lucide-react";
import type { ReactNode } from "react";

import { getAIProviderIcon } from "#/components/(icons)/helpers";
import Spacer from "#/components/(layouts)/spacer/spacer";
import RubiksCube from "#/components/(rubiks-cube)/rubiks-cube/rubiks-cube";
import { Button } from "#/components/ui/button";
import {
  generateAlgorithm,
  RubiksCubeStatus,
  type RubiksCube as RubiksCubeType,
} from "#/domain/rubiks-cube";

interface IProps {
  cube: RubiksCubeType;
  onBegin?: (id: string) => void;
}

const statusBadgeClasses: Record<RubiksCubeStatus, string> = {
  [RubiksCubeStatus.Created]: "bg-rubiks-blue text-white",
  [RubiksCubeStatus.InProgress]: "bg-rubiks-orange text-white",
  [RubiksCubeStatus.Completed]: "bg-rubiks-green text-white",
};

const statusIcons: Record<RubiksCubeStatus, ReactNode> = {
  [RubiksCubeStatus.Created]: <Circle className="size-3.5" />,
  [RubiksCubeStatus.InProgress]: (
    <LoaderCircle className="size-3.5 animate-spin" />
  ),
  [RubiksCubeStatus.Completed]: <CheckCircle className="size-3.5" />,
};

const RubiksCubeCard = ({ cube, onBegin }: IProps) => {
  const ProviderIcon = getAIProviderIcon(cube.llm.provider);
  return (
    <div className="flex flex-col items-center border-2 px-2 rounded-lg">
      <Spacer size="xs" />
      <div className="flex w-full items-start justify-between">
        <div className="flex flex-col items-start gap-2">
          <div className="inline-flex items-center gap-1.5 rounded-md bg-secondary px-3 py-1 text-xs text-secondary-foreground">
            <ProviderIcon className="size-4" />
            <span className="text-[14px]">{cube.llm.model}</span>
          </div>
          <span
            className={cn(
              "inline-flex items-center gap-1.5 rounded-sm px-2.5 py-1 text-xs font-medium uppercase tracking-wide",
              statusBadgeClasses[cube.status],
            )}
          >
            {statusIcons[cube.status]}
            {cube.status.replace("_", " ")}
          </span>
        </div>
        <div className="flex flex-col items-end gap-0.5 text-xs text-muted-foreground tabular-nums">
          <span>{format(new Date(cube.created_at), "yyyy-MM-dd")}</span>
          <span>{format(new Date(cube.created_at), "HH:mm:ss")}</span>
        </div>
      </div>
      <Spacer size="xs" />
      {cube.status === RubiksCubeStatus.Created && (
        <Button
          size="xs"
          onClick={() => onBegin?.(cube.id)}
          className="w-full rounded-sm"
          variant="outline"
        >
          <ArrowRight />
          Begin Solving
        </Button>
      )}
      <Spacer size="xs" />
      <RubiksCube
        algorithm={generateAlgorithm(cube.cube.rotations ?? [])}
        showRotationButtons={false}
        showResetAlgoButton={false}
        showResetCameraButton={false}
        showBorder={false}
      />
      <Spacer size="xs" />
    </div>
  );
};

export default RubiksCubeCard;
