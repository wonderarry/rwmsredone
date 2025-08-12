package domain

import "slices"

type GlobalRole string
type ProjectRole string
type ProcessRole string

const (
	RoleCanCreateProjects GlobalRole = "CanCreateProjects"
)

const (
	RoleProjectLeader ProjectRole = "ProjectLeader"
	RoleProjectMember ProjectRole = "ProjectMember"
)

const (
	RoleAdvisor  ProcessRole = "Advisor"
	RoleStudent  ProcessRole = "Student"
	RoleReviewer ProcessRole = "Reviewer"
)

func hasProjectRole(roles []ProjectRole, want ProjectRole) bool {
	return slices.Contains(roles, want)
}
