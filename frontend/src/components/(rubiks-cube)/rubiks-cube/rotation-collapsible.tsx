import cn from "cnfast";
import { ChevronDown } from "lucide-react";
import { useState } from "react";
import Spacer from "#/components/(layouts)/spacer/spacer";
import { Collapsible, CollapsibleContent } from "#/components/ui/collapsible";
import { ScrollArea } from "#/components/ui/scroll-area";
import type { CubeRotation, Rotation } from "#/domain/rubiks-cube";

interface IProps {
  rotations: CubeRotation[];
  disabled?: boolean;
  className?: string;
}

type RotationStyle = {
  background: string;
  text: string;
  border: string;
};

// ROTATION_CONFIG - maps each valid rotation to its visual style metadata
const ROTATION_CONFIG: Record<Rotation, RotationStyle> = {
  F: {
    background: "bg-rubiks-green",
    text: "text-rubiks-white",
    border: "",
  },
  "F'": {
    background: "bg-transparent",
    text: "text-foreground",
    border: "border-2 border-foreground",
  },
  R: {
    background: "bg-rubiks-red",
    text: "text-rubiks-white",
    border: "",
  },
  "R'": {
    background: "bg-transparent",
    text: "text-foreground",
    border: "border-2 border-foreground",
  },
  U: {
    background: "bg-rubiks-white",
    text: "text-black",
    border: "border-2 border-border",
  },
  "U'": {
    background: "bg-transparent",
    text: "text-foreground",
    border: "border-2 border-foreground",
  },
  L: {
    background: "bg-rubiks-orange",
    text: "text-rubiks-white",
    border: "",
  },
  "L'": {
    background: "bg-transparent",
    text: "text-foreground",
    border: "border-2 border-foreground",
  },
  B: {
    background: "bg-rubiks-blue",
    text: "text-rubiks-white",
    border: "",
  },
  "B'": {
    background: "bg-transparent",
    text: "text-foreground",
    border: "border-2 border-foreground",
  },
  D: {
    background: "bg-rubiks-yellow",
    text: "text-black",
    border: "",
  },
  "D'": {
    background: "bg-transparent",
    text: "text-foreground",
    border: "border-2 border-foreground",
  },
};

export const RotationCollapsible = ({
  rotations,
  disabled,
  className,
}: IProps) => {
  const [open, setOpen] = useState(true);

  return (
    <Collapsible open={open} onOpenChange={setOpen} className={className}>
      <div className="flex justify-center">
        <button
          type="button"
          disabled={disabled}
          onClick={() => setOpen((prev) => !prev)}
          className="inline-flex items-center gap-1 p-0 text-sm text-muted-foreground hover:text-foreground disabled:opacity-50 disabled:cursor-not-allowed cursor-pointer"
        >
          <p>{open ? "Hide rotations" : "Show rotations"}</p>
          <ChevronDown
            className={cn(
              "size-4 shrink-0 leading-none transition-transform duration-200",
              open && "rotate-180",
            )}
          />
        </button>
      </div>
      <CollapsibleContent>
        <Spacer size="xs" />
        <ScrollArea className="h-31 w-48">
          <div className="grid grid-cols-5 gap-2">
            {rotations.map((rotation) => {
              const style = ROTATION_CONFIG[rotation.rotation];
              return (
                <div
                  key={rotation.index}
                  className={cn(
                    "flex items-center justify-center size-8 rounded-sm text-sm font-mono leading-none text-center",
                    style.background,
                    style.text,
                    style.border,
                  )}
                  title={rotation.rotation}
                >
                  {rotation.rotation}
                </div>
              );
            })}
          </div>
        </ScrollArea>
      </CollapsibleContent>
    </Collapsible>
  );
};
