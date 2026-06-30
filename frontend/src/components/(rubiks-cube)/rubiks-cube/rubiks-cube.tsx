import cn from "cnfast";
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
import Spacer from "#/components/(layouts)/spacer/spacer";
import type { RubiksCube as RubiksCubeType } from "#/domain/rubiks-cube";
import { useRubiksCubeStore } from "#/stores/rubiks-cube-store";
import { RotationButtons } from "./rotation-buttons";
import { RotationCollapsible } from "./rotation-collapsible";

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

interface RubiksCubeHTMLElementProps extends HTMLAttributes<RubiksCubeHTMLElement> {
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
  cube?: RubiksCubeType;
  algorithm?: string;
  showRotationButtons?: boolean;
  showRotationsCollapsible?: boolean;
  showResetAlgoButton?: boolean;
  showResetCameraButton?: boolean;
  showBorder?: boolean;
}

const RubiksCube = forwardRef<RubiksCubeRef, IProps>(
  (
    {
      cube,
      algorithm = "",
      showRotationButtons = true,
      showRotationsCollapsible = true,
      showResetAlgoButton = true,
      showResetCameraButton = true,
      showBorder = true,
    },
    ref,
  ) => {
    const [loading, setLoading] = useState(true);

    // cubeRef - reference to the underlying <twisty-player> web component DOM element
    const cubeRef = useRef<RubiksCubeHTMLElement | null>(null);
    // apiRef - stable public API object exposed to parent refs and registered in the global store
    const apiRef = useRef<RubiksCubeRef>({
      // rotate - applies a single move to the cube
      rotate: async (move: string) => {
        const cubeElement = cubeRef.current;
        if (!cubeElement) return;

        if (cubeElement.experimentalAddMove) {
          await cubeElement.experimentalAddMove(move);
        } else {
          const currentAlg = cubeElement.alg || "";
          cubeElement.alg = currentAlg + (currentAlg ? " " : "") + move;
        }
      },
      // reset - clears the cube's current algorithm
      reset: async () => {
        const cubeElement = cubeRef.current;
        if (!cubeElement) return;

        if (cubeElement.experimentalSetAlg) {
          await cubeElement.experimentalSetAlg("");
        } else {
          cubeElement.alg = "";
        }
      },
      // resetCamera - resets the camera to its default latitude and longitude
      resetCamera: () => {
        const cubeElement = cubeRef.current;
        if (!cubeElement) return;

        cubeElement.setAttribute("camera-latitude", "35");
        cubeElement.setAttribute("camera-longitude", "30");
      },
      // getCube - exposes the underlying twisty-player element
      getCube: () => cubeRef.current,
    });

    // Expose the stable apiRef to parent components via the forwarded ref
    useImperativeHandle(ref, () => apiRef.current);

    // Store actions used to register and unregister this cube instance globally
    const AddRecord = useRubiksCubeStore((state) => state.AddRecord);
    const RemoveRecord = useRubiksCubeStore((state) => state.RemoveRecord);

    // xrotations - Subscribe to live rotations from the store so websocket updates are
    // reflected in the rotation list without waiting for a query refetch.
    const cubeID = cube?.id;
    const xrotations = useRubiksCubeStore(
      useCallback(
        (state) =>
          cubeID ? state.records.get(cubeID)?.cube.cube.rotations : undefined,
        [cubeID],
      ),
    );

    // Register this cube in the global store when a cube prop is provided, and
    // remove it when the component unmounts or the cube prop changes.
    useEffect(() => {
      if (!cube) return;

      AddRecord({ ref: apiRef.current, cube });

      return () => RemoveRecord(cube.id);
    }, [cube, AddRecord, RemoveRecord]);

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
      await apiRef.current.rotate(move);
    }, []);

    // handleReset - clears the cube's current algorithm
    const handleReset = useCallback(async () => {
      await apiRef.current.reset();
    }, []);

    // handleResetCamera - resets the camera to its default latitude and longitude
    const handleResetCamera = useCallback(() => {
      apiRef.current.resetCamera();
    }, []);

    return (
      <div className="flex flex-col items-center">
        <div
          className={cn(
            "relative bg-transparent border-2 rounded-sm shadow-sm",
            !showBorder && "border-transparent",
          )}
        >
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
        {showRotationsCollapsible && xrotations && xrotations.length > 0 && (
          <>
            <Spacer size="xs" />
            <RotationCollapsible rotations={xrotations} disabled={loading} />
          </>
        )}
        <div className="flex flex-col">
          {showRotationButtons && (
            <>
              <Spacer size="xs" />
              <RotationButtons
                rotations={ROTATIONS}
                onRotation={handleRotation}
                disabled={loading}
              />
            </>
          )}
          {(showResetAlgoButton || showResetCameraButton) && (
            <>
              <Spacer size="xs" />
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
            </>
          )}
        </div>
      </div>
    );
  },
);

RubiksCube.displayName = "RubiksCube";

export default RubiksCube;
