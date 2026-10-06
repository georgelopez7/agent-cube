import React, { forwardRef } from "react";

const TypesafeIcon = forwardRef<SVGSVGElement, React.HTMLProps<SVGSVGElement>>(
  (props, ref) => (
    <svg
      viewBox="0 0 400 400"
      xmlns="http://www.w3.org/2000/svg"
      ref={ref}
      {...props}
    >
      <title>TypeSafe</title>
      <circle cx="200" cy="200" r="200" fill="#E551BA" />
      <path
        d="M199,50 L97,116 L97,230 L143,260 L143,312 L200,349 L302,283 L302,168 L256,138 L256,87 Z M274,277 L200,326 L171,308 L245,259 Z M153,205 L181,224 L152,243 L124,224 Z M257,162 L283,179 L282,259 L256,242 Z M236,161 L237,241 L163,289 L162,260 L209,229 L209,179 Z M163,161 L190,179 L189,206 L162,188 Z M198,125 L228,143 L200,162 L171,143 Z M210,79 L237,97 L236,125 L209,107 Z M189,79 L190,107 L143,138 L143,189 L117,206 L116,127 Z"
        fill="#171717"
        fillRule="evenodd"
      />
    </svg>
  ),
);

TypesafeIcon.displayName = "TypesafeIcon";

export default TypesafeIcon;
