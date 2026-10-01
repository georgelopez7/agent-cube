import { create } from "zustand";
import type { RubiksCubeRef } from "#/components/(rubiks-cube)/rubiks-cube/rubiks-cube";
import type {
  CubeRotation,
  RubiksCube,
  RubiksCubeStatus,
  TokenUsage,
} from "#/domain/rubiks-cube";

export type _RubiksCube = {
  ref: RubiksCubeRef;
  cube: RubiksCube;
};

export type RubikCubeStore = {
  records: Map<string, _RubiksCube>;
  reasoning: Map<string, string>;
  AddRecord: (record: _RubiksCube) => void;
  RemoveRecord: (id: string) => void;
  AppendReasoning: (id: string, text: string) => void;
  ClearReasoning: (id: string) => void;
  SetStatus: (id: string, status: RubiksCubeStatus) => void;
  SetUsage: (id: string, usage: TokenUsage) => void;
  SetInvokedAt: (id: string, invokedAt: string) => void;
  AddRotation: (id: string, rotation: CubeRotation) => void;
};

export const useRubiksCubeStore = create<RubikCubeStore>((set, get) => ({
  records: new Map(),
  reasoning: new Map(),
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
  AppendReasoning: (id, text) =>
    set((state) => {
      const next = new Map(state.reasoning);
      next.set(id, `${next.get(id) ?? ""}${text}`);
      return { reasoning: next };
    }),
  ClearReasoning: (id) =>
    set((state) => {
      if (!state.reasoning.has(id)) return state;

      const next = new Map(state.reasoning);
      next.delete(id);
      return { reasoning: next };
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
  SetUsage: (id, usage) => {
    const record = get().records.get(id);
    if (!record) return;

    set((state) => {
      const next = new Map(state.records);
      next.set(id, {
        ...record,
        cube: { ...record.cube, usage },
      });
      return { records: next };
    });
  },
  SetInvokedAt: (id, invokedAt) => {
    const record = get().records.get(id);
    if (!record) return;

    set((state) => {
      const next = new Map(state.records);
      next.set(id, {
        ...record,
        cube: { ...record.cube, invoked_at: invokedAt },
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
