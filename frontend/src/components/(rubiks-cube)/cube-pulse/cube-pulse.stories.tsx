import type { Meta, StoryObj } from "@storybook/tanstack-react";
import CubePulse from "./cube-pulse";

const meta: Meta<typeof CubePulse> = {
  title: "Components/(rubiks-cube)/Cube Pulse",
  component: CubePulse,
  parameters: {
    layout: "centered",
  },
  argTypes: {
    size: {
      control: "select",
      options: ["sm", "md", "lg"],
    },
    animationSpeed: {
      control: "select",
      options: ["slow", "medium", "fast"],
    },
  },
};

export default meta;
type Story = StoryObj<typeof CubePulse>;

export const Default: Story = {
  args: {
    size: "md",
    animationSpeed: "medium",
  },
};

export const Small: Story = {
  args: {
    size: "sm",
    animationSpeed: "medium",
  },
};

export const Large: Story = {
  args: {
    size: "lg",
    animationSpeed: "medium",
  },
};

export const FastPulse: Story = {
  args: {
    size: "md",
    animationSpeed: "fast",
  },
};

export const SlowPulse: Story = {
  args: {
    size: "md",
    animationSpeed: "slow",
  },
};
