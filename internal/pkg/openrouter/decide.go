package openrouter

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	JevModelDefault  = "typesafe/jev-1.13"
	DecisionsURL     = "https://openrouter.ai/api/alpha/decisions"
	DecisionsTimeout = 60 * time.Second // HTTP timeout for Decisions API calls.
)

type DecisionsRequest struct {
	Model     string         `json:"model"`
	State     any            `json:"state"`
	Questions map[string]any `json:"questions"`
}

type ChoiceQuestion struct {
	Type         string            `json:"type"`
	Instructions string            `json:"instructions"`
	Criteria     map[string]string `json:"criteria"`
}

type DecisionAnswer struct {
	Type          string             `json:"type"`
	Choice        string             `json:"choice,omitempty"`
	Confidence    float64            `json:"confidence,omitempty"`
	Probabilities map[string]float64 `json:"probabilities,omitempty"`
	Noul          float64            `json:"noul,omitempty"`
	Score         float64            `json:"score,omitempty"`
}

type DecisionsUsage struct {
	InputTokens  int     `json:"input_tokens"`
	OutputTokens int     `json:"output_tokens"`
	Cost         float64 `json:"cost"`
}

type DecisionsResponse struct {
	ID       string                    `json:"id"`
	Model    string                    `json:"model"`
	Provider string                    `json:"provider"`
	Answers  map[string]DecisionAnswer `json:"answers"`
	Usage    DecisionsUsage            `json:"usage"`
}

// Decide - sends state + questions to the Decisions API
func (o *OpenRouter) Decide(ctx context.Context, model string, state any, questions map[string]any) (*DecisionsResponse, error) {
	if o.APIKey == "" {
		return nil, fmt.Errorf("openrouter: missing API key for Decisions call")
	}

	if model == "" {
		model = JevModelDefault
	}

	body, err := json.Marshal(DecisionsRequest{
		Model:     model,
		State:     state,
		Questions: questions,
	})

	if err != nil {
		return nil, fmt.Errorf("openrouter: failed to marshal decisions request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, DecisionsURL, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("openrouter: failed to build decisions request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+o.APIKey)
	req.Header.Set("Content-Type", "application/json")

	if o.HTTPReferer != "" {
		req.Header.Set("HTTP-Referer", o.HTTPReferer)
	}

	if o.XTitle != "" {
		req.Header.Set("X-Title", o.XTitle)
	}

	client := &http.Client{Timeout: DecisionsTimeout}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("openrouter: decisions request failed: %w", err)
	}

	defer resp.Body.Close()

	raw, err := io.ReadAll(io.LimitReader(resp.Body, 4<<20))
	if err != nil {
		return nil, fmt.Errorf("openrouter: failed to read decisions response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("openrouter: decisions request failed with status %d: %s", resp.StatusCode, string(raw))
	}

	var out DecisionsResponse
	if err := json.Unmarshal(raw, &out); err != nil {
		return nil, fmt.Errorf("openrouter: failed to decode decisions response: %w", err)
	}

	return &out, nil
}
