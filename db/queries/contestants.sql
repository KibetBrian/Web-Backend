-- name: RegisterCandidate :one
INSERT INTO contestants(full_name,password,position, description)
VALUES ($1,$2,$3,$4) RETURNING *;