package domain

type Policy string

const (
	PolicyNoChecks Policy = "NoChecks"

	PolicyCheckConfirmation Policy = "CheckConfirmation"

	PolicyCheckApproves Policy = "CheckApproves"
)

type StageRule struct {
	Key          StageKey
	Policy       Policy
	RequiredRole string
	ApproveEdge  StageKey
	RejectEdge   StageKey
	Threshold    int
	Terminal     bool
}

// This one doesnt get transferred to the db, its an internal definition of a template
type CompiledTemplate struct {
	TemplateKey TemplateKey
	Start       StageKey
	Stages      map[StageKey]StageRule
}

func Evaluate(tpl CompiledTemplate, current StageKey, lastDecision Decision, approvalsCount int) (next StageKey, done bool) {
	/*
		Basic purpose:
		 - This func responds to updates to the process state, and it is supplied with the branch to choose (approval or rejection) if it's possible
		 - It decides, if the stage has to change or stay the same, or maybe if the process was completed (reached terminal stage).

		Decision logic:
		 - Depending on the current stage's policy of transition, it either transitions to the next state (policy satisfied) or says nothing's changed
		 - If the transition is possible, it defaults to approval but if the lastDecision is rejections the verdict's overturned.
	*/

	rule, ok := tpl.Stages[current]

	if !ok {
		return current, false
	}

	switch rule.Policy {

	case PolicyNoChecks:
		next = rule.ApproveEdge

	case PolicyCheckConfirmation:
		if approvalsCount >= 1 {
			next = rule.ApproveEdge
		} else {
			return current, false
		}

	case PolicyCheckApproves:
		thr := rule.Threshold
		if thr <= 0 {
			thr = 1
		}
		if approvalsCount >= thr {
			next = rule.ApproveEdge
		} else {
			return current, false
		}
	default:
		return current, false
	}

	if lastDecision == Reject && rule.RejectEdge != "" {
		next = rule.RejectEdge
	}

	if next == "" || next == current {
		return current, false
	}

	if ns, ok := tpl.Stages[next]; ok && ns.Terminal {
		return next, true
	}
	return next, false
}
