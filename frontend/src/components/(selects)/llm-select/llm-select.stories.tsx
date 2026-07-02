import type { Meta, StoryObj } from "@storybook/tanstack-react";
import { LLMSelect } from "./llm-select";
import { AIProvider, type LLM } from "@/domain/ai";

const meta: Meta<typeof LLMSelect> = {
  title: "Components/(selects)/LLM Select",
  component: LLMSelect,
  parameters: {
    layout: "centered",
  },
};

export default meta;
type Story = StoryObj<typeof LLMSelect>;

export const Default: Story = {
  render: () => {
    const LLMs = [
      { provider: AIProvider.Google, model: "google/gemini-3.1-flash-lite-preview" },
      { provider: AIProvider.OpenAI, model: "openai/gpt-3.5-turbo" },
    ];

    const handleChange = (llm: LLM) => {
      console.log(llm);
    };

    return <LLMSelect llms={LLMs} onChange={handleChange} />;
  },
};
