package domain

import "time"

type AccountID = string
type IdentityID = string

/* ---------- Account & Roles ---------- */

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

/* ---------- Identity (auth) ---------- */

type IdentityProvider string

const (
	ProviderLocal          IdentityProvider = "local"
	ProviderUniversityOIDC IdentityProvider = "university-oidc"
)

// Identity represents a login method for an Account.
// One Account can have many Identities (local, OIDC, etc).
//   - local:    Subject = login (username)
//   - oidc:     Subject = "sub" claim (or stable unique per issuer)
type Identity struct {
	ID        IdentityID
	AccountID AccountID
	Provider  IdentityProvider
	Subject   string
	Email     *string // not sure whether we need it for each identity type

	PasswordHash *string // nil for non-local identities

	RefreshToken *string
	ExpiresAt    *time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}
