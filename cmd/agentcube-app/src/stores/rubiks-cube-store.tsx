import { create } from "zustand";
import type { RubiksCubeRef } from "#/components/(rubiks-cube)/rubiks-cube/rubiks-cube";
import type {
  CubeRotation,
  RubiksCube,
  RubiksCubeStatus,
} from "#/domain/rubiks-cube";

export type _RubiksCube = {
  ref: RubiksCubeRef;
  cube: RubiksCube;
};

export type RubikCubeStore = {
  records: Map<string, _RubiksCube>;
  AddRecord: (record: _RubiksCube) => void;
  RemoveRecord: (id: string) => void;
  SetStatus: (id: string, status: RubiksCubeStatus) => void;
  AddRotation: (id: string, rotation: CubeRotation) => void;
};

export const useRubiksCubeStore = create<RubikCubeStore>((set, get) => ({
  records: new Map(),
  AddRecord: (record) =>
    set((state) => {
      const next = new Map(state.records);
      next.set(record.cube.id, record);
      return { records: next };
    }),
  RemoveRecord: (id) =>
    set((state) => {
      if (!state.records.has(id)) return state;

      const next = new Map(state.records);
      next.delete(id);
      return { records: next };
    }),
  SetStatus: (id, status) => {
    const record = get().records.get(id);
    if (!record) return;

    set((state) => {
      const next = new Map(state.records);
      next.set(id, {
        ...record,
        cube: { ...record.cube, status },
      });
      return { records: next };
    });
  },
  AddRotation: (id, rotation) => {
    const record = get().records.get(id);
    if (!record) return;

    set((state) => {
      const next = new Map(state.records);
      next.set(id, {
        ...record,
        cube: {
          ...record.cube,
          cube: {
            ...record.cube.cube,
            rotations: [...(record.cube.cube.rotations ?? []), rotation],
          },
        },
      });
      return { records: next };
    });
  },
}));
