import cn from "cnfast";
import { ChevronDown } from "lucide-react";
import { useState } from "react";
import Spacer from "#/components/(layouts)/spacer/spacer";
import { Collapsible, CollapsibleContent } from "#/components/ui/collapsible";
import { ScrollArea } from "#/components/ui/scroll-area";
import type { CubeRotation } from "#/domain/rubiks-cube";

interface IProps {
  rotations: CubeRotation[];
  disabled?: boolean;
  className?: string;
}

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
          className="inline-flex items-center gap-1 p-0 font-mono text-sm text-muted-foreground hover:text-foreground disabled:opacity-50 disabled:cursor-not-allowed cursor-pointer"
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
              return (
                <div
                  key={`${rotation.index}-${rotation.created_at}-${rotation.rotation}`}
                  className={cn(
                    "flex size-8 items-center justify-center border border-neutral-700 bg-transparent font-mono text-sm font-extrabold leading-none text-white text-center",
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
