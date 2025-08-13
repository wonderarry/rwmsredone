-- name: AppendOutbox :exec
INSERT INTO outbox (topic, payload)
VALUES ($1, $2);
