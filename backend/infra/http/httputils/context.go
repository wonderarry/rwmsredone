package httputils

import (
	"context"

	"github.com/wonderarry/rwmsredone/internal/domain"
)

type ctxKey string

const actorKey ctxKey = "actorID"

func WithActor(ctx context.Context, id domain.AccountID) context.Context {
	return context.WithValue(ctx, actorKey, id)
}

func ActorIDFrom(ctx context.Context) domain.AccountID {
	if v := ctx.Value(actorKey); v != nil {
		return v.(domain.AccountID)
	}
	return ""
}
