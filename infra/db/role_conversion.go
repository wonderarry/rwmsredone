package db

import (
	"slices"

	"github.com/wonderarry/rwmsredone/internal/domain"
)

func parseGlobalRole(s string) (domain.GlobalRole, error) {
	r := domain.GlobalRole(s)
	if slices.Contains(domain.AllGlobalRoles, r) {
		return r, nil
	}
	return domain.GlobalRole(""), domain.ErrUnknownRole
}

func parseProjectRole(s string) (domain.ProjectRole, error) {
	r := domain.ProjectRole(s)
	if slices.Contains(domain.AllProjectRoles, r) {
		return r, nil
	}
	return domain.ProjectRole(""), domain.ErrUnknownRole
}

func parseProcessRole(s string) (domain.ProcessRole, error) {
	r := domain.ProcessRole(s)
	if slices.Contains(domain.AllProcessRoles, r) {
		return r, nil
	}
	return domain.ProcessRole(""), domain.ErrUnknownRole
}
