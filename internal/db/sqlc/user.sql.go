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
INSERT INTO users (
    id, email, password, oauth_id, phone,
    name, role, img, is_premium, lvl,
    dob, institution, created_at, updated_at
)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`

type AddUserParams struct {
	ID          string
	Email       string
	Password    sql.NullString
	OauthID     sql.NullString
	Phone       sql.NullString
	Name        string
	Role        string
	Img         sql.NullString
	IsPremium   sql.NullBool
	Lvl         sql.NullInt32
	Dob         sql.NullTime
	Institution sql.NullString
	CreatedAt   sql.NullTime
	UpdatedAt   sql.NullTime
}

func (q *Queries) AddUser(ctx context.Context, arg AddUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, addUser,
		arg.ID,
		arg.Email,
		arg.Password,
		arg.OauthID,
		arg.Phone,
		arg.Name,
		arg.Role,
		arg.Img,
		arg.IsPremium,
		arg.Lvl,
		arg.Dob,
		arg.Institution,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
}

const emailExists = `-- name: EmailExists :one
SELECT COUNT(1) FROM users WHERE email = ?
`

func (q *Queries) EmailExists(ctx context.Context, email string) (int64, error) {
	row := q.db.QueryRowContext(ctx, emailExists, email)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const findByEmail = `-- name: FindByEmail :one
SELECT id, email, password, oauth_id, name, role, phone, img, is_premium, lvl, dob, institution, created_at, updated_at FROM users WHERE email = ?
`

func (q *Queries) FindByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, findByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.OauthID,
		&i.Name,
		&i.Role,
		&i.Phone,
		&i.Img,
		&i.IsPremium,
		&i.Lvl,
		&i.Dob,
		&i.Institution,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
