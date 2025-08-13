package domain

// StageDef describes a stage in the process workflow.
type StageDef struct {
	Key         string   // unique key/ID for this stage (within the process)
	Name        string   // human-readable name
	Description string   // optional description
	PolicyKey   string   // key of the policy that governs transitions from this stage
	IsTerminal  bool     // true if this is a final stage (process ends here)
	Roles       []string // allowed roles that can act in this stage (e.g., Advisor, Reviewer)
}

// EdgeDef describes a possible transition between stages.
type EdgeDef struct {
	FromStageKey string // source stage key
	ToStageKey   string // destination stage key
	ConditionKey string // optional key for condition to check before transition
}
