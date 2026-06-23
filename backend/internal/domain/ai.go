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
