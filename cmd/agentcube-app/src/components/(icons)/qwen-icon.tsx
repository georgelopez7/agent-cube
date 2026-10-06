"use client";

import React, { forwardRef } from "react";

const QwenIcon = forwardRef<SVGSVGElement, React.HTMLProps<SVGSVGElement>>(
  (props, ref) => {
    return (
      <svg
        width="1024"
        height="1012"
        viewBox="0 0 1024 1012"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
        {...props}
        ref={ref}
      >
        <path
          d="M1023 634.864L887.966 400.37L959.234 280.517C964.93 270.512 964.93 258.145 959.234 248.14L887.966 128.287C884.354 121.825 877.477 117.795 869.906 117.795H606.645L545.936 11.4914C542.324 5.02983 535.447 1 527.876 1H387.285C379.713 1 372.837 5.02983 369.225 11.4914L234.191 245.916H93.5999C86.0285 245.916 79.1518 249.946 75.5398 256.408L4.27191 376.26C-1.42397 386.265 -1.42397 398.633 4.27191 408.638L124.858 626.179L64.148 732.483C58.4522 742.488 58.4522 754.855 64.148 764.861L135.416 884.713C139.028 891.175 145.905 895.205 153.476 895.205H416.737L477.446 1001.51C481.058 1007.97 487.935 1012 495.506 1012H636.097C643.669 1012 650.545 1007.97 654.157 1001.51L789.191 767.084H909.777C917.348 767.084 924.225 763.054 927.837 756.592L999.105 636.74C1004.8 626.735 1004.8 614.367 999.105 604.362L1023 634.864Z"
          fill="url(#paint0_radial_6_2)"
        />
        <path
          d="M635.957 1011.61H495.407L416.659 884.368H153.474L234.165 757.124H369.161L75.5595 263.289H234.165L387.216 11.4408L458.463 138.685L387.216 263.289H930.809L859.561 383.101L994.556 617.447H859.561L789.008 492.842L511.378 1011.61H635.957Z"
          fill="white"
        />
        <path
          d="M696.65 433.735H337.495L511.378 727.396L696.65 433.735Z"
          fill="url(#paint1_radial_6_2)"
        />
        <defs>
          <radialGradient
            id="paint0_radial_6_2"
            cx="0"
            cy="0"
            r="1"
            gradientUnits="userSpaceOnUse"
            gradientTransform="translate(503.286 574.069) rotate(90) scale(694.798 694.619)"
          >
            <stop stopColor="#665CEE" />
            <stop offset="1" stopColor="#332E91" />
          </radialGradient>
          <radialGradient
            id="paint1_radial_6_2"
            cx="0"
            cy="0"
            r="1"
            gradientUnits="userSpaceOnUse"
            gradientTransform="translate(503.184 573.828) rotate(90) scale(694.563 694.421)"
          >
            <stop stopColor="#665CEE" />
            <stop offset="1" stopColor="#332E91" />
          </radialGradient>
        </defs>
      </svg>
    );
  },
);

QwenIcon.displayName = "QwenIcon";

export default QwenIcon;
