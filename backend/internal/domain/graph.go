package domain

type StageDef struct {
	Key         string
	Name        string
	Description string
	PolicyKey   string
	IsTerminal  bool
	Roles       []string
}

type EdgeDef struct {
	FromStageKey string
	ToStageKey   string
	ConditionKey string
}
