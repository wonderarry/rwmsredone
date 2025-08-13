package db

import (
	"context"

	"github.com/wonderarry/rwmsredone/infra/db/sqlc"
	"github.com/wonderarry/rwmsredone/internal/app/contract"
	"github.com/wonderarry/rwmsredone/internal/domain"
)

type approvalRepo struct{ q *sqlc.Queries }

var _ contract.ApprovalRepo = (*approvalRepo)(nil)

func (r *approvalRepo) Upsert(ctx context.Context, a domain.Approval) error {
	return r.q.UpsertApproval(ctx, sqlc.UpsertApprovalParams{
		ProcessID:   a.ProcessID,
		StageKey:    a.StageKey,
		ByAccountID: a.ByAccountID,
		ByRole:      string(a.ByRole),
		Decision:    string(a.Decision),
		Comment:     a.Comment,
	})
}

func (r *approvalRepo) CountByDecisionAndRole(ctx context.Context, processID domain.ProcessID, stage domain.StageKey, role domain.ProcessRole, decision domain.Decision) (int, error) {
	n, err := r.q.CountApprovalByDecisionAndRole(ctx, sqlc.CountApprovalByDecisionAndRoleParams{
		ProcessID: processID,
		StageKey:  stage,
		ByRole:    string(role),
		Decision:  string(decision),
	})
	return int(n), err
}

func (r *approvalRepo) ListForStage(ctx context.Context, processID domain.ProcessID, stage domain.StageKey) ([]domain.Approval, error) {
	rows, err := r.q.ListApprovalsByProcessAndStage(ctx, sqlc.ListApprovalsByProcessAndStageParams{
		ProcessID: processID,
		StageKey:  stage,
	})
	if err != nil {
		return nil, err
	}
	out := make([]domain.Approval, 0, len(rows))
	for _, rw := range rows {
		out = append(out, domain.Approval{
			ProcessID:   rw.ProcessID,
			StageKey:    rw.StageKey,
			ByAccountID: rw.ByAccountID,
			ByRole:      domain.ProcessRole(rw.ByRole),
			Decision:    domain.Decision(rw.Decision),
			Comment:     rw.Comment,
		})
	}
	return out, nil
}
