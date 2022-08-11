-- name: RegisterUser :one
INSERT INTO users (full_name, email, password)
VALUES($1, $2, $3) RETURNING *;

-- name: UpdateUser :one
UPDATE users SET email = $1, password = $2 WHERE email = $3 RETURNING *;

-- name: CheckEmail :one
SELECT email, count(*) FROM users WHERE email = $1 GROUP BY email;

-- name: GetUser :one
SELECT * FROM users WHERE id=$1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetTotalUsersNum :one
SELECT COUNT(email) FROM users;

-- name: UpdateRegisterationState :one
UPDATE users SET registered_voter = true RETURNING *;

-- name: UpdateVotedPresident :one
UPDATE users SET voted_president = true WHERE email = $1 RETURNING *;

-- name: UpdateVotedGovernor :one
UPDATE users SET voted_governor = true WHERE email = $1 RETURNING *;

