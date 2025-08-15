-- name: CreateIdentity :exec
INSERT INTO identities (
  id, account_id, provider, subject, email, password_hash, refresh_token, expires_at
) VALUES ($1,$2,$3,$4,$5,$6,$7,$8);

-- name: GetIdentityByProviderSubject :one
SELECT id, account_id, provider, subject, email, password_hash, refresh_token, expires_at,
       created_at, updated_at
FROM identities
WHERE provider = $1 AND subject = $2;

-- name: ListIdentitiesByAccount :many
SELECT id, account_id, provider, subject, email, password_hash, refresh_token, expires_at,
       created_at, updated_at
FROM identities
WHERE account_id = $1
ORDER BY created_at ASC;
