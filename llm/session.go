package llm

import (
	"log"
	"sync"
	"time"
)

const sessionTimeout = 30 * time.Minute

var (
	sessionHistories = sync.Map{}
	mu               sync.Mutex
)

func CleanupExpiredSessions() {
	for {
		time.Sleep(5 * time.Minute)

		now := time.Now()
		sessionHistories.Range(func(key, value interface{}) bool {
			session := value.(Session)
			if now.Sub(session.LastActivity) > sessionTimeout {
				sessionHistories.Delete(key)
			}
			return true
		})
	}
}

func UpdateSessionHistory(sessionID, role, content string) *[]Message {
	mu.Lock()
	defer mu.Unlock()

	session, ok := sessionHistories.Load(sessionID)

	sess := session.(Session)
	sess.LastActivity = time.Now()
	if !ok {
		sess.History = append(sess.History, Message{
			Role:    "system",
			Content: "You are an advanced AI neural system for a NLU Chatbot. The NLU and you are different component of the same brain. Your job is to provide short answers for what the NLU is not capable of. Limit your response to one sentences. ",
		})
	}
	sess.History = append(sess.History, Message{
		Role:    role,
		Content: content,
	})
	sessionHistories.Store(sessionID, sess)
	return &sess.History
}

func GetSessionHistory(sessionID string) []Message {
	session, ok := sessionHistories.Load(sessionID)
	if !ok {
		log.Print("New session created for sessionID: ", sessionID)
		var history []Message
		history = append(history, Message{
			Role:    "system",
			Content: "You are an advanced AI neural system for a NLU Chatbot. The NLU and you are different component of the same brain. Your job is to provide short answers for what the NLU is not capable of. Limit your response to one sentences. ",
		})
		return history
	}
	sess := session.(Session)
	sess.LastActivity = time.Now()
	sessionHistories.Store(sessionID, sess)
	return sess.History
}
