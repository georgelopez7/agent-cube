import React, { forwardRef } from "react";

const XAIIcon = forwardRef<SVGSVGElement, React.HTMLProps<SVGSVGElement>>(
  (props, ref) => {
    return (
      <svg
        xmlns="http://www.w3.org/2000/svg"
        fill="#fff"
        viewBox="62.7 17.7 136.4 149.8"
        {...props}
        ref={ref}
      >
        <path d="m173.22 65.92 2.58 101.46h20.71l2.59-138.33zM199.09 17.7H167.48L117.96 88.49l15.79 22.55zM62.69 167.49h31.61l15.81-22.55-15.81-22.57zM62.69 65.92l71.15 101.46h31.61L94.29 65.92z" />
      </svg>
    );
  },
);

XAIIcon.displayName = "XAIIcon";

export default XAIIcon;
