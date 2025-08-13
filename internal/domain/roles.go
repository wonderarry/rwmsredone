package domain

import "slices"

type GlobalRole string
type ProjectRole string
type ProcessRole string

const (
	RoleCanCreateProjects GlobalRole = "CanCreateProjects"
)

var AllGlobalRoles = []GlobalRole{
	RoleCanCreateProjects,
}

const (
	RoleProjectLeader ProjectRole = "ProjectLeader"
	RoleProjectMember ProjectRole = "ProjectMember"
)

var AllProjectRoles = []ProjectRole{
	RoleProjectLeader,
	RoleProjectMember,
}

const (
	RoleAdvisor  ProcessRole = "Advisor"
	RoleStudent  ProcessRole = "Student"
	RoleReviewer ProcessRole = "Reviewer"
)

var AllProcessRoles = []ProcessRole{
	RoleAdvisor,
	RoleStudent,
	RoleReviewer,
}

func hasProjectRole(roles []ProjectRole, want ProjectRole) bool {
	return slices.Contains(roles, want)
}
