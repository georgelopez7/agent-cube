import type { Meta, StoryObj } from "@storybook/tanstack-react";
import { AIProvider, type LLM } from "@/domain/ai";
import CreateAgentModal, {
  type CreateAgentFormData,
} from "./create-agent-modal";

const SAMPLE_LLMS: LLM[] = [
  { provider: AIProvider.OpenAI, model: "openai/gpt-4o" },
  { provider: AIProvider.Anthropic, model: "anthropic/claude-3-5-sonnet" },
  { provider: AIProvider.Google, model: "google/gemini-1.5-pro" },
];

const meta: Meta<typeof CreateAgentModal> = {
  title: "Components/(modals)/Create Agent Modal",
  component: CreateAgentModal,
  parameters: {
    layout: "centered",
  },
};

export default meta;
type Story = StoryObj<typeof CreateAgentModal>;

export const Default: Story = {
  render: () => {
    const handleSubmit = (data: CreateAgentFormData) => {
      console.log(data);
    };

    return <CreateAgentModal llms={SAMPLE_LLMS} onSubmit={handleSubmit} />;
  },
};
