// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: voter.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const registerVoter = `-- name: RegisterVoter :one
INSERT INTO voters(id, full_name, email,password,registered_at,voters_public_address)
VALUES($1,$2,$3,$4,$5,$6) RETURNING id
`

type RegisterVoterParams struct {
	ID                  uuid.UUID `json:"id"`
	FullName            string    `json:"fullName"`
	Email               string    `json:"email"`
	Password            string    `json:"password"`
	RegisteredAt        time.Time `json:"registeredAt"`
	VotersPublicAddress string    `json:"votersPublicAddress"`
}

func (q *Queries) RegisterVoter(ctx context.Context, arg RegisterVoterParams) (uuid.UUID, error) {
	row := q.queryRow(ctx, q.registerVoterStmt, registerVoter,
		arg.ID,
		arg.FullName,
		arg.Email,
		arg.Password,
		arg.RegisteredAt,
		arg.VotersPublicAddress,
	)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const updateVoter = `-- name: UpdateVoter :one
UPDATE voters SET email = $1, password = $2 WHERE email=$3 RETURNING id, full_name, email, password, registered_at, voted_at, voters_public_address
`

type UpdateVoterParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Email_2  string `json:"email2"`
}

func (q *Queries) UpdateVoter(ctx context.Context, arg UpdateVoterParams) (Voter, error) {
	row := q.queryRow(ctx, q.updateVoterStmt, updateVoter, arg.Email, arg.Password, arg.Email_2)
	var i Voter
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Email,
		&i.Password,
		&i.RegisteredAt,
		&i.VotedAt,
		&i.VotersPublicAddress,
	)
	return i, err
}
