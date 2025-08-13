-- name: CreateProcess :exec
INSERT INTO processes (id, project_id, template_key, name, current_stage, state)
VALUES ($1, $2, $3, $4, $5, $6);

-- name: GetProcess :one
SELECT id, project_id, template_key, name, current_stage, state, created_at, updated_at
FROM processes
WHERE id = $1;

-- name: SetProcessCurrentStage :exec
UPDATE processes
SET current_stage = $2, updated_at = now()
WHERE id = $1;

-- name: SetProcessState :exec
UPDATE processes
SET state = $2, updated_at = now()
WHERE id = $1;

-- name: GetParentProjectID :one
SELECT project_id
FROM processes
WHERE id = $1;

-- name: AddProcessMember :exec
INSERT INTO process_members (process_id, account_id, role_key)
VALUES ($1, $2, $3)
ON CONFLICT (process_id, account_id, role_key) DO NOTHING;

-- name: RemoveProcessMember :exec
DELETE FROM process_members
WHERE process_id = $1 AND account_id = $2 AND role_key = $3;

-- name: IsProcessMember :one
SELECT EXISTS (
  SELECT 1 FROM process_members
  WHERE process_id = $1 AND account_id = $2 AND role_key = $3
) AS ok;

-- name: ListProcessMembers :many
SELECT process_id, account_id, role_key
FROM process_members
WHERE process_id = $1
ORDER BY account_id, role_key;
