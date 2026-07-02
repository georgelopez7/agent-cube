import type { Meta, StoryObj } from "@storybook/react";

import CubeLogo from "./cube-logo";

const meta: Meta<typeof CubeLogo> = {
  title: "Components/Icons/Cube Logo",
  component: CubeLogo,
  parameters: {
    layout: "centered",
  },
};

export default meta;

type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    size: 64,
  },
};
