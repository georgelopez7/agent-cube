import { Button as ButtonPrimitive } from "@base-ui/react/button";
import { cva, type VariantProps } from "class-variance-authority";

import { cn } from "cnfast";

const buttonVariants = cva(
  "group/button inline-flex shrink-0 items-center justify-center rounded-none border border-neutral-700 bg-black bg-clip-padding font-mono text-xs font-bold tracking-wide whitespace-nowrap uppercase transition-colors outline-none select-none cursor-pointer focus-visible:border-neutral-400 focus-visible:ring-1 focus-visible:ring-neutral-500 active:not-aria-[haspopup]:translate-y-px disabled:pointer-events-none disabled:opacity-40 aria-invalid:border-red-600 [&_svg]:pointer-events-none [&_svg]:shrink-0 [&_svg:not([class*='size-'])]:size-3.5",
  {
    variants: {
      variant: {
        default: "text-white hover:bg-neutral-900",
        outline:
          "border-neutral-700 bg-black text-white hover:bg-neutral-900 aria-expanded:bg-neutral-900",
        secondary:
          "border-neutral-700 bg-[#101010] text-white hover:bg-neutral-900 aria-expanded:bg-neutral-900",
        ghost:
          "border-transparent hover:border-neutral-800 hover:bg-neutral-900 hover:text-white aria-expanded:bg-neutral-900",
        destructive:
          "border-red-600 bg-red-600 text-white hover:bg-red-700 focus-visible:border-red-500",
        link: "border-transparent text-white underline underline-offset-4 hover:bg-transparent hover:underline",
      },
      size: {
        default:
          "h-10 gap-1.5 px-6 has-data-[icon=inline-end]:pr-4 has-data-[icon=inline-start]:pl-4",
        xs: "h-7 gap-1 px-3 has-data-[icon=inline-end]:pr-2 has-data-[icon=inline-start]:pl-2 [&_svg:not([class*='size-'])]:size-3",
        sm: "h-9 gap-1 px-4 has-data-[icon=inline-end]:pr-3 has-data-[icon=inline-start]:pl-3",
        lg: "h-11 gap-1.5 px-8 has-data-[icon=inline-end]:pr-5 has-data-[icon=inline-start]:pl-5",
        icon: "size-10",
        "icon-xs": "size-7 [&_svg:not([class*='size-'])]:size-3",
        "icon-sm": "size-9",
        "icon-lg": "size-11",
      },
    },
    defaultVariants: {
      variant: "default",
      size: "default",
    },
  },
);

function Button({
  className,
  variant = "default",
  size = "default",
  ...props
}: ButtonPrimitive.Props & VariantProps<typeof buttonVariants>) {
  return (
    <ButtonPrimitive
      data-slot="button"
      className={cn(buttonVariants({ variant, size, className }))}
      {...props}
    />
  );
}

export { Button, buttonVariants };
