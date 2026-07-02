import { cn } from "cnfast";
import { useId, useMemo } from "react";

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

const hashStringToSeed = (value: string): number => {
  let hash = 0;

  for (let i = 0; i < value.length; i++) {
    const char = value.charCodeAt(i);
    hash = (hash << 5) - hash + char;
    hash |= 0;
  }

  return Math.abs(hash);
};

const createSeededRandom = (seed: number) => {
  let state = seed;

  return () => {
    state = (state * 9301 + 49297) % 233280;
    return state / 233280;
  };
};

const shuffle = (tiles: CubeTile[], seed?: number): CubeTile[] => {
  const shuffled = [...tiles];
  const random = seed !== undefined ? createSeededRandom(seed) : Math.random;

  for (let i = shuffled.length - 1; i > 0; i--) {
    const j = Math.floor(random() * (i + 1));
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
  const id = useId();

  const tiles = useMemo(() => {
    if (!randomize) return CUBE_TILES;
    return shuffle(CUBE_TILES, hashStringToSeed(id));
  }, [randomize, id]);

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
