export enum AIProvider {
  OpenAI = "openai",
  Anthropic = "anthropic",
  Mistral = "mistral",
  Google = "google",
  Meta = "meta",
  Deepseek = "deepseek",
  XAI = "x-ai",
  Qwen = "qwen",
  OpenRouter = "openrouter",
  Minimax = "minimax",
  ZAI = "zai",
  Kimi = "kimi",
  Acree = "acree",
  Moonshot = "moonshot",
  Typesafe = "typesafe",
}

export type LLM = {
  provider: AIProvider;
  model: string;
};
