-- name: CreateAccount :exec
INSERT INTO accounts (id, first_name, middle_name, last_name, grp)
VALUES ($1, $2, $3, $4, $5);

-- name: GetAccount :one
SELECT id, first_name, middle_name, last_name, grp, created_at, updated_at
FROM accounts
WHERE id = $1;

-- name: GrantGlobalRole :exec
INSERT INTO account_global_roles (account_id, role_key)
VALUES ($1, $2)
ON CONFLICT (account_id, role_key) DO NOTHING;

-- name: ListGlobalRoles :many
SELECT role_key
FROM account_global_roles
WHERE account_id = $1
ORDER BY role_key;

-- name: HasGlobalRole :one
SELECT EXISTS (
  SELECT 1 FROM account_global_roles
  WHERE account_id = $1 AND role_key = $2
) AS ok;