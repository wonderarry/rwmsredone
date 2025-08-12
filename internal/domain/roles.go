package domain

import "slices"

type GlobalRole string

const (
	RoleCanCreateProjects GlobalRole = "CanCreateProjects"
)

const (
	RoleProjectLeader = "ProjectLeader"
	RoleProjectMember = "ProjectMember"
)

const (
	RoleAdvisor  = "Advisor"
	RoleStudent  = "Student"
	RoleReviewer = "Reviewer"
)

func hasRole(roles []string, want string) bool {
	return slices.Contains(roles, want)
}
