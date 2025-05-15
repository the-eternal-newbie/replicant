package replicant

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type OllamaResponse struct {
	Response string `json:"response"`
}

func QueryLLM(prompt string) (string, error) {
	payload := OllamaRequest{
		Model:  "mistral", // or any model loaded in Ollama
		Prompt: prompt,
		Stream: false,
	}

	data, _ := json.Marshal(payload)
	resp, err := http.Post("http://localhost:11434/api/generate", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result OllamaResponse
	json.NewDecoder(resp.Body).Decode(&result)
	return result.Response, nil
}
