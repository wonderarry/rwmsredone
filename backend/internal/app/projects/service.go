package projects

import (
	"context"

	"github.com/wonderarry/rwmsredone/internal/app/contract"
	"github.com/wonderarry/rwmsredone/internal/domain"
)

type service struct {
	uow   contract.UnitOfWork
	idgen contract.IDGen
}

func New(uow contract.UnitOfWork, idgen contract.IDGen) Service {
	return &service{uow: uow, idgen: idgen}
}

func (s *service) CreateProject(ctx context.Context, cmd CreateProject) (domain.ProjectID, error) {
	var id domain.ProjectID

	err := s.uow.WithTx(ctx, func(ctx context.Context, tx contract.Tx) error {
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
func (s *service) EditProjectMeta(ctx context.Context, cmd EditProjectMeta) error {
	return s.uow.WithTx(ctx, func(ctx context.Context, tx contract.Tx) error {
		ok, err := tx.Projects().IsMember(ctx, cmd.ProjectID, cmd.ActorID, domain.RoleProjectLeader)
		if err != nil {
			return err
		}
		if !ok {
			return domain.ErrForbidden
		}

		p, err := tx.Projects().Get(ctx, cmd.ProjectID)

		if err != nil {
			return err
		}

		if cmd.Name != "" {
			p.Name = cmd.Name
		}
		if cmd.Theme != "" {
			p.Theme = cmd.Theme
		}
		if cmd.Description != "" {
			p.Description = cmd.Description
		}

		if err := tx.Projects().UpdateMeta(ctx, p); err != nil {
			return err
		}
		return nil
	})
}
func (s *service) AddProjectMember(ctx context.Context, cmd AddProjectMember) error {
	return s.uow.WithTx(ctx, func(ctx context.Context, tx contract.Tx) error {
		ok, err := tx.Projects().IsMember(ctx, cmd.ProjectID, cmd.ActorID, domain.RoleProjectLeader)
		if err != nil {
			return err
		}
		if !ok {
			return domain.ErrForbidden
		}
		m := domain.ProjectMember{
			ProjectID: cmd.ProjectID,
			AccountID: cmd.AccountID,
			RoleKey:   cmd.Role,
		}
		if err := tx.Projects().AddMember(ctx, m); err != nil {
			return err
		}

		return tx.Outbox().Append(ctx, domain.ProjectMemberAdded{
			ProjectID: cmd.ProjectID,
			AccountID: cmd.AccountID,
			RoleKey:   cmd.Role,
			AddedBy:   cmd.ActorID,
		})
	})
}

func (s *service) RemoveProjectMember(ctx context.Context, cmd RemoveProjectMember) error {
	return s.uow.WithTx(ctx, func(ctx context.Context, tx contract.Tx) error {
		ok, err := tx.Projects().IsMember(ctx, cmd.ProjectID, cmd.ActorID, domain.RoleProjectLeader)
		if err != nil {
			return err
		}
		if !ok {
			return domain.ErrForbidden
		}

		m := domain.ProjectMember{
			ProjectID: cmd.ProjectID,
			AccountID: cmd.AccountID,
			RoleKey:   cmd.Role,
		}
		if err := tx.Projects().RemoveMember(ctx, m); err != nil {
			return err
		}

		return tx.Outbox().Append(ctx, domain.ProjectMemberRemoved{
			ProjectID: cmd.ProjectID,
			AccountID: cmd.AccountID,
			RoleKey:   cmd.Role,
			RemovedBy: cmd.ActorID,
		})
	})
}
func (s *service) ListMyProjects(ctx context.Context, actorID domain.AccountID) ([]ProjectCard, error) {
	var cards []ProjectCard
	err := s.uow.WithTx(ctx, func(ctx context.Context, tx contract.Tx) error {
		ps, err := tx.Projects().ListForAccount(ctx, actorID)
		if err != nil {
			return err
		}
		cards = make([]ProjectCard, 0, len(ps))
		for _, p := range ps {
			cards = append(cards, ProjectCard{
				ID:          p.ID,
				Name:        p.Name,
				Theme:       p.Theme,
				Description: p.Description,
			})
		}
		return nil
	})
	return cards, err
}
func (s *service) GetProject(ctx context.Context, id domain.ProjectID) (*ProjectDetail, error) {
	var out *ProjectDetail
	err := s.uow.WithTx(ctx, func(ctx context.Context, tx contract.Tx) error {
		p, err := tx.Projects().Get(ctx, id)
		if err != nil {
			return err
		}
		members, err := tx.Projects().ListMembers(ctx, id)
		if err != nil {
			return err
		}
		out = &ProjectDetail{Project: *p, Members: members}
		return nil
	})
	return out, err
}
