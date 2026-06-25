import { Loader2, Video } from "lucide-react";
import {
  createElement,
  forwardRef,
  type HTMLAttributes,
  useCallback,
  useEffect,
  useImperativeHandle,
  useRef,
  useState,
} from "react";
import { RotationButtons } from "./rotation-buttons";

const TWISTY_PLAYER_NAME = "twisty-player";
const TWISTY_PLAYER_SCRIPT = "https://cdn.cubing.net/v0/js/cubing/twisty";

const ROTATIONS = [
  { label: "F", value: "F" },
  { label: "R", value: "R" },
  { label: "U", value: "U" },
  { label: "L", value: "L" },
  { label: "B", value: "B" },
  { label: "D", value: "D" },
  { label: "F'", value: "F'" },
  { label: "R'", value: "R'" },
  { label: "U'", value: "U'" },
  { label: "L'", value: "L'" },
  { label: "B'", value: "B'" },
  { label: "D'", value: "D'" },
];

export interface RubiksCubeHTMLElement extends HTMLElement {
  alg: string;
  experimentalAddMove?: (move: string) => Promise<void>;
  experimentalSetAlg?: (alg: string) => Promise<void>;
}

interface RubiksCubeHTMLElementProps
  extends HTMLAttributes<RubiksCubeHTMLElement> {
  alg?: string;
  visualization?: string;
  background?: string;
  puzzle?: string;
  controlPanel?: string;
  hintFacelets?: string;
}

export interface RubiksCubeRef {
  rotate: (move: string) => Promise<void>;
  reset: () => Promise<void>;
  resetCamera: () => void;
  getCube: () => RubiksCubeHTMLElement | null;
}

interface IProps {
  algorithm?: string;
  showRotationButtons?: boolean;
  showResetAlgoButton?: boolean;
  showResetCameraButton?: boolean;
}

const RubiksCube = forwardRef<RubiksCubeRef, IProps>(
  (
    {
      algorithm = "",
      showRotationButtons = true,
      showResetAlgoButton = true,
      showResetCameraButton = true,
    },
    ref,
  ) => {
    const [loading, setLoading] = useState(true);

    const cubeRef = useRef<RubiksCubeHTMLElement | null>(null);

    useImperativeHandle(ref, () => ({
      // rotate - applies a single move to the cube
      rotate: async (move: string) => {
        if (!cubeRef.current) return;

        const cube = cubeRef.current;

        if (cube.experimentalAddMove) {
          await cube.experimentalAddMove(move);
        } else {
          const currentAlg = cube.alg || "";
          cube.alg = currentAlg + (currentAlg ? " " : "") + move;
        }
      },
      // reset - clears the cube's current algorithm
      reset: async () => {
        if (!cubeRef.current) return;

        const cube = cubeRef.current;

        if (cube.experimentalSetAlg) {
          await cube.experimentalSetAlg("");
        } else {
          cube.alg = "";
        }
      },
      // resetCamera - resets the camera to its default latitude and longitude
      resetCamera: () => {
        if (!cubeRef.current) return;

        const cube = cubeRef.current;
        cube.setAttribute("camera-latitude", "35");
        cube.setAttribute("camera-longitude", "30");
      },
      // getCube - exposes the underlying twisty-player element
      getCube: () => cubeRef.current,
    }));

    useEffect(() => {
      // initCube - loads the cubing/twisty web component script once on mount
      const initCube = async () => {
        if (customElements.get(TWISTY_PLAYER_NAME)) {
          setLoading(false);
          return;
        }

        const existingScript = document.querySelector(
          `script[src*="${TWISTY_PLAYER_SCRIPT}"]`,
        );

        if (existingScript) {
          const checkLoaded = setInterval(() => {
            if (customElements.get(TWISTY_PLAYER_NAME)) {
              clearInterval(checkLoaded);
              setLoading(false);
            }
          }, 50);

          return () => clearInterval(checkLoaded);
        }

        const script = document.createElement("script");
        script.src = TWISTY_PLAYER_SCRIPT;
        script.type = "module";

        script.onload = () => {
          setLoading(false);
        };

        document.head.appendChild(script);
      };

      initCube();
    }, []);

    // applyAlgorithm - applies the preset algorithm to the cube when it changes or the player loads
    useEffect(() => {
      if (loading || !cubeRef.current) return;

      const cube = cubeRef.current;

      if (cube.experimentalSetAlg) {
        void cube.experimentalSetAlg(algorithm);
      } else {
        cube.alg = algorithm;
      }
    }, [loading, algorithm]);

    // handleRotation - handles an applied rotation
    const handleRotation = useCallback(async (move: string) => {
      if (!cubeRef.current) return;

      const cube = cubeRef.current;

      if (cube.experimentalAddMove) {
        await cube.experimentalAddMove(move);
      } else {
        const currentAlg = cube.alg || "";
        cube.alg = currentAlg + (currentAlg ? " " : "") + move;
        await new Promise((resolve) => setTimeout(resolve, 500));
      }
    }, []);

    // handleReset - clears the cube's current algorithm
    const handleReset = useCallback(async () => {
      if (!cubeRef.current) return;

      const cube = cubeRef.current;

      if (cube.experimentalSetAlg) {
        await cube.experimentalSetAlg("");
      } else {
        cube.alg = "";
      }
    }, []);

    // handleResetCamera - resets the camera to its default latitude and longitude
    const handleResetCamera = useCallback(async () => {
      if (!cubeRef.current) return;

      const cube = cubeRef.current;
      cube.setAttribute("camera-latitude", "35");
      cube.setAttribute("camera-longitude", "30");
    }, []);

    return (
      <div className="flex flex-col items-center gap-4">
        <div className="p-2">
          <div className="relative bg-transparent border-2 rounded-sm shadow-sm mb-4">
            {!loading ? (
              createElement("twisty-player", {
                ref: cubeRef,
                alg: algorithm,
                visualization: "3D",
                background: "none",
                puzzle: "3x3x3",
                "control-panel": "none",
                "hint-facelets": "none",
                style: {
                  width: "315px",
                  height: "315px",
                },
              } as RubiksCubeHTMLElementProps)
            ) : (
              <div className="flex items-center justify-center w-78.75 h-78.75">
                <Loader2 className="w-8 h-8 animate-spin text-muted-foreground" />
              </div>
            )}
          </div>
          <div className="flex flex-col gap-2">
            {showRotationButtons && (
              <RotationButtons
                rotations={ROTATIONS}
                onRotation={handleRotation}
                disabled={loading}
              />
            )}
            <div className="flex justify-center gap-2">
              {showResetAlgoButton && (
                <button
                  type="button"
                  onClick={handleReset}
                  className="px-4 py-1 text-sm border-2 rounded hover:bg-accent cursor-pointer"
                  disabled={loading}
                >
                  Reset
                </button>
              )}
              {showResetCameraButton && (
                <button
                  type="button"
                  onClick={handleResetCamera}
                  className="flex items-center gap-2 px-4 py-1 text-sm border-2 rounded hover:bg-accent cursor-pointer"
                  disabled={loading}
                >
                  <Video className="size-4" />
                  Reset Camera
                </button>
              )}
            </div>
          </div>
        </div>
      </div>
    );
  },
);

RubiksCube.displayName = "RubiksCube";

export default RubiksCube;
