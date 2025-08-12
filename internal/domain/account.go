package domain

import "time"

type AccountID = string
type IdentityID = string

type Account struct {
	ID          AccountID
	FirstName   string
	MiddleName  string
	LastName    string
	GroupNumber string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type AccountRole struct {
	AccountID AccountID
	Role      GlobalRole
}

type IdentityProvider string

const (
	ProviderLocal          IdentityProvider = "local"
	ProviderUnivercityOIDC IdentityProvider = "university-oidc"
)

type Identity struct {
	ID        int64
	AccountID AccountID
	Provider  IdentityProvider
	Subject   string
	Email     string
	CreatedAt time.Time
}

type LocalCredentialHash struct {
	IdentityID   int64
	PasswordHash string
	PasswordSalt string
}
