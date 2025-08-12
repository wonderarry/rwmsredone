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
	RoleKey   ProjectRole
}

func (p Project) CanManageMembers(actorRoles []ProjectRole) bool {

	return hasProjectRole(actorRoles, RoleProjectLeader)
}
