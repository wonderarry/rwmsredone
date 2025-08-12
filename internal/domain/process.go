package domain

type ProcessID = int64
type StageKey = string
type TemplateKey = string

type ProcessState string

const (
	ProcessActive    ProcessState = "active"
	ProcessCompleted ProcessState = "completed"
	ProcessArchived  ProcessState = "archived"
)

type Process struct {
	ID           ProcessID
	ProjectID    ProcessID
	TemplateKey  TemplateKey
	Name         string
	CurrentStage StageKey
	State        ProcessState
}

func (p *Process) Advance(to StageKey) {
	p.CurrentStage = to
}

type ProcessMember struct {
	ProcessID ProcessID
	AccountID AccountID
	RoleKey   ProcessRole
}
