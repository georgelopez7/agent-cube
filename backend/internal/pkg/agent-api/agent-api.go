package agentapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type AgentAPI struct {
	URL string `json:"url"`
}

func NewAgentAPI(url string) AgentAPI {
	return AgentAPI{
		URL: url,
	}
}

// InvokeAgent - invokes the Rubik's Cube agent for the given cube ID and model.
func (a AgentAPI) InvokeAgent(id string, model string) (string, error) {
	url := fmt.Sprintf("%s/api/v1/rubiks-cubes/%s/agents/invoke", a.URL, id)

	type InvokeAgentPayload struct {
		Model string `json:"model"`
	}

	body := InvokeAgentPayload{
		Model: model,
	}

	payload, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("marshal request: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(payload))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to invoke agent: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	type InvokeAgentResponse struct {
		Message string `json:"message"`
	}

	var result InvokeAgentResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	return result.Message, nil
}
