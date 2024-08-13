package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"webhook/es"
	"webhook/handler"
	"webhook/llm"
)

type SerializedMessages struct {
	Messages []llm.Message `json:"messages"`
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	var req es.Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	handler.PerformAction(&req.QueryResult.Action)
	handler.HandleIntent(&req.QueryResult.Intent.DisplayName)

	log.Printf("QueryText: %s", req.QueryResult.QueryText)
	// log.Print("QueryText_MemAddr: ", &req.QueryResult.QueryText)
	log.Printf("SessionID: %s", req.Session)

	start := time.Now()
	responseText := llm.CallLocalLLM(&req.Session, &req.QueryResult.QueryText)
	duration := time.Since(start)
	log.Printf("llm_proc_time: %v\n", duration)

	if responseText == nil {
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}

	response := es.Response{
		FulfillmentText:     "",
		FulfillmentMessages: []es.FulfillmentMessage{{Text: es.Text{Text: []string{*responseText}}}},
		Source:              "",
		Payload:             map[string]interface{}{},
		OutputContexts:      []es.Context{},
		FollowUpEventInput:  es.EventInput{},
		SessionEntityTypes:  []es.SessionEntityTypes{},
	}
	if *responseText == `` {
		response.FulfillmentMessages = nil
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode response: %v", err)
	}
	// json.NewEncoder(w).Encode(response)
}

func main() {
	go llm.CleanupExpiredSessions()

	http.HandleFunc("/inference", webhookHandler)
	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
