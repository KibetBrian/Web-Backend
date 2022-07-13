-- name: RegisterVoter :one
INSERT INTO voters(full_name, email,national_id_number,image_address,ethereum_address, region)
VALUES($1,$2,$3,$4,$5, $6) RETURNING id;

-- name: UpdateVoter :one
UPDATE voters SET email = $1 WHERE email=$2 RETURNING *;

-- name: TotalVotersNum :one
SELECT COUNT(email) FROM voters WHERE verified= true;

-- name: TotalVotedVoters :one 
SELECT COUNT(email) FROM voters WHERE voted = true;

-- name: DeleteVoter :one
DELETE FROM voters WHERE ethereum_address = $1 RETURNING *;

-- name: PendingVoters :many
SELECT * FROM voters WHERE verified = false;

-- name: UpdatePendingState :one
UPDATE voters SET verified = true WHERE ethereum_address = $1 RETURNING *;

-- name: GetAddress :one
SELECT * FROM voters WHERE ethereum_address = $1;

-- name: VerifiedVoters :many
SELECT * FROM voters WHERE verified = true;

