import type { LLM } from "./ai";

export type RubiksCube = {
  id: string;
  llm: LLM;
  status: RubiksCubeStatus;
  cube: Cube;
  created_at: string;
  updated_at: string;
};

export type Cube = {
  centers: Record<string, Cubie>;
  corners: Record<string, Cubie>;
  edges: Record<string, Cubie>;
  rotations: CubeRotation[] | null;
};

export type Rotation =
  | "F"
  | "F'"
  | "B"
  | "B'"
  | "U"
  | "U'"
  | "D"
  | "D'"
  | "L"
  | "L'"
  | "R"
  | "R'";

export type Cubie = {
  stickers: Record<string, string>;
};

export type CubeRotation = {
  rotation: Rotation;
  index: number;
  from_scramble: boolean;
  created_at: string;
};

export enum RubiksCubeStatus {
  Created = "created",
  InProgress = "in_progress",
  Completed = "completed",
  Stopped = "stopped",
}

// generateAlgorithm - generates a string representation of the cubes rotations
export const generateAlgorithm = (rotations: CubeRotation[]): string => {
  return rotations.map((rotation) => rotation.rotation).join(" ");
};
