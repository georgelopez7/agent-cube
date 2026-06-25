import type { Meta, StoryObj } from "@storybook/tanstack-react";
import RubiksCube from "./rubiks-cube";

const meta: Meta<typeof RubiksCube> = {
  title: "Components/Rubiks Cube",
  component: RubiksCube,
  parameters: {
    layout: "centered",
  },
  argTypes: {
    algorithm: {
      control: "text",
    },
    showRotationButtons: {
      control: "boolean",
    },
    showResetAlgoButton: {
      control: "boolean",
    },
    showResetCameraButton: {
      control: "boolean",
    },
  },
};

export default meta;
type Story = StoryObj<typeof RubiksCube>;

export const Default: Story = {
  args: {
    algorithm: "",
    showRotationButtons: true,
    showResetAlgoButton: true,
    showResetCameraButton: true,
  },
};

export const WithAlgorithm: Story = {
  args: {
    algorithm: "R U R' U'",
    showRotationButtons: true,
    showResetAlgoButton: true,
    showResetCameraButton: true,
  },
};

export const HiddenRotationButtons: Story = {
  args: {
    algorithm: "",
    showRotationButtons: false,
    showResetAlgoButton: true,
    showResetCameraButton: true,
  },
};
