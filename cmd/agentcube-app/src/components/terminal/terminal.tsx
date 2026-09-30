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
  const bottomRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    if (bottomRef.current) {
      bottomRef.current.scrollIntoView({ behavior: "smooth" });
    }
  }, [logs]);

  return (
    <div
      className={cn(
        "border border-border rounded-lg font-mono text-sm bg-black overflow-hidden",
        className,
      )}
    >
      <div className="flex items-center gap-2 px-4 py-2 bg-zinc-900 border-b border-zinc-800">
        <TerminalIcon className="size-4 text-zinc-400" />
        <span className="text-xs text-zinc-400">{title}</span>
      </div>
      <ScrollArea className="h-48 p-4">
        <div className="text-white space-y-1">
          {logs.map((log, index) => (
            <div className="whitespace-pre" key={index}>
              <p>{log}</p>
            </div>
          ))}
        </div>
        {showCursor && <div className="terminal-cursor" />}
        <div ref={bottomRef} />
        <ScrollBar orientation="horizontal" />
      </ScrollArea>
    </div>
  );
}
