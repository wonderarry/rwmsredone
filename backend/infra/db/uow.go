package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/wonderarry/rwmsredone/infra/db/sqlc"
	"github.com/wonderarry/rwmsredone/internal/app/contract"
)

type UoW struct {
	pool *pgxpool.Pool
}

func NewUoW(pool *pgxpool.Pool) *UoW {
	return &UoW{pool: pool}
}

func (u *UoW) WithTx(ctx context.Context, fn func(ctx context.Context, tx contract.Tx) error) error {
	pgxtx, err := u.pool.BeginTx(ctx, pgx.TxOptions{
		// IsoLevel: pgx.Serializable,
		// AccessMode: pgx.ReadWrite,
	})
	if err != nil {
		return err
	}

	q := sqlc.New(pgxtx)

	t := &txImpl{
		tx: pgxtx,
		q:  q,
	}

	if err := fn(ctx, t); err != nil {
		_ = pgxtx.Rollback(ctx)
		return err
	}
	return pgxtx.Commit(ctx)
}

type txImpl struct {
	tx pgx.Tx
	q  *sqlc.Queries
}

func (t *txImpl) Accounts() contract.AccountRepo    { return &accountRepo{q: t.q} }
func (t *txImpl) Identities() contract.IdentityRepo { return &identityRepo{q: t.q} }
func (t *txImpl) Projects() contract.ProjectRepo    { return &projectRepo{q: t.q} }
func (t *txImpl) Processes() contract.ProcessRepo   { return &processRepo{q: t.q} }
func (t *txImpl) Approvals() contract.ApprovalRepo  { return &approvalRepo{q: t.q} }
func (t *txImpl) Outbox() contract.OutboxRepo       { return &outboxRepo{q: t.q} }
