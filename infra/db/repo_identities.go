package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/wonderarry/rwmsredone/infra/db/sqlc"
	"github.com/wonderarry/rwmsredone/internal/app/contract"
	"github.com/wonderarry/rwmsredone/internal/domain"
)

type identityRepo struct{ q *sqlc.Queries }

var _ contract.IdentityRepo = (*identityRepo)(nil)

func (r *identityRepo) Create(ctx context.Context, i *domain.Identity) error {
	return r.q.CreateIdentity(ctx, sqlc.CreateIdentityParams{
		ID:           i.ID,
		AccountID:    i.AccountID,
		Provider:     string(i.Provider),
		Subject:      i.Subject,
		Email:        textFromPtr(i.Email),
		PasswordHash: textFromPtr(i.PasswordHash),
		RefreshToken: textFromPtr(i.RefreshToken),
		ExpiresAt:    tsFromPtr(i.ExpiresAt),
	})
}

func (r *identityRepo) GetByProviderSubject(ctx context.Context, p domain.IdentityProvider, sub string) (*domain.Identity, error) {
	row, err := r.q.GetIdentityByProviderSubject(ctx, sqlc.GetIdentityByProviderSubjectParams{
		Provider: string(p),
		Subject:  sub,
	})
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil // not found -> nil, nil for service logic convenience
		}
		return nil, err
	}
	return &domain.Identity{
		ID:           row.ID,
		AccountID:    row.AccountID,
		Provider:     domain.IdentityProvider(row.Provider),
		Subject:      row.Subject,
		Email:        textPtr(row.Email),
		PasswordHash: textPtr(row.PasswordHash),
		RefreshToken: textPtr(row.RefreshToken),
		ExpiresAt:    tsPtr(row.ExpiresAt),
		CreatedAt:    *tsPtr(row.CreatedAt),
		UpdatedAt:    *tsPtr(row.UpdatedAt),
	}, nil
}

func (r *identityRepo) ListByAccount(ctx context.Context, id domain.AccountID) ([]*domain.Identity, error) {
	rows, err := r.q.ListIdentitiesByAccount(ctx, id)
	if err != nil {
		return nil, err
	}
	out := make([]*domain.Identity, 0, len(rows))
	for _, rw := range rows {
		out = append(out, &domain.Identity{
			ID:           rw.ID,
			AccountID:    rw.AccountID,
			Provider:     domain.IdentityProvider(rw.Provider),
			Subject:      rw.Subject,
			Email:        textPtr(rw.Email),
			PasswordHash: textPtr(rw.PasswordHash),
			RefreshToken: textPtr(rw.RefreshToken),
			ExpiresAt:    tsPtr(rw.ExpiresAt),
			CreatedAt:    *tsPtr(rw.CreatedAt),
			UpdatedAt:    *tsPtr(rw.UpdatedAt),
		})
	}
	return out, nil
}
