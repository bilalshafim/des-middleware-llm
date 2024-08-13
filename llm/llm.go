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

func CallLocalLLM(sessionID, message *string) *string {
	// Implement your logic to interact with the local LLM
	llmResponse, err := SendAPIRequest(sessionID, message)
	// log.Println("Response_MemAddr: ", llmResponse)
	if err != nil {
		log.Println("Error: ", err)
		if llmResponse != nil {
			return llmResponse
		}
		return new(string)
	}
	return llmResponse
}

func SendAPIRequest(sessionID, message *string) (*string, error) {
	// Create request payload
	// history := GetSessionHistory(*sessionID)

	// user message not being added here to session history
	var history *[]Message
	history = UpdateSessionHistory(*sessionID, "user", *message)
	log.Print(*history)
	/*
		history = append(history, Message{
			Role:    "user",
			Content: *message,
		})
	*/

	// messages := []Message{
	// 	{
	// 		Role:    "user",
	// 		Content: *message,
	// 	},
	// 	{
	// 		Role:    "user",
	// 		Content: "You are an advanced AI neural system for a NLU Chatbot. The NLU and you are different component of the same brain. Your job is to provide short answers for what the NLU is not capable of. Limit your response to one sentences. ",
	// 	},
	// }
	payload := RequestPayload{
		Model:    "llama3:8b-instruct-q5_0",
		Messages: *history,
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
	log.Printf("Response: %+v", response.Message.Content)

	// LLM response not being added here to session history
	if response.Message.Content != "" {
		history = UpdateSessionHistory(*sessionID, "assistant", response.Message.Content)
	}

	log.Print(*history)
	// Print parsed response
	return &response.Message.Content, nil
}
