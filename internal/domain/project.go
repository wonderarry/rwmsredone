package domain

type ProjectID = int64

type Project struct {
	ID        ProjectID
	Name      string
	CreatedBy AccountID
}

type ProjectMember struct {
	ProjectID ProjectID
	AccountID AccountID
	RoleKey   string // could be ProjectLeader or ProjectMember for now
}

func (p Project) CanManageMembers(actorRoles []string) bool {

	return hasRole(actorRoles, RoleProjectLeader)
}
