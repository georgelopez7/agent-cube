import type { Meta, StoryObj } from "@storybook/tanstack-react";

import Navbar from "./navbar";

const meta: Meta<typeof Navbar> = {
  title: "Components/(layouts)/Navbar",
  component: Navbar,
  parameters: {
    layout: "centered",
  },
};

export default meta;
type Story = StoryObj<typeof Navbar>;

export const Default: Story = {
  render: () => (
    <div className="w-[80vw]">
      <Navbar />
    </div>
  ),
};
