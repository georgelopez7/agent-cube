import type { Meta, StoryObj } from "@storybook/react";

import XLogo from "./x-logo";

const meta: Meta<typeof XLogo> = {
  title: "Components/Icons/X Logo",
  component: XLogo,
  parameters: {
    layout: "centered",
  },
};

export default meta;

type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    size: 24,
  },
};
