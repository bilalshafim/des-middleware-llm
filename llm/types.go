package llm

import (
	"fmt"
	"time"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type RequestPayload struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
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

type Session struct {
	LastActivity time.Time
	History      []Message
}

func (m Message) String() string {
	return fmt.Sprintf("{Role: %s, Content: %s}", m.Role, m.Content)
}
