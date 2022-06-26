-- name: RegisterVoter :one
INSERT INTO voters(full_name, email, registered_at,voters_public_address)
VALUES($1,$2,$3,$4) RETURNING id;

-- name: UpdateVoter :one
UPDATE voters SET email = $1 WHERE email=$2 RETURNING *;

