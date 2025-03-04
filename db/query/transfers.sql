-- name: CreateTransfers :one
INSERT INTO transfers (
    from_account_id,
    to_account_id,
    amount
) VALUES(
    $1,$2,$3
) 
RETURNING *;

-- name: GetTransfers :one
SELECT * FROM transfers WHERE id = $1 LIMIT 1;

-- name: GetListTransfers :many
SELECT * FROM transfers ORDER by id LIMIT $1 OFFSET $2;

-- name: UpdateTransfers :one
UPDATE transfers 
SET amount = $2 
WHERE id = $1
RETURNING *;

-- name: DeleteTransfers :exec
DELETE FROM transfers 
WHERE id = $1;