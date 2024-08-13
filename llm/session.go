package llm

import (
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

func UpdateSessionHistory(sessionID, role, content string) {
	mu.Lock()
	defer mu.Unlock()

	session, _ := sessionHistories.LoadOrStore(sessionID, Session{
		LastActivity: time.Now(),
		History:      []Message{},
	})
	sess := session.(Session)
	sess.LastActivity = time.Now()
	sess.History = append(sess.History, Message{
		Role:    role,
		Content: content,
	})
	sessionHistories.Store(sessionID, sess)
}

func GetSessionHistory(sessionID string) []Message {
	session, ok := sessionHistories.Load(sessionID)
	if !ok {
		return []Message{}
	}
	sess := session.(Session)
	sess.LastActivity = time.Now()
	sessionHistories.Store(sessionID, sess)
	return sess.History
}
