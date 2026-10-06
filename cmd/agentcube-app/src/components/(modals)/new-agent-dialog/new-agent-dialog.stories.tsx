import type { Meta, StoryObj } from "@storybook/tanstack-react";
import { AIProvider, type LLM } from "@/domain/ai";
import NewAgentDialog, { type NewAgentFormData } from "./new-agent-dialog";

const SAMPLE_LLMS: LLM[] = [
  { provider: AIProvider.OpenAI, model: "openai/gpt-4o" },
  { provider: AIProvider.Anthropic, model: "anthropic/claude-3-5-sonnet" },
  { provider: AIProvider.Google, model: "google/gemini-1.5-pro" },
];

const meta: Meta<typeof NewAgentDialog> = {
  title: "Components/(modals)/New Agent Dialog",
  component: NewAgentDialog,
  parameters: {
    layout: "centered",
  },
};

export default meta;
type Story = StoryObj<typeof NewAgentDialog>;

export const Default: Story = {
  render: () => {
    const handleSubmit = (data: NewAgentFormData) => {
      console.log(data);
    };

    return <NewAgentDialog llms={SAMPLE_LLMS} onSubmit={handleSubmit} />;
  },
};
