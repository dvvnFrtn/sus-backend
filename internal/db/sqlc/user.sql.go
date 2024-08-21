// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: user.sql

package sqlc

import (
	"context"
	"database/sql"
)

const addUser = `-- name: AddUser :execresult
INSERT INTO users (id, email, password, oauth_id, name, role)
VALUES (?, ?, ?, ?, ?, ?)
`

type AddUserParams struct {
	ID       string
	Email    string
	Password sql.NullString
	OauthID  sql.NullString
	Name     string
	Role     string
}

func (q *Queries) AddUser(ctx context.Context, arg AddUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, addUser,
		arg.ID,
		arg.Email,
		arg.Password,
		arg.OauthID,
		arg.Name,
		arg.Role,
	)
}