-- name: CreateProject :exec
INSERT INTO projects (id, name, theme, descr, created_by)
VALUES ($1, $2, $3, $4, $5);

-- name: GetProject :one
SELECT id, name, theme, descr, created_by, created_at, updated_at
FROM projects
WHERE id = $1;

-- name: UpdateProjectMeta :exec
UPDATE projects
SET name = $2, theme = $3, descr = $4, updated_at = now()
WHERE id = $1;

-- name: IsProjectMember :one
SELECT EXISTS (
  SELECT 1 FROM project_members
  WHERE project_id = $1 AND account_id = $2 AND role_key = $3
) AS ok;

-- name: AddProjectMember :exec
INSERT INTO project_members (project_id, account_id, role_key)
VALUES ($1, $2, $3)
ON CONFLICT (project_id, account_id, role_key) DO NOTHING;

-- name: RemoveProjectMember :exec
DELETE FROM project_members
WHERE project_id = $1 AND account_id = $2 AND role_key = $3;

-- name: ListProjectsForAccount :many
SELECT p.id, p.name, p.theme, p.descr, p.created_by, p.created_at, p.updated_at
FROM projects p
JOIN project_members m ON m.project_id = p.id
WHERE m.account_id = $1
GROUP BY p.id
ORDER BY p.created_at DESC;

-- name: ListProjectMembers :many
SELECT project_id, account_id, role_key
FROM project_members
WHERE project_id = $1
ORDER BY account_id, role_key;
