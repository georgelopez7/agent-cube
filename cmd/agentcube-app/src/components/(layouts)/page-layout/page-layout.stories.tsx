import type { Meta, StoryObj } from "@storybook/tanstack-react";

import { PageLayout } from "./page-layout";

const meta: Meta<typeof PageLayout> = {
  title: "Components/(layouts)/Page Layout",
  component: PageLayout,
  parameters: {
    layout: "fullscreen",
  },
  argTypes: {
    navbar: {
      control: "boolean",
    },
  },
};

export default meta;
type Story = StoryObj<typeof PageLayout>;

export const Default: Story = {
  args: {
    navbar: true,
  },
  render: (args) => (
    <PageLayout {...args}>
      <div className="text-center">
        <h1 className="text-4xl font-bold">Page Content</h1>
        <p className="text-muted-foreground">Centered inside the page layout</p>
      </div>
    </PageLayout>
  ),
};

export const WithoutNavbar: Story = {
  args: {
    navbar: false,
  },
  render: (args) => (
    <PageLayout {...args}>
      <div className="text-center">
        <h1 className="text-4xl font-bold">Page Content</h1>
        <p className="text-muted-foreground">Without the navbar</p>
      </div>
    </PageLayout>
  ),
};
