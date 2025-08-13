package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/wonderarry/rwmsredone/infra/db/sqlc"
	"github.com/wonderarry/rwmsredone/internal/app/contract"
	"github.com/wonderarry/rwmsredone/internal/domain"
)

type accountRepo struct{ q *sqlc.Queries }

var _ contract.AccountRepo = (*accountRepo)(nil)

func (r *accountRepo) Get(ctx context.Context, id domain.AccountID) (*domain.Account, error) {
	row, err := r.q.GetAccount(ctx, id)
	if err != nil {
		return nil, err
	}
	return &domain.Account{
		ID:          row.ID,
		FirstName:   row.FirstName.String,
		MiddleName:  row.MiddleName.String,
		LastName:    row.LastName.String,
		GroupNumber: row.Grp.String,
		CreatedAt:   row.CreatedAt.Time,
		UpdatedAt:   row.UpdatedAt.Time,
	}, nil
}

func (r *accountRepo) Create(ctx context.Context, a *domain.Account) error {
	return r.q.CreateAccount(ctx, sqlc.CreateAccountParams{
		ID:         a.ID,
		FirstName:  pgtype.Text{String: a.FirstName, Valid: true},
		MiddleName: pgtype.Text{String: a.MiddleName, Valid: true},
		LastName:   pgtype.Text{String: a.LastName, Valid: true},
		Grp:        pgtype.Text{String: a.GroupNumber, Valid: true},
	})
}

func (r *accountRepo) GrantGlobalRole(ctx context.Context, id domain.AccountID, role domain.GlobalRole) error {
	return r.q.GrantGlobalRole(ctx, sqlc.GrantGlobalRoleParams{
		AccountID: id,
		RoleKey:   string(role),
	})
}

func (r *accountRepo) HasGlobalRole(ctx context.Context, id domain.AccountID, role domain.GlobalRole) (bool, error) {
	ok, err := r.q.HasGlobalRole(ctx, sqlc.HasGlobalRoleParams{
		AccountID: id,
		RoleKey:   string(role),
	})
	if err != nil {
		return false, err
	}
	return ok, nil
}

func (r *accountRepo) ListGlobalRoles(ctx context.Context, a *domain.Account) ([]domain.GlobalRole, error) {
	rows, err := r.q.ListGlobalRoles(ctx, a.ID)
	if err != nil {
		return nil, err
	}
	out := make([]domain.GlobalRole, 0, len(rows))
	for _, rr := range rows {
		parsed, notok := parseGlobalRole(rr)
		if notok != nil {
			return nil, err
		}
		out = append(out, parsed)
	}
	return out, nil
}
