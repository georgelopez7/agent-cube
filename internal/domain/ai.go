package domain

type LLM struct {
	Provider string `json:"provider" bson:"provider"`
	Model    string `json:"model" bson:"model"`
}

func NewLLM(provider string, model string) LLM {
	return LLM{
		Provider: provider,
		Model:    model,
	}
}

var LLMs = []LLM{
	NewLLM("openai", "openai/gpt-5.4-mini"),
	NewLLM("anthropic", "anthropic/claude-haiku-4.5"),
	NewLLM("google", "google/gemini-3.1-flash-lite"),
	NewLLM("moonshot", "moonshotai/kimi-k2.7-code"),
	NewLLM("x-ai", "x-ai/grok-4.3"),
	NewLLM("deepseek", "deepseek/deepseek-v4-pro"),
}
