package accounts

import (
	"context"

	"github.com/wonderarry/rwmsredone/internal/domain"
)

type Service interface {
	GetMe(ctx context.Context, actorID domain.AccountID) (*Me, error)
}

type Me struct {
	Account    domain.Account
	Roles      []domain.GlobalRole
	Identities []IdentityInfo
}

type IdentityInfo struct {
	Provider string
	Info     string
	Email    string
}
