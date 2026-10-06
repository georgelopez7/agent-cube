import * as React from "react";
import { Input as InputPrimitive } from "@base-ui/react/input";

import { cn } from "cnfast";

function Input({ className, type, ...props }: React.ComponentProps<"input">) {
  return (
    <InputPrimitive
      type={type}
      data-slot="input"
      className={cn(
        "h-10 w-full min-w-0 border border-neutral-700 bg-[#101010] px-3 py-1 font-mono text-sm transition-colors outline-none file:inline-flex file:h-7 file:border-0 file:bg-transparent file:text-sm file:font-medium file:text-white placeholder:text-neutral-500 focus-visible:border-neutral-400 disabled:pointer-events-none disabled:cursor-not-allowed disabled:opacity-40 aria-invalid:border-red-600",
        className,
      )}
      {...props}
    />
  );
}

export { Input };
