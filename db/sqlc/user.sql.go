// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: user.sql

package db

import (
	"context"
)

const insertUser = `-- name: InsertUser :one
INSERT INTO users (full_name, email, password)
VALUES($1, $2, $3) RETURNING id, full_name, email, password
`

type InsertUserParams struct {
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) (User, error) {
	row := q.queryRow(ctx, q.insertUserStmt, insertUser, arg.FullName, arg.Email, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Email,
		&i.Password,
	)
	return i, err
}