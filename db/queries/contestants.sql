-- name: RegisterContestant :one
INSERT INTO contestants(full_name,email,position, description, region,ethereum_address,national_id_number, image_address)
VALUES ($1,$2,$3,$4,$5, $6,$7,$8) RETURNING *;

-- name: GetAllCandidates :many
SELECT * FROM contestants;