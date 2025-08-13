package processes

import (
	"context"

	"github.com/wonderarry/rwmsredone/internal/app/contract"
	"github.com/wonderarry/rwmsredone/internal/domain"
)

type service struct {
	uow       contract.UnitOfWork
	templates contract.TemplateProvider
	idgen     contract.IDGen
}

func New(uow contract.UnitOfWork, templates contract.TemplateProvider, idgen contract.IDGen) Service {
	return &service{uow: uow, templates: templates, idgen: idgen}
}

func (s *service) CreateProcess(ctx context.Context, cmd CreateProcess) (domain.ProcessID, error) {
	var pid domain.ProcessID

	err := s.uow.WithTx(ctx, func(ctx context.Context, tx contract.Tx) error {
		isLeader, err := tx.Projects().IsMember(ctx, cmd.ProjectID, cmd.ActorID, domain.RoleProjectLeader)
		if err != nil {
			return err
		}
		if !isLeader {
			return domain.ErrForbidden
		}

		tpl, err := s.templates.Load(ctx, cmd.TemplateKey)
		if err != nil {
			return err
		}
		pr := &domain.Process{
			ID:           domain.ProcessID(s.idgen.NewID()),
			ProjectID:    cmd.ProjectID,
			TemplateKey:  cmd.TemplateKey,
			Name:         cmd.Name,
			CurrentStage: tpl.Start,
			State:        domain.ProcessActive,
		}

		if err := tx.Processes().Create(ctx, pr); err != nil {
			return err
		}

		if err := tx.Outbox().Append(ctx, domain.ProcessCreated{
			ProcessID:   pr.ID,
			ProjectID:   pr.ProjectID,
			TemplateKey: pr.TemplateKey,
			Name:        pr.Name,
		}); err != nil {
			return err
		}

		pid = pr.ID
		return nil
	})

	return pid, err
}

func (s *service) AddMember(ctx context.Context, cmd AddProcessMember) error {
	return s.uow.WithTx(ctx, func(ctx context.Context, tx contract.Tx) error {
		projectID, err := tx.Processes().ParentProjectID(ctx, cmd.ProcessID)
		if err != nil {
			return err
		}

		isLeader, err := tx.Projects().IsMember(ctx, projectID, cmd.ActorID, domain.RoleProjectLeader)
		if err != nil {
			return err
		}
		if !isLeader {
			return domain.ErrForbidden
		}

		isProjMem, err := tx.Projects().IsMember(ctx, projectID, cmd.AccountID, domain.RoleProjectMember)
		if err != nil {
			return err
		}
		if !isProjMem {
			return domain.ErrForbidden
		}

		if err := tx.Processes().AddMember(ctx, domain.ProcessMember{
			ProcessID: cmd.ProcessID,
			AccountID: cmd.AccountID,
			RoleKey:   cmd.Role,
		}); err != nil {
			return err
		}

		return tx.Outbox().Append(ctx, domain.ProcessMemberAdded{
			ProcessID: cmd.ProcessID,
			AccountID: cmd.AccountID,
			RoleKey:   cmd.Role,
		})
	})
}

func (s *service) RemoveMember(ctx context.Context, cmd RemoveProcessMember) error {
	return s.uow.WithTx(ctx, func(ctx context.Context, tx contract.Tx) error {
		projectID, err := tx.Processes().ParentProjectID(ctx, cmd.ProcessID)
		if err != nil {
			return err
		}

		isLeader, err := tx.Projects().IsMember(ctx, projectID, cmd.ActorID, domain.RoleProjectLeader)
		if err != nil {
			return err
		}
		if !isLeader {
			return domain.ErrForbidden
		}

		if err := tx.Processes().RemoveMember(ctx, domain.ProcessMember{
			ProcessID: cmd.ProcessID,
			AccountID: cmd.AccountID,
			RoleKey:   cmd.Role,
		}); err != nil {
			return err
		}

		return tx.Outbox().Append(ctx, domain.ProcessMemberRemoved{
			ProcessID: cmd.ProcessID,
			AccountID: cmd.AccountID,
			RoleKey:   cmd.Role,
		})
	})
}

func (s *service) RecordApproval(ctx context.Context, cmd RecordApproval) error {
	return s.uow.WithTx(ctx, func(ctx context.Context, tx contract.Tx) error {
		pr, err := tx.Processes().Get(ctx, cmd.ProcessID)
		if err != nil {
			return err
		}

		ok, err := tx.Processes().IsMember(ctx, pr.ID, cmd.ActorID, cmd.ActorRole)
		if err != nil {
			return err
		}
		if !ok {
			return domain.ErrForbidden
		}

		if err := tx.Approvals().Upsert(ctx, domain.Approval{
			ProcessID:   pr.ID,
			StageKey:    pr.CurrentStage,
			ByAccountID: cmd.ActorID,
			ByRole:      cmd.ActorRole,
			Decision:    cmd.Decision,
			Comment:     cmd.Comment,
		}); err != nil {
			return err
		}

		tpl, err := s.templates.Load(ctx, pr.TemplateKey)
		if err != nil {
			return err
		}
		required := tpl.Stages[pr.CurrentStage].RequiredRole

		count, err := tx.Approvals().CountByDecisionAndRole(
			ctx, pr.ID, pr.CurrentStage, required, domain.Approve,
		)
		if err != nil {
			return err
		}

		next, done := domain.Evaluate(tpl, pr.CurrentStage, cmd.Decision, count)

		// Always emit vote event.
		if err := tx.Outbox().Append(ctx, domain.ApprovalRecorded{
			ProcessID: pr.ID,
			StageKey:  pr.CurrentStage,
			ByAccount: cmd.ActorID,
			Decision:  cmd.Decision,
		}); err != nil {
			return err
		}

		// Advance if needed; finalize if terminal.
		if next != pr.CurrentStage {
			if err := tx.Processes().SetCurrentStage(ctx, pr.ID, next); err != nil {
				return err
			}
			if err := tx.Outbox().Append(ctx, domain.StageAdvanced{
				ProcessID: pr.ID,
				From:      pr.CurrentStage,
				To:        next,
			}); err != nil {
				return err
			}
			if done {
				if err := tx.Processes().SetState(ctx, pr.ID, domain.ProcessCompleted); err != nil {
					return err
				}
				if err := tx.Outbox().Append(ctx, domain.ProcessFinalized{
					ProcessID: pr.ID,
				}); err != nil {
					return err
				}
			}
		}

		return nil
	})
}

func (s *service) GetProcess(ctx context.Context, id domain.ProcessID) (*ProcessDetail, error) {
	var out *ProcessDetail
	err := s.uow.WithTx(ctx, func(ctx context.Context, tx contract.Tx) error {
		pr, err := tx.Processes().Get(ctx, id)
		if err != nil {
			return err
		}
		members, err := tx.Processes().ListMembers(ctx, id)
		if err != nil {
			return err
		}
		out = &ProcessDetail{
			Process: *pr,
			Members: members,
		}
		return nil
	})
	return out, err
}

func (s *service) GetProcessGraph(ctx context.Context, id domain.ProcessID) (*Graph, error) {
	return nil, domain.ErrNotImplemented
}

func (s *service) ListApprovals(ctx context.Context, id domain.ProcessID, stage domain.StageKey) ([]domain.Approval, error) {
	var out []domain.Approval
	err := s.uow.WithTx(ctx, func(ctx context.Context, tx contract.Tx) error {
		apprs, err := tx.Approvals().ListForStage(ctx, id, stage)
		if err != nil {
			return err
		}
		out = apprs
		return nil
	})
	return out, err
}
