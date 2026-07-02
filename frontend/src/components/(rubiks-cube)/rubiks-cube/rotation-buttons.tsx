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
          className="px-3 py-1 text-sm border-2 rounded hover:bg-accent cursor-pointer disabled:opacity-50 disabled:cursor-not-allowed"
        >
          {rotation.label}
        </button>
      ))}
    </div>
  );
};
