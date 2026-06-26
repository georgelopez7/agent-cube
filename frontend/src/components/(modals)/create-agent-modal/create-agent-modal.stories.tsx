import type { Meta, StoryObj } from "@storybook/tanstack-react";
import CreateAgentModal from "./create-agent-modal";

const meta: Meta<typeof CreateAgentModal> = {
  title: "Components/(modals)/Create Agent Modal",
  component: CreateAgentModal,
  parameters: {
    layout: "centered",
  },
};

export default meta;
type Story = StoryObj<typeof CreateAgentModal>;

export const Default: Story = {};
