package application

import (
	"context"
	"time"

	"github.com/diablovocado/declutr/modules/lifeos/domain"
	"github.com/diablovocado/declutr/shared/observability"
)

type ProjectEngine struct{}

func NewProjectEngine() *ProjectEngine {
	return &ProjectEngine{}
}

func (pe *ProjectEngine) CreateProject(
	ctx context.Context,
	userID string,
	lifeAreaID string,
	name string,
	description string,
	targetDate time.Time,
) (*domain.Project, error) {
	now := time.Now().UTC()
	return &domain.Project{
		ID:             "prj-" + observability.GenerateID(8),
		UserID:         userID,
		LifeAreaID:     lifeAreaID,
		Name:           name,
		Description:    description,
		Status:         domain.ProjectInProgress,
		AssociatedDocs: []string{},
		People:         []string{},
		TargetDate:     targetDate,
		CreatedAt:      now,
		UpdatedAt:      now,
	}, nil
}
