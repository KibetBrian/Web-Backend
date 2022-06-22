-- name: RegisterVoter :one
INSERT INTO voters(id, full_name, email,password,registered_at,voters_public_address)
VALUES($1,$2,$3,$4,$5,$6) RETURNING id;

-- name: UpdateVoter :one
UPDATE voters SET email = $1, password = $2 WHERE email=$3 RETURNING *;
