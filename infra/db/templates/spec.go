package templates

import (
	"fmt"
	"slices"
	"strings"

	"github.com/wonderarry/rwmsredone/internal/domain"
)

type rawSpec struct {
	Project struct {
		Type string `yaml:"type"`
		Name struct {
			Full  string `yaml:"full"`
			Short string `yaml:"short"`
		} `yaml:"name"`
		Artifacts map[string]struct {
			Type     string `yaml:"type"`
			Name     string `yaml:"name"`
			Template string `yaml:"template"`
		} `yaml:"artifacts"`
	} `yaml:"project"`

	Dates map[string]string `yaml:"dates"` // TODO: impl dates

	Stages struct {
		Events map[string]struct {
			Name string `yaml:"name"`
			Date string `yaml:"date"`
			Type string `yaml:"type"`
		} `yaml:"events"`

		Operations map[string]struct {
			Name    string   `yaml:"name"`
			Date    string   `yaml:"date"`
			Policy  string   `yaml:"policy"` // NoChecks | CheckConfirmations | CheckApproves
			Role    string   `yaml:"role"`   // Teacher | Student | etc...
			Args    []string `yaml:"arguments"`
			Results []string `yaml:"results"`

			// Optional explicit routing (not sure if it'll stay)
			ApproveTo string `yaml:"approve_to"`
			RejectTo  string `yaml:"reject_to"`

			// Optional for CheckApproves
			Threshold *int `yaml:"threshold"`
		} `yaml:"operations"`
	} `yaml:"stages"`

	Gates []struct {
		From []string `yaml:"from"`
		To   []string `yaml:"to"`
		Type string   `yaml:"type"`
	} `yaml:"gates"`
}

func compileSpec(s rawSpec) (domain.CompiledTemplate, error) {
	// We'll derive:
	// - Start stage: gate with from "start"
	// - Stage rules for each operation
	// - Edges from gates + optional explicit approve_to/reject_to

	// 1) Build adjacency from gates
	outgoing := map[string][]string{}
	for _, g := range s.Gates {
		for _, f := range g.From {
			outgoing[f] = append(outgoing[f], g.To...)

		}
	}

	// 2) Infer start: first to-stage from a gate whose from includes "start"
	var startKey string
	if tos, ok := outgoing["start"]; ok && len(tos) > 0 {
		startKey = tos[0]
	}

	// 3) Compile operations into StageRules
	stages := make(map[domain.StageKey]domain.StageRule, len(s.Stages.Operations))
	for opKey, op := range s.Stages.Operations {
		rule := domain.StageRule{
			Key:          domain.StageKey(opKey),
			Policy:       mapPolicy(op.Policy),
			RequiredRole: mapRole(op.Role),
		}
		if rule.Policy == "" {
			rule.Policy = domain.PolicyNoChecks
		}

		// Threshold for CheckApproves
		if rule.Policy == domain.PolicyCheckApproves {
			thr := 1
			if op.Threshold != nil && *op.Threshold > 0 {
				thr = *op.Threshold
			}
			rule.Threshold = thr
		}

		// Prefer explicit approve_to / reject_to if present
		rule.ApproveEdge = domain.StageKey(strings.TrimSpace(op.ApproveTo))
		rule.RejectEdge = domain.StageKey(strings.TrimSpace(op.RejectTo))

		// If not explicit, try infer from gates
		if rule.ApproveEdge == "" || (rule.Policy != domain.PolicyNoChecks && rule.RejectEdge == "") {
			tos := outgoing[opKey]
			if len(tos) == 1 {
				if rule.ApproveEdge == "" {
					rule.ApproveEdge = domain.StageKey(tos[0])
				}
			} else if len(tos) >= 2 {
				var errTo, okTo string
				for _, cand := range tos {
					if strings.Contains(strings.ToLower(cand), "error") {
						errTo = cand
					} else if strings.Contains(strings.ToLower(cand), "done") || strings.Contains(strings.ToLower(cand), "ok") {
						okTo = cand
					}
				}
				if rule.ApproveEdge == "" {
					if okTo != "" {
						rule.ApproveEdge = domain.StageKey(okTo)
					} else {
						rule.ApproveEdge = domain.StageKey(tos[0])
					}
				}
				if rule.RejectEdge == "" {
					if errTo != "" {
						rule.RejectEdge = domain.StageKey(errTo)
					} else if len(tos) > 1 {
						rule.RejectEdge = domain.StageKey(tos[1])
					}
				}
			}
		}

		stages[domain.StageKey(opKey)] = rule
	}

	// 4) Mark terminal stages: any operation with no outgoing edges at all
	for k, r := range stages {
		if r.ApproveEdge == "" && len(outgoing[string(k)]) == 0 {
			r.Terminal = true
			stages[k] = r
		}
	}

	ct := domain.CompiledTemplate{
		Start:  domain.StageKey(startKey),
		Stages: stages,
	}
	if ct.Start == "" {
		keys := make([]string, 0, len(s.Stages.Operations))
		for k := range s.Stages.Operations {
			keys = append(keys, k)
		}
		if len(keys) == 0 {
			return domain.CompiledTemplate{}, fmt.Errorf("template has no operations")
		}
		slices.Sort(keys)
		ct.Start = domain.StageKey(keys[0])
	}
	return ct, nil
}

func mapPolicy(p string) domain.Policy {
	switch strings.TrimSpace(p) {
	case "NoChecks":
		return domain.PolicyNoChecks
	case "CheckConfirmations", "CheckConfirmation":
		return domain.PolicyCheckConfirmation
	case "CheckApproves":
		return domain.PolicyCheckApproves
	default:
		return ""
	}
}

func mapRole(r string) domain.ProcessRole {
	switch strings.TrimSpace(r) {
	case "Advisor":
		return domain.RoleAdvisor
	case "Student":
		return domain.RoleStudent
	case "Reviewer":
		return domain.RoleReviewer
	case "Teacher": // Not settled on roles currently
		return domain.RoleAdvisor
	default:
		return domain.RoleAdvisor
	}
}
