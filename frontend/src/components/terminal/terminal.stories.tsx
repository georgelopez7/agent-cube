import type { Meta, StoryObj } from "@storybook/tanstack-react";
import { Terminal } from "./terminal";

const meta: Meta<typeof Terminal> = {
  title: "Components/Terminal",
  component: Terminal,
  parameters: {
    layout: "centered",
  },
};

export default meta;
type Story = StoryObj<typeof Terminal>;

const logs = [
  "[AGENT] Initializing cube state analysis... [AGENT] Initializing cube state analysis... [AGENT] Initializing cube state analysis...",
  "[VISION] Scanning face colors: Front=RWBGRY",
  "[REASON] Current scramble depth: ~18 moves",
  "[PLAN] Attempting cross on white face...",
  "[MOVE] Executing: R U R' U'",
  "[EVAL] Cross progress: 2/4 edges placed",
  "[PLAN] Searching for edge piece WB...",
  "[MOVE] Executing: F' U' F",
  "[EVAL] Cross progress: 3/4 edges placed",
  "[REASON] Analyzing F2L opportunities...",
];

export const Default: Story = {
  args: {
    className: "w-100",
    logs,
  },
};

export const ScrollToBottom: Story = {
  args: {
    className: "w-100",
    logs: Array.from(
      { length: 50 },
      (_, i) => `[LOG] Line ${i + 1}: ${logs[i % logs.length]}`,
    ),
  },
};
