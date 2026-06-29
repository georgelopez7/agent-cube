import cn from "cnfast";
import { ChevronDown } from "lucide-react";
import { useState } from "react";
import Spacer from "#/components/(layouts)/spacer/spacer";
import { Collapsible, CollapsibleContent } from "#/components/ui/collapsible";
import { ScrollArea } from "#/components/ui/scroll-area";

interface Rotation {
  label: string;
  value: string;
}

interface IProps {
  rotations: Rotation[];
  disabled?: boolean;
  className?: string;
}

const RUBIKS_BACKGROUND: Record<string, string> = {
  F: "bg-rubiks-green",
  R: "bg-rubiks-red",
  U: "bg-rubiks-white",
  L: "bg-rubiks-orange",
  B: "bg-rubiks-blue",
  D: "bg-rubiks-yellow",
};

const getRotationStyle = (value: string) => {
  const isPrime = value.includes("'");
  const base = value.replace("'", "");

  if (isPrime) {
    return {
      background: "bg-transparent",
      text: "text-foreground",
      border: "border-2 border-foreground",
    };
  }

  return {
    background: RUBIKS_BACKGROUND[base] ?? "bg-muted",
    text: base === "U" || base === "D" ? "text-black" : "text-rubiks-white",
    border: base === "U" ? "border-2 border-border" : "",
  };
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
              const style = getRotationStyle(rotation.value);
              return (
                <div
                  key={rotation.value}
                  className={cn(
                    "flex items-center justify-center size-8 rounded-sm text-sm font-mono leading-none text-center",
                    style.background,
                    style.text,
                    style.border,
                  )}
                  title={rotation.label}
                >
                  {rotation.label}
                </div>
              );
            })}
          </div>
        </ScrollArea>
      </CollapsibleContent>
    </Collapsible>
  );
};
