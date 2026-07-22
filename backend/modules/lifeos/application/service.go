package application

import (
	"context"
	"time"

	"github.com/diablovocado/declutr/modules/lifeos/domain"
	"github.com/diablovocado/declutr/modules/lifeos/repository"
)

type LifeOSService struct {
	repo          repository.LifeOSRepository
	projectEngine *ProjectEngine
	goalEngine    *GoalEngine
	graphEngine   *LifeGraphEngine
}

func NewLifeOSService(repo repository.LifeOSRepository) *LifeOSService {
	return &LifeOSService{
		repo:          repo,
		projectEngine: NewProjectEngine(),
		goalEngine:    NewGoalEngine(),
		graphEngine:   NewLifeGraphEngine(),
	}
}

func (s *LifeOSService) GetUnifiedDashboard(ctx context.Context, userID string) (*domain.LifeDashboard, error) {
	areas, err := s.repo.ListLifeAreas(ctx, userID)
	if err != nil {
		return nil, err
	}

	projects, err := s.repo.ListProjects(ctx, userID)
	if err != nil {
		return nil, err
	}

	goals, err := s.repo.ListGoals(ctx, userID, "")
	if err != nil {
		return nil, err
	}

	timeline, err := s.repo.ListTimelineEvents(ctx, userID)
	if err != nil {
		return nil, err
	}

	priorities := []string{
		"Renew passport for Japan Trip",
		"Review Pitch Deck for Launch Startup",
		"Check Q3 tax filing receipts",
	}

	return &domain.LifeDashboard{
		LifeAreas:       areas,
		ActiveProjects:  projects,
		ActiveGoals:     goals,
		PrioritiesToday: priorities,
		RecentTimeline:  timeline,
		HealthScore:     94,
	}, nil
}

func (s *LifeOSService) ListLifeAreas(ctx context.Context, userID string) ([]domain.LifeArea, error) {
	return s.repo.ListLifeAreas(ctx, userID)
}

func (s *LifeOSService) ListProjects(ctx context.Context, userID string) ([]domain.Project, error) {
	return s.repo.ListProjects(ctx, userID)
}

func (s *LifeOSService) CreateProject(ctx context.Context, userID string, lifeAreaID string, name string, desc string, targetDate time.Time) (*domain.Project, error) {
	prj, err := s.projectEngine.CreateProject(ctx, userID, lifeAreaID, name, desc, targetDate)
	if err != nil {
		return nil, err
	}
	if err := s.repo.CreateProject(ctx, prj); err != nil {
		return nil, err
	}
	return prj, nil
}

func (s *LifeOSService) ListGoals(ctx context.Context, userID string, projectID string) ([]domain.ProjectGoal, error) {
	return s.repo.ListGoals(ctx, userID, projectID)
}

func (s *LifeOSService) CreateGoal(ctx context.Context, userID string, projectID string, title string, desc string, dueDate time.Time) (*domain.ProjectGoal, error) {
	goal, err := s.goalEngine.CreateGoal(ctx, userID, projectID, title, desc, dueDate)
	if err != nil {
		return nil, err
	}
	if err := s.repo.CreateGoal(ctx, goal); err != nil {
		return nil, err
	}
	return goal, nil
}

func (s *LifeOSService) UpdateGoalProgress(ctx context.Context, goalID string, progressPct int, isCompleted bool) error {
	return s.repo.UpdateGoalProgress(ctx, goalID, progressPct, isCompleted)
}

func (s *LifeOSService) GetTimeline(ctx context.Context, userID string) ([]domain.LifeTimelineEvent, error) {
	return s.repo.ListTimelineEvents(ctx, userID)
}

func (s *LifeOSService) GetMetrics(ctx context.Context, userID string) (*domain.LifeMetric, error) {
	return s.repo.GetLifeMetrics(ctx, userID)
}
