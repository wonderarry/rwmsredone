package accounts

import (
	"context"

	"github.com/wonderarry/rwmsredone/internal/domain"
)

type Service interface {
	RegisterLocal(ctx context.Context, cmd RegisterLocal) (domain.AccountID, error)
	LoginLocal(ctx context.Context, login, password string) (Token, error)
	GetMe(ctx context.Context, actorID domain.AccountID) (*Me, error)
}

/* ---------- DTOs ---------- */

type RegisterLocal struct {
	Login      string
	Password   string
	FirstName  string
	MiddleName string
	LastName   string
	Group      string
	Roles      []domain.GlobalRole
}

type Token struct {
	AccessToken string
}

type Me struct {
	Account    domain.Account
	Roles      []domain.GlobalRole
	Identities []IdentityInfo
}

type IdentityInfo struct {
	Provider domain.IdentityProvider
	Subject  string
	Email    string
}
