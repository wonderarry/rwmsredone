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

func ParseGlobalRole(s string) (GlobalRole, error) {
	r := GlobalRole(s)
	if slices.Contains(AllGlobalRoles, r) {
		return r, nil
	}
	return GlobalRole(""), ErrUnknownRole
}

func ParseProjectRole(s string) (ProjectRole, error) {
	r := ProjectRole(s)
	if slices.Contains(AllProjectRoles, r) {
		return r, nil
	}
	return ProjectRole(""), ErrUnknownRole
}

func ParseProcessRole(s string) (ProcessRole, error) {
	r := ProcessRole(s)
	if slices.Contains(AllProcessRoles, r) {
		return r, nil
	}
	return ProcessRole(""), ErrUnknownRole
}
