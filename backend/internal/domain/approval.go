package domain

type Decision string

const (
	Approve Decision = "approve"
	Reject  Decision = "reject"
)

type Approval struct {
	ProcessID   ProcessID
	StageKey    StageKey
	ByAccountID AccountID
	ByRole      ProcessRole
	Decision    Decision
	Comment     string
}
