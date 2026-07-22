package application

import (
	"context"
	"time"

	"github.com/diablovocado/declutr/modules/lifeos/domain"
	"github.com/diablovocado/declutr/shared/observability"
)

type GoalEngine struct{}

func NewGoalEngine() *GoalEngine {
	return &GoalEngine{}
}

func (ge *GoalEngine) CreateGoal(
	ctx context.Context,
	userID string,
	projectID string,
	title string,
	description string,
	dueDate time.Time,
) (*domain.ProjectGoal, error) {
	now := time.Now().UTC()
	return &domain.ProjectGoal{
		ID:            "gol-los-" + observability.GenerateID(8),
		ProjectID:     projectID,
		UserID:        userID,
		Title:         title,
		Description:   description,
		ProgressPct:   0,
		IsCompleted:   false,
		MissingAssets: []string{},
		DueDate:       dueDate,
		CreatedAt:     now,
		UpdatedAt:     now,
	}, nil
}
