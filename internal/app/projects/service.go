package projects

import (
	"context"

	"github.com/wonderarry/rwmsredone/internal/app"
	"github.com/wonderarry/rwmsredone/internal/domain"
)

type service struct {
	uow   app.UnitOfWork
	idgen app.IDGen
}

func New(uow app.UnitOfWork) Service { return &service{uow: uow} }

func (s *service) CreateProject(ctx context.Context, cmd CreateProject) (domain.ProjectID, error) {
	var id domain.ProjectID

	err := s.uow.WithTx(ctx, func(ctx context.Context, tx app.Tx) error {
		ok, err := tx.Accounts().HasGlobalRole(ctx, cmd.ActorID, domain.RoleCanCreateProjects)
		if err != nil {
			return err
		}
		if !ok {
			return domain.ErrForbidden
		}

		pid := domain.ProjectID(s.idgen.NewID())
		p := &domain.Project{
			ID:          pid,
			Name:        cmd.Name,
			Theme:       cmd.Theme,
			Description: cmd.Description,
			CreatedBy:   cmd.ActorID,
		}
		if err := tx.Projects().Create(ctx, p); err != nil {
			return err
		}

		_ = tx.Projects().AddMember(ctx, domain.ProjectMember{ProjectID: p.ID, AccountID: cmd.ActorID, RoleKey: domain.RoleProjectLeader})
		_ = tx.Projects().AddMember(ctx, domain.ProjectMember{ProjectID: p.ID, AccountID: cmd.ActorID, RoleKey: domain.RoleProjectMember})

		if err := tx.Outbox().Append(ctx, domain.ProjectCreated{ProjectID: p.ID, Name: p.Name, By: cmd.ActorID}); err != nil {
			return err
		}
		id = p.ID

		return nil
	})

	return id, err
}
func (s *service) EditProjectMeta(ctx context.Context, cmd EditProjectMeta) error
func (s *service) AddProjectMember(ctx context.Context, cmd AddProjectMember) error
func (s *service) RemoveProjectMember(ctx context.Context, cmd RemoveProjectMember) error
func (s *service) ListMyProjects(ctx context.Context, actorID domain.AccountID) ([]ProjectCard, error)
func (s *service) GetProject(ctx context.Context, id domain.ProjectID) (*ProjectDetail, error)
