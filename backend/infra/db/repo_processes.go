package db

import (
	"context"

	"github.com/wonderarry/rwmsredone/infra/db/sqlc"
	"github.com/wonderarry/rwmsredone/internal/app/contract"
	"github.com/wonderarry/rwmsredone/internal/domain"
)

type processRepo struct{ q *sqlc.Queries }

var _ contract.ProcessRepo = (*processRepo)(nil)

func (r *processRepo) Create(ctx context.Context, pr *domain.Process) error {
	return r.q.CreateProcess(ctx, sqlc.CreateProcessParams{
		ID:           pr.ID,
		ProjectID:    pr.ProjectID,
		TemplateKey:  pr.TemplateKey,
		Name:         pr.Name,
		CurrentStage: pr.CurrentStage,
		State:        string(pr.State),
	})
}

func (r *processRepo) Get(ctx context.Context, id domain.ProcessID) (*domain.Process, error) {
	row, err := r.q.GetProcess(ctx, id)
	if err != nil {
		return nil, err
	}
	return &domain.Process{
		ID:           row.ID,
		ProjectID:    row.ProjectID,
		TemplateKey:  row.TemplateKey,
		Name:         row.Name,
		CurrentStage: row.CurrentStage,
		State:        domain.ProcessState(row.State),
	}, nil
}

func (r *processRepo) SetCurrentStage(ctx context.Context, id domain.ProcessID, stage domain.StageKey) error {
	return r.q.SetProcessCurrentStage(ctx, sqlc.SetProcessCurrentStageParams{
		ID:           id,
		CurrentStage: stage,
	})
}

func (r *processRepo) SetState(ctx context.Context, id domain.ProcessID, state domain.ProcessState) error {
	return r.q.SetProcessState(ctx, sqlc.SetProcessStateParams{
		ID:    id,
		State: string(state),
	})
}

func (r *processRepo) AddMember(ctx context.Context, m domain.ProcessMember) error {
	return r.q.AddProcessMember(ctx, sqlc.AddProcessMemberParams{
		ProcessID: m.ProcessID,
		AccountID: m.AccountID,
		RoleKey:   string(m.RoleKey),
	})
}

func (r *processRepo) RemoveMember(ctx context.Context, m domain.ProcessMember) error {
	return r.q.RemoveProcessMember(ctx, sqlc.RemoveProcessMemberParams{
		ProcessID: m.ProcessID,
		AccountID: m.AccountID,
		RoleKey:   string(m.RoleKey),
	})
}

func (r *processRepo) IsMember(ctx context.Context, processID domain.ProcessID, accountID domain.AccountID, role domain.ProcessRole) (bool, error) {
	ok, err := r.q.IsProcessMember(ctx, sqlc.IsProcessMemberParams{
		ProcessID: processID,
		AccountID: accountID,
		RoleKey:   string(role),
	})
	if err != nil {
		return false, err
	}
	return ok, nil
}

func (r *processRepo) ParentProjectID(ctx context.Context, processID domain.ProcessID) (domain.ProjectID, error) {
	row, err := r.q.GetParentProjectID(ctx, processID)
	if err != nil {
		return "", err
	}
	return domain.ProjectID(row), nil
}

func (r *processRepo) ListMembers(ctx context.Context, id domain.ProcessID) ([]domain.ProcessMember, error) {
	rows, err := r.q.ListProcessMembers(ctx, id)
	if err != nil {
		return nil, err
	}
	out := make([]domain.ProcessMember, 0, len(rows))
	for _, rw := range rows {
		out = append(out, domain.ProcessMember{
			ProcessID: rw.ProcessID,
			AccountID: rw.AccountID,
			RoleKey:   domain.ProcessRole(rw.RoleKey),
		})
	}
	return out, nil
}
