import { Terminal as TerminalIcon } from "lucide-react";
import { ScrollArea, ScrollBar } from "@/components/ui/scroll-area";
import { useEffect, useRef } from "react";
import cn from "cnfast";

interface IProps {
  className?: string;
  logs: string[];
  title?: string;
  showCursor?: boolean;
}

export function Terminal({
  className,
  logs,
  title = "terminal.log",
  showCursor = true,
}: IProps) {
  const viewportRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const viewport = viewportRef.current;
    if (!viewport) return;

    viewport.scrollTop = viewport.scrollHeight;
  }, [logs]);

  return (
    <div
      className={cn(
        "overflow-hidden border border-neutral-700 bg-black font-mono text-xs",
        className,
      )}
    >
      <div className="flex items-center gap-2 border-b border-neutral-800 bg-[#101010] px-4 py-2">
        <TerminalIcon className="size-4 text-neutral-400" />
        <span className="text-xs text-neutral-400">{title}</span>
      </div>
      <ScrollArea className="h-48 p-4" viewportRef={viewportRef}>
        <div className="space-y-1 text-neutral-400">
          {logs.map((log, index) => (
            <div className="whitespace-pre-wrap break-words" key={index}>
              <p>{log}</p>
            </div>
          ))}
        </div>
        {showCursor && <div className="terminal-cursor" />}
        <ScrollBar orientation="horizontal" />
      </ScrollArea>
    </div>
  );
}
