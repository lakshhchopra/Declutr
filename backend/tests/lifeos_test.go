package tests

import (
	"context"
	"testing"
	"time"

	lifeApp "github.com/diablovocado/declutr/modules/lifeos/application"
	lifeRepo "github.com/diablovocado/declutr/modules/lifeos/repository"
)

func TestLifeOSUnifiedDashboardAndProjectHub(t *testing.T) {
	repo := lifeRepo.NewInMemoryLifeOSRepository()
	service := lifeApp.NewLifeOSService(repo)
	ctx := context.Background()

	// 1. Get Unified Life Dashboard
	dash, err := service.GetUnifiedDashboard(ctx, "usr-default")
	if err != nil || dash == nil {
		t.Fatalf("Failed to fetch LifeOS dashboard: %v", err)
	}

	if len(dash.LifeAreas) < 12 {
		t.Errorf("Expected 12 default Life Areas, got %d", len(dash.LifeAreas))
	}

	if len(dash.ActiveProjects) < 2 {
		t.Errorf("Expected at least 2 active projects, got %d", len(dash.ActiveProjects))
	}

	// 2. Create New First-Class Project
	newPrj, err := service.CreateProject(ctx, "usr-default", "area-Education", "Masters Application", "Prepare SOP and transcripts", time.Now().Add(60*24*time.Hour))
	if err != nil || newPrj == nil {
		t.Fatalf("Failed to create project: %v", err)
	}

	// 3. Create Goal belonging to Project
	goal, err := service.CreateGoal(ctx, "usr-default", newPrj.ID, "Submit TOEFL Score", "Send official ETS scores to university", time.Now().Add(30*24*time.Hour))
	if err != nil || goal == nil {
		t.Fatalf("Failed to create goal: %v", err)
	}

	// 4. Update Goal Progress
	err = service.UpdateGoalProgress(ctx, goal.ID, 100, true)
	if err != nil {
		t.Fatalf("Failed to update goal progress: %v", err)
	}

	goals, _ := service.ListGoals(ctx, "usr-default", newPrj.ID)
	if len(goals) != 1 || !goals[0].IsCompleted {
		t.Errorf("Expected goal completed, got %v", goals)
	}
}

func TestLifeOSMetricsAndTimeline(t *testing.T) {
	repo := lifeRepo.NewInMemoryLifeOSRepository()
	service := lifeApp.NewLifeOSService(repo)
	ctx := context.Background()

	metrics, err := service.GetMetrics(ctx, "usr-default")
	if err != nil || metrics.LifeBalance < 80 {
		t.Errorf("Expected health life balance score >= 80, got %v", metrics)
	}

	timeline, err := service.GetTimeline(ctx, "usr-default")
	if err != nil || len(timeline) == 0 {
		t.Errorf("Expected life timeline events recorded, got %d", len(timeline))
	}
}
