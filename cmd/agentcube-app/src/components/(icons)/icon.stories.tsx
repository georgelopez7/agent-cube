import type { Meta, StoryObj } from "@storybook/tanstack-react";
import AcreeIcon from "./acree-icon";
import AnthropicIcon from "./anthropic-icon";
import DeepseekIcon from "./deepseek-icon";
import GoogleIcon from "./google-icon";
import KimiIcon from "./kimi-icon";
import MetaIcon from "./meta-icon";
import MinimaxLogo from "./minimax-logo";
import MistralIcon from "./mistral-icon";
import MoonshotIcon from "./moonshot-icon";
import OpenAIIcon from "./openai-icon";
import OpenRouterIcon from "./openrouter-icon";
import QwenIcon from "./qwen-icon";
import TypesafeIcon from "./typesafe-icon";
import XAIIcon from "./xai-icon";
import ZAIIcon from "./zai-icon";

const meta: Meta = {
  title: "Components/(icons)/Icons",
  parameters: {
    layout: "centered",
  },
};

export default meta;
type Story = StoryObj<typeof meta>;

export const OpenAI: Story = {
  render: () => <OpenAIIcon className="w-16 h-16" />,
};

export const Anthropic: Story = {
  render: () => <AnthropicIcon className="w-16 h-16" />,
};

export const Mistral: Story = {
  render: () => <MistralIcon className="w-16 h-16" />,
};

export const Google: Story = {
  render: () => <GoogleIcon className="w-16 h-16" />,
};

export const _Meta: Story = {
  render: () => <MetaIcon className="w-16 h-16" />,
};

export const Deepseek: Story = {
  render: () => <DeepseekIcon className="w-16 h-16" />,
};

export const XAI: Story = {
  render: () => <XAIIcon className="w-16 h-16" />,
};

export const Qwen: Story = {
  render: () => <QwenIcon className="w-16 h-16" />,
};

export const OpenRouter: Story = {
  render: () => <OpenRouterIcon className="w-16 h-16" />,
};

export const Minimax: Story = {
  render: () => <MinimaxLogo className="w-16 h-16" />,
};

export const ZAI: Story = {
  render: () => <ZAIIcon className="w-16 h-16" />,
};

export const Kimi: Story = {
  render: () => <KimiIcon className="w-16 h-16" />,
};

export const Acree: Story = {
  render: () => <AcreeIcon className="w-16 h-16" />,
};

export const Moonshot: Story = {
  render: () => <MoonshotIcon className="w-16 h-16" />,
};

export const Typesafe: Story = {
  render: () => <TypesafeIcon className="w-16 h-16" />,
};
