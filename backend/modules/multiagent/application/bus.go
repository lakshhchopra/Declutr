package application

import (
	"context"
	"sync"
	"time"

	"github.com/diablovocado/declutr/modules/multiagent/domain"
	"github.com/diablovocado/declutr/shared/observability"
)

type MessageHandler func(msg *domain.AgentMessage) (*domain.AgentMessage, error)

type MessageBus struct {
	mu           sync.RWMutex
	handlers     map[string]MessageHandler // Receiver Agent ID -> Handler
	auditHistory []*domain.AgentMessage
}

func NewMessageBus() *MessageBus {
	return &MessageBus{
		handlers:     make(map[string]MessageHandler),
		auditHistory: []*domain.AgentMessage{},
	}
}

func (mb *MessageBus) RegisterHandler(agentID string, handler MessageHandler) {
	mb.mu.Lock()
	defer mb.mu.Unlock()
	mb.handlers[agentID] = handler
}

func (mb *MessageBus) Send(ctx context.Context, msg *domain.AgentMessage) (*domain.AgentMessage, error) {
	if msg.ID == "" {
		msg.ID = "msg-" + observability.GenerateID(8)
	}
	if msg.Timestamp.IsZero() {
		msg.Timestamp = time.Now().UTC()
	}

	mb.mu.Lock()
	mb.auditHistory = append(mb.auditHistory, msg)
	handler, exists := mb.handlers[msg.Receiver]
	mb.mu.Unlock()

	if !exists && msg.Receiver != "BROADCAST" {
		// Mock response if handler not registered
		resp := &domain.AgentMessage{
			ID:            "msg-" + observability.GenerateID(8),
			CorrelationID: msg.CorrelationID,
			GoalID:        msg.GoalID,
			TaskID:        msg.TaskID,
			Sender:        msg.Receiver,
			Receiver:      msg.Sender,
			MessageType:   "RESPONSE",
			Payload:       map[string]interface{}{"status": "SUCCESS", "message": "Specialist task executed successfully"},
			Timestamp:     time.Now().UTC(),
		}
		mb.mu.Lock()
		mb.auditHistory = append(mb.auditHistory, resp)
		mb.mu.Unlock()
		return resp, nil
	}

	if exists {
		return handler(msg)
	}

	return msg, nil
}

func (mb *MessageBus) GetAuditHistory() []*domain.AgentMessage {
	mb.mu.RLock()
	defer mb.mu.RUnlock()
	history := make([]*domain.AgentMessage, len(mb.auditHistory))
	copy(history, mb.auditHistory)
	return history
}
