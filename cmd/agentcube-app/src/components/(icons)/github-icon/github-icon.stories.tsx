import type { Meta, StoryObj } from "@storybook/react";

import GitHubIcon from "./github-icon";

const meta: Meta<typeof GitHubIcon> = {
  title: "Components/Icons/GitHub Icon",
  component: GitHubIcon,
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
