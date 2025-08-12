package app

import (
	"context"

	"github.com/wonderarry/rwmsredone/internal/domain"
)

type UnitOfWork interface {
	WithTx(ctx context.Context, fn func(ctx context.Context, tx Tx) error) error
}

type Tx interface {
	Accounts() AccountRepo
	Identities() IdentityRepo
	Projects() ProjectRepo
	Processes() ProcessRepo
	Approvals() ApprovalRepo
	Outbox() OutboxRepo
}

type AccountRepo interface {
	Get(ctx context.Context, id domain.AccountID) (*domain.Account, error)
	Create(ctx context.Context, a *domain.Account) error
	UpdateProfile(ctx context.Context, a *domain.Account) error
	AddGlobalRole(ctx context.Context, a *domain.Account) error
	HasGlobalRole(ctx context.Context, a *domain.Account) error
}

type IdentityRepo interface {
	Create(ctx context.Context, i *domain.Identity) error
	FindByProviderSubject(ctx context.Context, p domain.IdentityProvider, sub string) (*domain.Identity, error)

	// Unsure about these currently

	// FindLocalByLogin(ctx context.Context, login string) (*domain.Identity, error)
	// GetLocalCredentials(ctx context.Context, identityID int64) (*domain.LocalCredentials, error)
	// UpsertLocalCredentials(ctx context.Context, c domain.LocalCredentials) error
}
