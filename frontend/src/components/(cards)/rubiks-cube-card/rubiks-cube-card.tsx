import { cn } from "cnfast";
import { format, formatDistance } from "date-fns";
import {
  ArrowRight,
  Ban,
  CheckCircle,
  Circle,
  Clock,
  LoaderCircle,
  Square,
  Trash2,
} from "lucide-react";
import { type ReactNode, useCallback, useState } from "react";
import { getAIProviderIcon } from "#/components/(icons)/helpers";
import Spacer from "#/components/(layouts)/spacer/spacer";
import RubiksCube from "#/components/(rubiks-cube)/rubiks-cube/rubiks-cube";
import { useRubiksCubeStore } from "#/stores/rubiks-cube-store";
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from "#/components/ui/alert-dialog";
import { Button } from "#/components/ui/button";
import {
  generateAlgorithm,
  RubiksCubeStatus,
  type RubiksCube as RubiksCubeType,
} from "#/domain/rubiks-cube";

interface IProps {
  cube: RubiksCubeType;
  onInvoke?: (id: string) => Promise<void>;
  onStop?: (id: string) => Promise<void>;
  onDelete?: (id: string) => Promise<void>;
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
  [RubiksCubeStatus.Stopped]: {
    styles: "bg-muted text-muted-foreground",
    icon: <Ban className="size-3.5" />,
    label: "Stopped",
  },
  [RubiksCubeStatus.TimedOut]: {
    styles: "bg-yellow-500 text-white",
    icon: <Clock className="size-3.5" />,
    label: "Timed Out",
  },
};

const formatMaxDuration = (ms: number): string => {
  if (ms >= 86400000) return "Infinity";
  if (ms < 60000) {
    const seconds = Math.round(ms / 1000);
    return `${seconds} second${seconds === 1 ? "" : "s"}`;
  }
  return formatDistance(new Date(0), new Date(ms));
};

const RubiksCubeCard = ({ cube, onInvoke, onStop, onDelete }: IProps) => {
  const ProviderIcon = getAIProviderIcon(cube.llm.provider);

  // Subscribe to live status updates from the global store (e.g. websocket timeout events).
  const storeStatus = useRubiksCubeStore(
    useCallback(
      (state) => state.records.get(cube.id)?.cube.status,
      [cube.id],
    ),
  );

  const cubeStatus = storeStatus ?? cube.status;

  const algorithm = generateAlgorithm(cube.cube.rotations ?? []);
  const date = format(new Date(cube.created_at), "yyyy-MM-dd");
  const time = format(new Date(cube.created_at), "HH:mm:ss");
  const status = STATUS_CONFIG[cubeStatus];
  const maxDuration = formatMaxDuration(cube.max_duration_ms);

  const [invoking, setInvoking] = useState(false);
  const [stopping, setStopping] = useState(false);
  const [deleting, setDeleting] = useState(false);
  const [deleteDialogOpen, setDeleteDialogOpen] = useState(false);

  const handleInvoke = async (id: string) => {
    if (!onInvoke || invoking) return;

    setInvoking(true);
    try {
      await onInvoke(id);
    } finally {
      setInvoking(false);
    }
  };

  const handleStop = async (id: string) => {
    if (!onStop || stopping) return;

    setStopping(true);
    try {
      await onStop(id);
    } finally {
      setStopping(false);
    }
  };

  const handleDelete = async (id: string) => {
    if (!onDelete || deleting) return;

    setDeleting(true);
    try {
      await onDelete(id);
      setDeleteDialogOpen(false);
    } finally {
      setDeleting(false);
    }
  };

  return (
    <div className="flex flex-col items-center border-2 px-3 rounded-lg">
      <Spacer size="xs" />
      <div className="flex w-full items-start justify-between">
        <div className="flex flex-col items-start">
          <div className="inline-flex items-center gap-1.5 rounded-md bg-secondary px-3 py-1 text-xs text-secondary-foreground">
            <ProviderIcon className="size-4" />
            <span className="text-[14px]">{cube.llm.model}</span>
          </div>
          <Spacer size="xxs" />
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
        <AlertDialog open={deleteDialogOpen} onOpenChange={setDeleteDialogOpen}>
          <AlertDialogTrigger
            disabled={cubeStatus === RubiksCubeStatus.InProgress}
            render={
              <Button
                size="icon-sm"
                variant="ghost"
                disabled={cubeStatus === RubiksCubeStatus.InProgress}
                className="text-muted-foreground hover:text-destructive"
                aria-label="Delete cube"
              >
                <Trash2 className="size-5" />
              </Button>
            }
          />
          <AlertDialogContent>
            <AlertDialogHeader>
              <AlertDialogTitle>Delete Cube</AlertDialogTitle>
              <AlertDialogDescription>
                Are you sure you want to delete this cube? This action cannot be
                undone.
              </AlertDialogDescription>
            </AlertDialogHeader>
            <AlertDialogFooter>
              <AlertDialogCancel disabled={deleting}>Cancel</AlertDialogCancel>
              <AlertDialogAction
                variant="destructive"
                onClick={() => handleDelete(cube.id)}
                disabled={!onDelete || deleting}
              >
                {deleting && <LoaderCircle className="size-4 animate-spin" />}
                Delete
              </AlertDialogAction>
            </AlertDialogFooter>
          </AlertDialogContent>
        </AlertDialog>
      </div>
      <Spacer size="xxs" />
      <div className="flex w-full items-start justify-between">
        <div className="flex flex-col items-start gap-0.5 text-xs text-muted-foreground tabular-nums">
          <span>{date}</span>
          <span>{time}</span>
        </div>
        <div className="flex flex-col items-end gap-0.5 text-right text-xs text-muted-foreground tabular-nums">
          <span>Max Duration:</span>
          <span>{maxDuration}</span>
        </div>
      </div>
      <Spacer size="xs" />
      {cubeStatus === RubiksCubeStatus.Created && (
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
      {cubeStatus === RubiksCubeStatus.InProgress && (
        <Button
          size="xs"
          onClick={() => handleStop(cube.id)}
          disabled={stopping}
          className="w-full rounded-sm bg-red-600 text-white hover:bg-red-700 gap-2"
          variant="destructive"
        >
          {stopping ? (
            <LoaderCircle className="size-3 animate-spin" />
          ) : (
            <Square className="size-3 fill-current" />
          )}
          Stop Agent
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
