-- name: UpsertApproval :exec
INSERT INTO approvals (process_id, stage_key, by_account_id, by_role, decision, comment, created_at)
VALUES ($1, $2, $3, $4, $5, $6, now())
ON CONFLICT (process_id, stage_key, by_account_id)
DO UPDATE SET decision = EXCLUDED.decision,
              comment  = EXCLUDED.comment,
              created_at = now();

-- name: CountApprovalByDecisionAndRole :one
SELECT COUNT(*)
FROM approvals
WHERE process_id = $1
  AND stage_key  = $2
  AND by_role    = $3
  AND decision   = $4;

-- name: ListApprovalsByProcessAndStage :many
SELECT process_id, stage_key, by_account_id, by_role, decision, comment, created_at
FROM approvals
WHERE process_id = $1 AND stage_key = $2
ORDER BY created_at ASC;
