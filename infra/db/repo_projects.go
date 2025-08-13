package db

import (
	"context"

	"github.com/wonderarry/rwmsredone/infra/db/sqlc"
	"github.com/wonderarry/rwmsredone/internal/app/contract"
	"github.com/wonderarry/rwmsredone/internal/domain"
)

type projectRepo struct{ q *sqlc.Queries }

var _ contract.ProjectRepo = (*projectRepo)(nil)

func (r *projectRepo) Create(ctx context.Context, p *domain.Project) error {
	return r.q.CreateProject(ctx, sqlc.CreateProjectParams{
		ID:        p.ID,
		Name:      p.Name,
		Theme:     textFromPtr(&p.Theme),
		Descr:     textFromPtr(&p.Description),
		CreatedBy: p.CreatedBy,
	})
}

func (r *projectRepo) Get(ctx context.Context, id domain.ProjectID) (*domain.Project, error) {
	row, err := r.q.GetProject(ctx, id)
	if err != nil {
		return nil, err
	}
	return &domain.Project{
		ID:          row.ID,
		Name:        row.Name,
		Theme:       *textPtr(row.Theme),
		Description: *textPtr(row.Descr),
		CreatedBy:   row.CreatedBy,
	}, nil
}

func (r *projectRepo) AddMember(ctx context.Context, m domain.ProjectMember) error {
	return r.q.AddProjectMember(ctx, sqlc.AddProjectMemberParams{
		ProjectID: m.ProjectID,
		AccountID: m.AccountID,
		RoleKey:   string(m.RoleKey),
	})
}

func (r *projectRepo) RemoveMember(ctx context.Context, m domain.ProjectMember) error {
	return r.q.RemoveProjectMember(ctx, sqlc.RemoveProjectMemberParams{
		ProjectID: m.ProjectID,
		AccountID: m.AccountID,
		RoleKey:   string(m.RoleKey),
	})
}

func (r *projectRepo) IsMember(ctx context.Context, projectID domain.ProjectID, accountID domain.AccountID, role domain.ProjectRole) (bool, error) {
	ok, err := r.q.IsProjectMember(ctx, sqlc.IsProjectMemberParams{
		ProjectID: projectID,
		AccountID: accountID,
		RoleKey:   string(role),
	})
	if err != nil {
		return false, err
	}
	return ok, nil
}

func (r *projectRepo) UpdateMeta(ctx context.Context, p *domain.Project) error {
	return r.q.UpdateProjectMeta(ctx, sqlc.UpdateProjectMetaParams{
		ID:    p.ID,
		Name:  p.Name,
		Theme: textFromPtr(&p.Theme),
		Descr: textFromPtr(&p.Description),
	})
}

func (r *projectRepo) ListForAccount(ctx context.Context, id domain.AccountID) ([]*domain.Project, error) {
	rows, err := r.q.ListProjectsForAccount(ctx, id)
	if err != nil {
		return nil, err
	}
	out := make([]*domain.Project, 0, len(rows))
	for _, row := range rows {
		out = append(out, &domain.Project{
			ID:          row.ID,
			Name:        row.Name,
			Theme:       *textPtr(row.Theme),
			Description: *textPtr(row.Descr),
			CreatedBy:   row.CreatedBy,
		})
	}
	return out, nil
}

func (r *projectRepo) ListMembers(ctx context.Context, id domain.ProjectID) ([]domain.ProjectMember, error) {
	rows, err := r.q.ListProjectMembers(ctx, id)
	if err != nil {
		return nil, err
	}
	out := make([]domain.ProjectMember, 0, len(rows))
	for _, rw := range rows {
		out = append(out, domain.ProjectMember{
			ProjectID: rw.ProjectID,
			AccountID: rw.AccountID,
			RoleKey:   domain.ProjectRole(rw.RoleKey),
		})
	}
	return out, nil
}
