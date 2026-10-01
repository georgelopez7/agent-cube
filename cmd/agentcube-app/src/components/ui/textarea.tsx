import * as React from "react";

import { cn } from "cnfast";

function Textarea({ className, ...props }: React.ComponentProps<"textarea">) {
  return (
    <textarea
      data-slot="textarea"
      className={cn(
        "flex field-sizing-content min-h-16 w-full resize-none rounded-none border border-neutral-700 bg-[#101010] px-3 py-3 font-mono text-sm transition-colors outline-none placeholder:text-neutral-500 focus-visible:border-neutral-400 disabled:cursor-not-allowed disabled:opacity-40 aria-invalid:border-red-600",
        className,
      )}
      {...props}
    />
  );
}

export { Textarea };
