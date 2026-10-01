import { AIProvider } from "@/domain/ai";
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

type IconComponent = React.ComponentType<{ className?: string }>;

const providerIcons: Record<AIProvider, IconComponent> = {
  [AIProvider.OpenAI]: OpenAIIcon,
  [AIProvider.Anthropic]: AnthropicIcon,
  [AIProvider.Mistral]: MistralIcon,
  [AIProvider.Google]: GoogleIcon,
  [AIProvider.Meta]: MetaIcon,
  [AIProvider.Deepseek]: DeepseekIcon,
  [AIProvider.XAI]: XAIIcon,
  [AIProvider.Qwen]: QwenIcon,
  [AIProvider.OpenRouter]: OpenRouterIcon,
  [AIProvider.Minimax]: MinimaxLogo,
  [AIProvider.ZAI]: ZAIIcon,
  [AIProvider.Kimi]: KimiIcon,
  [AIProvider.Acree]: AcreeIcon,
  [AIProvider.Moonshot]: MoonshotIcon,
  [AIProvider.Typesafe]: TypesafeIcon,
};

export function getAIProviderIcon(provider: AIProvider): IconComponent {
  return providerIcons[provider];
}
