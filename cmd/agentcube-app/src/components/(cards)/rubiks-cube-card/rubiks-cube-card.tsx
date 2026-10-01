import { cn } from "cnfast";
import { format, formatDistance } from "date-fns";
import {
  Ban,
  CheckCircle,
  Circle,
  Clock,
  LoaderCircle,
  MousePointerClick,
  Square,
  Trash2,
} from "lucide-react";
import { type ReactNode, useCallback, useState } from "react";
import { getAIProviderIcon } from "#/components/(icons)/helpers";
import Spacer from "#/components/(layouts)/spacer/spacer";
import RubiksCube from "#/components/(rubiks-cube)/rubiks-cube/rubiks-cube";
import { Terminal } from "#/components/terminal/terminal";
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
import { useRubiksCubeStore } from "#/stores/rubiks-cube-store";

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
    styles: "border border-neutral-700 bg-neutral-800 text-white",
    icon: <Circle className="size-3.5" />,
    label: "Created",
  },
  [RubiksCubeStatus.InProgress]: {
    styles: "border border-yellow-600 bg-yellow-600 text-white",
    icon: <LoaderCircle className="size-3.5 animate-spin" />,
    label: "In Progress",
  },
  [RubiksCubeStatus.Completed]: {
    styles: "border border-green-600 bg-green-600 text-white",
    icon: <CheckCircle className="size-3.5" />,
    label: "Completed",
  },
  [RubiksCubeStatus.Stopped]: {
    styles: "border border-neutral-800 bg-neutral-800 text-white",
    icon: <Ban className="size-3.5" />,
    label: "Stopped",
  },
  [RubiksCubeStatus.TimedOut]: {
    styles: "border border-yellow-600 bg-yellow-600 text-white",
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
    useCallback((state) => state.records.get(cube.id)?.cube.status, [cube.id]),
  );

  const cubeStatus = storeStatus ?? cube.status;

  const reasoning = useRubiksCubeStore(
    useCallback((state) => state.reasoning.get(cube.id) ?? "", [cube.id]),
  );

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
    <div className="flex h-full flex-col bg-[#101010] p-1">
      <div className="flex h-full w-full flex-col p-2">
        <div className="flex w-full items-center justify-center gap-3 border border-neutral-700 bg-black px-3 py-1 font-mono text-sm font-bold text-white">
          <ProviderIcon className="size-3.5 text-neutral-400" />
          <span>{cube.llm.model}</span>
        </div>
        <Spacer size="xs" />
        <div className="flex w-full flex-col border border-neutral-700 bg-black p-3">
          <div className="flex w-full items-center justify-between gap-3">
            <span
              className={cn(
                "inline-flex items-center gap-1.5 px-2.5 py-1 font-mono text-xs font-bold uppercase tracking-wide",
                status.styles,
              )}
            >
              {status.icon}
              {status.label}
            </span>
            <span className="flex flex-col text-right text-xs text-white tabular-nums">
              <span>{date}</span>
              <span>{time}</span>
            </span>
          </div>
          <Spacer size="xxs" />
          <div className="flex w-full justify-start">
            <span className="text-left text-xs text-muted-foreground tabular-nums">
              Max duration:{" "}
              <span
                className={cn(
                  "text-white",
                  maxDuration === "1 minute" && "font-bold",
                )}
              >
                {maxDuration}
              </span>
            </span>
          </div>
        </div>
        <Spacer size="xs" />
        <div className="flex w-full items-center justify-between gap-2">
          <div className="flex min-w-0 flex-1 gap-2">
            {cubeStatus === RubiksCubeStatus.Created && (
              <Button
                size="xs"
                onClick={() => handleInvoke(cube.id)}
                disabled={invoking}
                className="flex-1 gap-2"
                variant="outline"
              >
                {invoking ? (
                  <LoaderCircle className="size-4 animate-spin" />
                ) : (
                  <MousePointerClick className="size-3 -scale-x-100" />
                )}
                Invoke Agent
              </Button>
            )}
            {cubeStatus === RubiksCubeStatus.InProgress && (
              <Button
                size="xs"
                onClick={() => handleStop(cube.id)}
                disabled={stopping}
                className="flex-1 gap-2"
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
          </div>
          <AlertDialog
            open={deleteDialogOpen}
            onOpenChange={setDeleteDialogOpen}
          >
            <AlertDialogTrigger
              disabled={cubeStatus === RubiksCubeStatus.InProgress}
              render={
                <Button
                  size="xs"
                  variant="ghost"
                  disabled={cubeStatus === RubiksCubeStatus.InProgress}
                  className="gap-2 border border-neutral-700 text-white hover:border-red-600 hover:bg-transparent hover:text-red-600"
                >
                  <Trash2 className="size-3" />
                  Delete
                </Button>
              }
            />
            <AlertDialogContent>
              <AlertDialogHeader className="gap-2 border-0 p-0">
                <AlertDialogTitle className="sr-only">
                  Delete Cube
                </AlertDialogTitle>
                <AlertDialogDescription className="border border-neutral-700 bg-black px-4 py-3 text-sm text-white">
                  Are you sure you want to delete this cube? This action cannot
                  be undone.
                </AlertDialogDescription>
              </AlertDialogHeader>
              <AlertDialogFooter>
                <AlertDialogCancel size="xs" disabled={deleting}>
                  Cancel
                </AlertDialogCancel>
                <AlertDialogAction
                  size="xs"
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
        <Spacer size="xs" />
        <div className="flex w-full justify-center border border-neutral-700 bg-black p-2">
          <RubiksCube
            cube={cube}
            algorithm={algorithm}
            showRotationButtons={false}
            showResetAlgoButton={false}
            showBorder={false}
          />
        </div>
        <Terminal
          className="mt-3 w-full"
          logs={reasoning ? [reasoning] : []}
          title="agent.reasoning"
        />
      </div>
    </div>
  );
};

export default RubiksCubeCard;
