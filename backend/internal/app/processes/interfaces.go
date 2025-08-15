package processes

import (
	"context"

	"github.com/wonderarry/rwmsredone/internal/domain"
)

type Service interface {
	CreateProcess(ctx context.Context, cmd CreateProcess) (domain.ProcessID, error)
	AddMember(ctx context.Context, cmd AddProcessMember) error
	RemoveMember(ctx context.Context, cmd RemoveProcessMember) error
	RecordApproval(ctx context.Context, cmd RecordApproval) error

	GetProcess(ctx context.Context, id domain.ProcessID) (*ProcessDetail, error)
	GetProcessGraph(ctx context.Context, id domain.ProcessID) (*Graph, error)
	ListApprovals(ctx context.Context, id domain.ProcessID, stage domain.StageKey) ([]domain.Approval, error)
}

type ProcessDetail struct {
	Process domain.Process
	Members []domain.ProcessMember
}

type Graph struct {
	Stages []domain.StageDef // or a DTO; pull from compiled template
	Edges  []domain.EdgeDef
}

type CreateProcess struct {
	ProjectID   domain.ProjectID
	TemplateKey domain.TemplateKey
	Name        string
	ActorID     domain.AccountID
}

type AddProcessMember struct {
	ProcessID domain.ProcessID
	AccountID domain.AccountID
	Role      domain.ProcessRole
	ActorID   domain.AccountID
}

type RemoveProcessMember struct {
	ProcessID domain.ProcessID
	AccountID domain.AccountID
	Role      domain.ProcessRole
	ActorID   domain.AccountID
}

type RecordApproval struct {
	ProcessID domain.ProcessID
	Decision  domain.Decision
	Comment   string
	ActorID   domain.AccountID
	ActorRole domain.ProcessRole
}
