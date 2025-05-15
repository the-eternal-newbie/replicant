package replicant

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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
	host := os.Getenv("OLLAMA_HOST")
	if host == "" {
		host = "http://localhost:11434"
	}

	model := os.Getenv("OLLAMA_MODEL")
	if model == "" {
		model = "mistral"
	}

	payload := OllamaRequest{
		Model:  model,
		Prompt: prompt,
		Stream: false,
	}

	data, _ := json.Marshal(payload)
	resp, err := http.Post(fmt.Sprintf("%s/api/generate", host), "application/json", bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result OllamaResponse
	json.NewDecoder(resp.Body).Decode(&result)
	return result.Response, nil
}
