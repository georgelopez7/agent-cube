package openrouter

import (
	"net/http"

	"github.com/achetronic/adk-utils-go/genai/openai"
)

type OpenRouter struct {
	BaseURL     string
	APIKey      string
	HTTPReferer string
	XTitle      string
}

func NewOpenRouter(baseURL string, apiKey string, httpReferrer string, xTitle string) *OpenRouter {
	return &OpenRouter{
		BaseURL:     baseURL,
		APIKey:      apiKey,
		HTTPReferer: httpReferrer,
		XTitle:      xTitle,
	}
}

func (o *OpenRouter) NewModel(model string) *openai.Model {
	return openai.New(openai.Config{
		APIKey:    o.APIKey,
		BaseURL:   o.BaseURL,
		ModelName: model,
		HTTPOptions: openai.HTTPOptions{
			Headers: http.Header{
				"HTTP-Referer": []string{o.HTTPReferer},
				"X-Title":      []string{o.XTitle},
			},
		},
	})
}
