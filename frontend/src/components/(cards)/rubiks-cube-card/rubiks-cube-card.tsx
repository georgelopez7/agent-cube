import { cn } from "cnfast";
import { format } from "date-fns";
import { ArrowRight, CheckCircle, Circle, LoaderCircle } from "lucide-react";
import { useState, type ReactNode } from "react";

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
  onInvoke?: (id: string) => Promise<void>;
}

const STATUS_CONFIG: Record<
  RubiksCubeStatus,
  { styles: string; icon: ReactNode; label: string }
> = {
  [RubiksCubeStatus.Created]: {
    styles: "bg-rubiks-blue text-white",
    icon: <Circle className="size-3.5" />,
    label: "Created",
  },
  [RubiksCubeStatus.InProgress]: {
    styles: "bg-rubiks-orange text-white",
    icon: <LoaderCircle className="size-3.5 animate-spin" />,
    label: "In Progress",
  },
  [RubiksCubeStatus.Completed]: {
    styles: "bg-rubiks-green text-white",
    icon: <CheckCircle className="size-3.5" />,
    label: "Completed",
  },
};

const RubiksCubeCard = ({ cube, onInvoke }: IProps) => {
  const ProviderIcon = getAIProviderIcon(cube.llm.provider);

  const algorithm = generateAlgorithm(cube.cube.rotations ?? []);
  const date = format(new Date(cube.created_at), "yyyy-MM-dd");
  const time = format(new Date(cube.created_at), "HH:mm:ss");
  const status = STATUS_CONFIG[cube.status];

  const [invoking, setInvoking] = useState(false);

  const handleInvoke = async (id: string) => {
    if (invoking) return;

    setInvoking(true);
    try {
      await onInvoke?.(id);
    } finally {
      setInvoking(false);
    }
  };

  return (
    <div className="flex flex-col items-center border-2 px-3 rounded-lg">
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
              status.styles,
            )}
          >
            {status.icon}
            {status.label}
          </span>
        </div>
        <div className="flex flex-col items-end gap-0.5 text-xs text-muted-foreground tabular-nums">
          <span>{date}</span>
          <span>{time}</span>
        </div>
      </div>
      <Spacer size="xs" />
      {cube.status === RubiksCubeStatus.Created && (
        <Button
          size="xs"
          onClick={() => handleInvoke(cube.id)}
          disabled={invoking}
          className="w-full rounded-sm"
          variant="outline"
        >
          {invoking ? (
            <LoaderCircle className="size-4 animate-spin" />
          ) : (
            <ArrowRight />
          )}
          Invoke Agent
        </Button>
      )}
      <Spacer size="xs" />
      <RubiksCube
        cube={cube}
        algorithm={algorithm}
        showRotationButtons={false}
        showResetAlgoButton={false}
        showBorder={false}
      />
      <Spacer size="xs" />
    </div>
  );
};

export default RubiksCubeCard;
