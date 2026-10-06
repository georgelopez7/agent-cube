import type { Meta, StoryObj } from "@storybook/tanstack-react";

import Spacer from "./spacer";

const meta: Meta<typeof Spacer> = {
  title: "Components/(layouts)/Spacer",
  component: Spacer,
  parameters: {
    layout: "centered",
  },
  argTypes: {
    size: {
      control: "select",
      options: ["xs", "sm", "md", "lg", "xl"],
    },
  },
};

export default meta;
type Story = StoryObj<typeof Spacer>;

export const Default: Story = {
  args: {
    size: "md",
  },
  render: (args) => (
    <div className="w-64 border border-border">
      <div className="h-8 bg-primary/20" />
      <Spacer {...args} />
      <div className="h-8 bg-primary/20" />
    </div>
  ),
};

export const ExtraSmall: Story = {
  args: {
    size: "xs",
  },
  render: (args) => (
    <div className="w-64 border border-border">
      <div className="h-8 bg-primary/20" />
      <Spacer {...args} />
      <div className="h-8 bg-primary/20" />
    </div>
  ),
};

export const Small: Story = {
  args: {
    size: "sm",
  },
  render: (args) => (
    <div className="w-64 border border-border">
      <div className="h-8 bg-primary/20" />
      <Spacer {...args} />
      <div className="h-8 bg-primary/20" />
    </div>
  ),
};

export const Large: Story = {
  args: {
    size: "lg",
  },
  render: (args) => (
    <div className="w-64 border border-border">
      <div className="h-8 bg-primary/20" />
      <Spacer {...args} />
      <div className="h-8 bg-primary/20" />
    </div>
  ),
};

export const ExtraLarge: Story = {
  args: {
    size: "xl",
  },
  render: (args) => (
    <div className="w-64 border border-border">
      <div className="h-8 bg-primary/20" />
      <Spacer {...args} />
      <div className="h-8 bg-primary/20" />
    </div>
  ),
};
