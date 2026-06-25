export enum AIProvider {
  OpenAI = "openai",
  Anthropic = "anthropic",
  Mistral = "mistral",
  Google = "google",
  Meta = "meta",
  Deepseek = "deepseek",
  XAI = "xai",
  Qwen = "qwen",
  OpenRouter = "openrouter",
  Minimax = "minimax",
  ZAI = "zai",
  Kimi = "kimi",
  Acree = "acree",
  Moonshot = "moonshot",
}

export type LLM = {
  provider: AIProvider;
  model: string;
};
