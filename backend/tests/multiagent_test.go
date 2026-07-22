package tests

import (
	"context"
	"testing"

	multiApp "github.com/diablovocado/declutr/modules/multiagent/application"
	multiDomain "github.com/diablovocado/declutr/modules/multiagent/domain"
	multiRepo "github.com/diablovocado/declutr/modules/multiagent/repository"
)

func TestMultiAgentCoordinatorOrchestration(t *testing.T) {
	repo := multiRepo.NewInMemoryMultiAgentRepository()
	service := multiApp.NewMultiAgentService(repo)
	ctx := context.Background()

	// 1. Process User Goal via Coordinator
	graph, consensus, err := service.ProcessUserGoal(ctx, "Audit Q3 Tax Receipts and Expiring Passports")
	if err != nil || graph == nil || consensus == nil {
		t.Fatalf("Failed to orchestrate goal: %v", err)
	}

	if graph.Status != "COMPLETED" {
		t.Errorf("Expected task graph status COMPLETED, got %s", graph.Status)
	}

	if len(graph.Tasks) < 4 {
		t.Errorf("Expected 4 tasks in DAG task graph, got %d", len(graph.Tasks))
	}

	if !consensus.ConsensusAchieved {
		t.Error("Expected consensus achieved=true")
	}

	// 2. Structured Message Bus Audit Log Check
	logs, err := service.GetMessageAuditLogs(ctx)
	if err != nil || len(logs) == 0 {
		t.Errorf("Expected message bus audit logs recorded, got %d logs", len(logs))
	}

	// 3. Shared Memory Persistence Check
	mems, _ := repo.ListSharedMemory(ctx, graph.GoalID)
	if len(mems) < 1 {
		t.Errorf("Expected shared memory item recorded, got %d items", len(mems))
	}
}

func TestStructuredMessageBusRouting(t *testing.T) {
	bus := multiApp.NewMessageBus()
	ctx := context.Background()

	// Register Specialist Handler
	bus.RegisterHandler("agt-SEARCH_AGENT", func(msg *multiDomain.AgentMessage) (*multiDomain.AgentMessage, error) {
		return &multiDomain.AgentMessage{
			CorrelationID: msg.CorrelationID,
			Sender:        "agt-SEARCH_AGENT",
			Receiver:      msg.Sender,
			MessageType:   "RESPONSE",
			Payload:       map[string]interface{}{"found_count": 12},
		}, nil
	})

	req := &multiDomain.AgentMessage{
		CorrelationID: "corr-123",
		GoalID:        "gol-123",
		TaskID:        "tsk-123",
		Sender:        "agt-coordinator",
		Receiver:      "agt-SEARCH_AGENT",
		MessageType:   "REQUEST",
		Payload:       map[string]interface{}{"query": "tax receipts"},
	}

	resp, err := bus.Send(ctx, req)
	if err != nil || resp == nil {
		t.Fatalf("Message bus send failed: %v", err)
	}

	if resp.Sender != "agt-SEARCH_AGENT" {
		t.Errorf("Expected response sender agt-SEARCH_AGENT, got %s", resp.Sender)
	}

	count := resp.Payload["found_count"]
	if count != 12 {
		t.Errorf("Expected payload found_count=12, got %v", count)
	}
}
