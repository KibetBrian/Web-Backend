-- name: RegisterContestant :one
INSERT INTO contestants(full_name,email,password,position, description)
VALUES ($1,$2,$3,$4, $5) RETURNING *;