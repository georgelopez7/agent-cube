import { cn } from "cnfast";

interface IRotationButtonsProps {
  rotations: { label: string; value: string }[];
  onRotation: (move: string) => void;
  disabled?: boolean;
  className?: string;
}

export const RotationButtons = ({
  rotations,
  onRotation,
  disabled,
  className,
}: IRotationButtonsProps) => {
  return (
    <div className={cn("grid grid-cols-3 gap-2", className)}>
      {rotations.map((rotation) => (
        <button
          type="button"
          key={rotation.value}
          onClick={() => onRotation(rotation.value)}
          disabled={disabled}
          className="border border-neutral-700 px-3 py-1 font-mono text-xs hover:bg-neutral-900 cursor-pointer disabled:opacity-40 disabled:cursor-not-allowed"
        >
          {rotation.label}
        </button>
      ))}
    </div>
  );
};
