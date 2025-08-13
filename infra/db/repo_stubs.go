package db

import (
	"context"

	"github.com/wonderarry/rwmsredone/infra/db/sqlc"
	"github.com/wonderarry/rwmsredone/internal/app/contract"
	"github.com/wonderarry/rwmsredone/internal/domain"
)

/* ---------- Accounts ---------- */

type accountRepo struct{ q *sqlc.Queries }

var _ contract.AccountRepo = (*accountRepo)(nil)

func (r *accountRepo) Get(ctx context.Context, id domain.AccountID) (*domain.Account, error) {
	return nil, domain.ErrNotImplemented
}
func (r *accountRepo) Create(ctx context.Context, a *domain.Account) error {
	return domain.ErrNotImplemented
}
func (r *accountRepo) AddGlobalRole(ctx context.Context, a *domain.Account) error {
	return domain.ErrNotImplemented
}
func (r *accountRepo) HasGlobalRole(ctx context.Context, id domain.AccountID, role domain.GlobalRole) (bool, error) {
	return false, domain.ErrNotImplemented
}
func (r *accountRepo) ListGlobalRoles(ctx context.Context, a *domain.Account) ([]domain.GlobalRole, error) {
	return nil, domain.ErrNotImplemented
}

/* ---------- Identities ---------- */

type identityRepo struct{ q *sqlc.Queries }

var _ contract.IdentityRepo = (*identityRepo)(nil)

func (r *identityRepo) Create(ctx context.Context, i *domain.Identity) error {
	return domain.ErrNotImplemented
}
func (r *identityRepo) GetByProviderSubject(ctx context.Context, p domain.IdentityProvider, sub string) (*domain.Identity, error) {
	return nil, domain.ErrNotImplemented
}
func (r *identityRepo) ListByAccount(ctx context.Context, id domain.AccountID) ([]*domain.Identity, error) {
	return nil, domain.ErrNotImplemented
}

/* ---------- Projects ---------- */

type projectRepo struct{ q *sqlc.Queries }

var _ contract.ProjectRepo = (*projectRepo)(nil)

func (r *projectRepo) Create(ctx context.Context, p *domain.Project) error {
	return domain.ErrNotImplemented
}
func (r *projectRepo) Get(ctx context.Context, id domain.ProjectID) (*domain.Project, error) {
	return nil, domain.ErrNotImplemented
}
func (r *projectRepo) AddMember(ctx context.Context, m domain.ProjectMember) error {
	return domain.ErrNotImplemented
}
func (r *projectRepo) RemoveMember(ctx context.Context, m domain.ProjectMember) error {
	return domain.ErrNotImplemented
}
func (r *projectRepo) IsMember(ctx context.Context, projectID domain.ProjectID, accountID domain.AccountID, role domain.ProjectRole) (bool, error) {
	return false, domain.ErrNotImplemented
}
func (r *projectRepo) UpdateMeta(ctx context.Context, p *domain.Project) error {
	return domain.ErrNotImplemented
}
func (r *projectRepo) ListForAccount(ctx context.Context, id domain.AccountID) ([]*domain.Project, error) {
	return nil, domain.ErrNotImplemented
}
func (r *projectRepo) ListMembers(ctx context.Context, id domain.ProjectID) ([]domain.ProjectMember, error) {
	return nil, domain.ErrNotImplemented
}

/* ---------- Processes ---------- */

type processRepo struct{ q *sqlc.Queries }

var _ contract.ProcessRepo = (*processRepo)(nil)

func (r *processRepo) Create(ctx context.Context, pr *domain.Process) error {
	return domain.ErrNotImplemented
}
func (r *processRepo) Get(ctx context.Context, id domain.ProcessID) (*domain.Process, error) {
	return nil, domain.ErrNotImplemented
}
func (r *processRepo) SetCurrentStage(ctx context.Context, id domain.ProcessID, stage domain.StageKey) error {
	return domain.ErrNotImplemented
}
func (r *processRepo) SetState(ctx context.Context, id domain.ProcessID, state domain.ProcessState) error {
	return domain.ErrNotImplemented
}
func (r *processRepo) AddMember(ctx context.Context, m domain.ProcessMember) error {
	return domain.ErrNotImplemented
}
func (r *processRepo) RemoveMember(ctx context.Context, m domain.ProcessMember) error {
	return domain.ErrNotImplemented
}
func (r *processRepo) IsMember(ctx context.Context, processID domain.ProcessID, accountID domain.AccountID, role domain.ProcessRole) (bool, error) {
	return false, domain.ErrNotImplemented
}
func (r *processRepo) ParentProjectID(ctx context.Context, processID domain.ProcessID) (domain.ProjectID, error) {
	return "", domain.ErrNotImplemented
}
func (r *processRepo) ListMembers(ctx context.Context, id domain.ProcessID) ([]domain.ProcessMember, error) {
	return nil, domain.ErrNotImplemented
}

/* ---------- Approvals ---------- */

type approvalRepo struct{ q *sqlc.Queries }

var _ contract.ApprovalRepo = (*approvalRepo)(nil)

func (r *approvalRepo) Upsert(ctx context.Context, a domain.Approval) error {
	return domain.ErrNotImplemented
}
func (r *approvalRepo) CountByDecisionAndRole(ctx context.Context, processID domain.ProcessID, stage domain.StageKey, role domain.ProcessRole, decision domain.Decision) (int, error) {
	return 0, domain.ErrNotImplemented
}
func (r *approvalRepo) ListForStage(ctx context.Context, processID domain.ProcessID, stage domain.StageKey) ([]domain.Approval, error) {
	return nil, domain.ErrNotImplemented
}

/* ---------- Outbox ---------- */

type outboxRepo struct{ q *sqlc.Queries }

var _ contract.OutboxRepo = (*outboxRepo)(nil)

func (r *outboxRepo) Append(ctx context.Context, e domain.Event) error {
	return domain.ErrNotImplemented
}
