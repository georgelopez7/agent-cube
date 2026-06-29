import type { Meta, StoryObj } from "@storybook/tanstack-react";
import { AIProvider } from "#/domain/ai";
import { RubiksCubeStatus } from "#/domain/rubiks-cube";
import RubiksCubeCard from "./rubiks-cube-card";

const meta: Meta<typeof RubiksCubeCard> = {
  title: "Components/(cards)/Rubiks Cube Card",
  component: RubiksCubeCard,
  parameters: {
    layout: "centered",
  },
  args: {
    onBegin: (id) => console.log(id),
  },
};

export default meta;
type Story = StoryObj<typeof RubiksCubeCard>;

const baseCube = {
  id: "cube-1",
  llm: {
    provider: AIProvider.OpenAI,
    model: "gpt-4o",
  },
  cube: {
    centers: {},
    corners: {},
    edges: {},
    rotations: [
      {
        rotation: "R",
        index: 0,
        from_scramble: true,
        created_at: "2024-01-01T00:00:00Z",
      },
      {
        rotation: "U",
        index: 1,
        from_scramble: true,
        created_at: "2024-01-01T00:00:01Z",
      },
      {
        rotation: "R'",
        index: 2,
        from_scramble: true,
        created_at: "2024-01-01T00:00:02Z",
      },
      {
        rotation: "U'",
        index: 3,
        from_scramble: true,
        created_at: "2024-01-01T00:00:03Z",
      },
    ],
  },
  created_at: "2024-01-01T12:34:56Z",
  updated_at: "2024-01-01T12:34:56Z",
};

const withRotations = (rotations: typeof baseCube.cube.rotations) => ({
  ...baseCube,
  cube: { ...baseCube.cube, rotations },
});

export const Default: Story = {
  args: {
    cube: {
      ...withRotations(baseCube.cube.rotations),
      status: RubiksCubeStatus.Completed,
    },
  },
};

export const Created: Story = {
  args: {
    cube: {
      ...withRotations([]),
      status: RubiksCubeStatus.Created,
    },
  },
};

export const InProgress: Story = {
  args: {
    cube: {
      ...withRotations(baseCube.cube.rotations.slice(0, 2)),
      status: RubiksCubeStatus.InProgress,
    },
  },
};

export const Completed: Story = {
  args: {
    cube: {
      ...withRotations(baseCube.cube.rotations),
      status: RubiksCubeStatus.Completed,
    },
  },
};
