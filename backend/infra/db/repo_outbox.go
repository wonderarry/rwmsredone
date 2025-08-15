package db

import (
	"context"
	"encoding/json"

	"github.com/wonderarry/rwmsredone/infra/db/sqlc"
	"github.com/wonderarry/rwmsredone/internal/app/contract"
	"github.com/wonderarry/rwmsredone/internal/domain"
)

type outboxRepo struct{ q *sqlc.Queries }

var _ contract.OutboxRepo = (*outboxRepo)(nil)

func (r *outboxRepo) Append(ctx context.Context, e domain.Event) error {
	payload, err := json.Marshal(e)
	if err != nil {
		return err
	}
	return r.q.AppendOutbox(ctx, sqlc.AppendOutboxParams{
		Topic:   e.Topic(),
		Payload: payload,
	})
}
