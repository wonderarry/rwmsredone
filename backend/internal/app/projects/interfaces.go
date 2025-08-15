package projects

import (
	"context"

	"github.com/wonderarry/rwmsredone/internal/domain"
)

type Service interface {
	CreateProject(ctx context.Context, cmd CreateProject) (domain.ProjectID, error)
	EditProjectMeta(ctx context.Context, cmd EditProjectMeta) error
	AddProjectMember(ctx context.Context, cmd AddProjectMember) error
	RemoveProjectMember(ctx context.Context, cmd RemoveProjectMember) error

	ListMyProjects(ctx context.Context, actorID domain.AccountID) ([]ProjectCard, error)
	GetProject(ctx context.Context, id domain.ProjectID) (*ProjectDetail, error)
}

type ProjectCard struct {
	ID          domain.ProjectID
	Name        string
	Theme       string
	Description string
}

type ProjectDetail struct {
	Project domain.Project
	Members []domain.ProjectMember
}

type CreateProject struct {
	Name        string
	Theme       string
	Description string
	ActorID     domain.AccountID
}

type EditProjectMeta struct {
	ProjectID                domain.ProjectID
	Name, Theme, Description string
	ActorID                  domain.AccountID
}

type AddProjectMember struct {
	ProjectID domain.ProjectID
	AccountID domain.AccountID
	Role      domain.ProjectRole
	ActorID   domain.AccountID
}

type RemoveProjectMember struct {
	ProjectID domain.ProjectID
	AccountID domain.AccountID
	Role      domain.ProjectRole
	ActorID   domain.AccountID
}
