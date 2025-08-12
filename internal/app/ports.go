package app

import (
	"context"
	"time"

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

/* ---------- Accounts & Identities ---------- */

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

/* ---------- Projects & Processes ---------- */

type ProjectRepo interface {
	Create(ctx context.Context, p *domain.Project) error
	Get(ctx context.Context, id domain.ProjectID) (*domain.Project, error)
	AddMember(ctx context.Context, m domain.ProjectMember) error
	RemoveMember(ctx context.Context, m domain.ProjectMember) error
	IsMember(ctx context.Context, projectID domain.ProjectID, accountID domain.AccountID, role domain.ProjectRole) (bool, error)
}

type ProcessRepo interface {
	Create(ctx context.Context, pr *domain.Process) error
	Get(ctx context.Context, id domain.ProcessID) (*domain.Process, error)
	SetCurrentStage(ctx context.Context, id domain.ProcessID, stage domain.StageKey) error
	SetState(ctx context.Context, id domain.ProcessID, state domain.ProcessState) error
	AddMember(ctx context.Context, m domain.ProcessMember) error
	RemoveMember(ctx context.Context, m domain.ProcessMember) error
	IsMember(ctx context.Context, processID domain.ProcessID, accountID domain.AccountID, role domain.ProcessRole) (bool, error)
	ParentProjectID(ctx context.Context, processID domain.ProcessID) (domain.ProjectID, error)
}

/* ---------- Approvals ---------- */

type ApprovalRepo interface {
	Upsert(ctx context.Context, a domain.Approval) error
	CountByDecisionAndRole(ctx context.Context, processID domain.ProcessID, stage domain.StageKey, role domain.ProcessRole, decision domain.Decision) (int, error)
	// optional for UI/audit:
	ListForStage(ctx context.Context, processID domain.ProcessID, stage domain.StageKey) ([]domain.Approval, error)
}

/* ---------- Templates & Events ---------- */

type TemplateProvider interface {
	Load(ctx context.Context, key domain.TemplateKey) (domain.CompiledTemplate, error)
	List(ctx context.Context) ([]domain.TemplateKey, error) // optional
}

type OutboxRepo interface {
	Append(ctx context.Context, e domain.Event) error
}

/* ---------- Auth & utilities ---------- */

type PasswordHasher interface {
	Hash(password string) (string, error)
	Verify(hash, password string) bool
}

type OIDCVerifier interface {
	VerifyIDToken(ctx context.Context, rawIDToken, expectedNonce string) (OIDCClaims, error)
}

type OIDCClaims struct {
	Subject     string
	Email       string
	FirstName   string
	MiddleName  string
	LastName    string
	GroupNumber string
	ExpiresAt   time.Time
}

type Clock interface{ Now() time.Time }
type IDGen interface{ NewID() string }
