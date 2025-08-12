package domain

const (
	RoleCanCreateProjects = "CanCreateProjects"
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
	for _, r := range roles {
		if r == want {
			return true
		}
	}
	return false
}
