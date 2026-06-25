import { cn } from "cnfast";
import { useMemo } from "react";

interface IProps {
  className?: string;
  size?: "sm" | "md" | "lg";
  animationSpeed?: "slow" | "medium" | "fast";
  randomize?: boolean;
}

type CubeTile = {
  id: string;
  color: string;
};

const SIZE_CLASSES = {
  sm: "w-40 h-40 md:w-48 md:h-48",
  md: "w-64 h-64 md:w-80 md:h-80",
  lg: "w-80 h-80 md:w-96 md:h-96",
};

const ANIMATION_CLASSES = {
  slow: "animate-[pulse_4s_cubic-bezier(0.4,0,0.6,1)_infinite]",
  medium: "animate-[pulse_2s_cubic-bezier(0.4,0,0.6,1)_infinite]",
  fast: "animate-[pulse_1s_cubic-bezier(0.4,0,0.6,1)_infinite]",
};

const COLORS = [
  "bg-rubiks-red",
  "bg-rubiks-blue",
  "bg-rubiks-green",
  "bg-rubiks-orange",
  "bg-rubiks-yellow",
  "bg-rubiks-white",
  "bg-rubiks-blue",
  "bg-rubiks-green",
  "bg-rubiks-yellow",
];

const CUBE_TILES: CubeTile[] = COLORS.map((color, index) => ({
  id: `tile-${index}`,
  color,
}));

const shuffle = (tiles: CubeTile[]): CubeTile[] => {
  const shuffled = [...tiles];

  for (let i = shuffled.length - 1; i > 0; i--) {
    const j = Math.floor(Math.random() * (i + 1));
    [shuffled[i], shuffled[j]] = [shuffled[j], shuffled[i]];
  }

  return shuffled;
};

export const CubePulse = ({
  className = "",
  size = "md",
  animationSpeed = "medium",
  randomize = false,
}: IProps) => {
  const tiles = useMemo(() => {
    if (!randomize) return CUBE_TILES;
    return shuffle(CUBE_TILES);
  }, [randomize]);

  return (
    <div
      className={cn(
        "relative transform rotate-12 perspective-1000",
        SIZE_CLASSES[size],
        className,
      )}
    >
      <div className="absolute inset-0 bg-linear-to-br from-slate-500/20 to-transparent rounded-lg blur-3xl" />
      <div
        className={cn(
          "relative w-full h-full grid grid-cols-3 grid-rows-3 gap-2 p-2 bg-slate-900/80 rounded-lg border border-slate-500/30 shadow-[0_0_10px_rgba(100,116,139,0.2)]",
          ANIMATION_CLASSES[animationSpeed],
        )}
      >
        {tiles.map(({ id, color }) => (
          <div
            key={id}
            className={cn(
              color,
              "rounded-sm shadow-[inset_0_0_10px_rgba(0,0,0,0.5)]",
            )}
          />
        ))}
      </div>
    </div>
  );
};

export default CubePulse;
