package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	apiURL             = "https://gpt.sat.ae:10444/ollama/api/chat"
	authorizationToken = "sk-d471af019699429b825f5013333dc27e"
)

type RequestPayload struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ResponsePayload struct {
	Model              string `json:"model"`
	CreatedAt          string `json:"created_at"`
	Message            Message
	DoneReason         string `json:"done_reason"`
	Done               bool   `json:"done"`
	TotalDuration      int64  `json:"total_duration"`
	LoadDuration       int64  `json:"load_duration"`
	PromptEvalDuration int64  `json:"prompt_eval_duration"`
	EvalCount          int    `json:"eval_count"`
	EvalDuration       int64  `json:"eval_duration"`
}

func SendAPIRequest(message *string) (*string, error) {
	// Create request payload
	messages := []Message{
		{
			Role:    "user",
			Content: *message,
		},
		{
			Role:    "user",
			Content: "You are an advanced AI neural system for a NLU Chatbot. The NLU and you are different component of the same brain. Your job is to provide short answers for what the NLU is not capable of. Limit your response to one sentences. ",
		},
	}
	payload := RequestPayload{
		Model:    "llama3:8b-instruct-q5_0",
		Messages: messages,
		Stream:   false,
	}

	// Marshal payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshalling payload: %w", err)
	}

	// Create request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+authorizationToken)

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	// log.Printf("Raw Response: %s", body) // Print raw response for debugging

	// Unmarshal response
	var response ResponsePayload
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	// Print parsed response
	log.Printf("Response: %+v", response.Message.Content)
	return &response.Message.Content, nil
}
