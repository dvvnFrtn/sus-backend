// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: organization.sql

package sqlc

import (
	"context"
	"database/sql"
)

const addOrganization = `-- name: AddOrganization :execresult
INSERT INTO organizations (
    id, name, description, header_img, profile_img, created_at, updated_at
) VALUES (?, ?, ?, ?, ?, ?, ?)
`

type AddOrganizationParams struct {
	ID          string
	Name        string
	Description string
	HeaderImg   sql.NullString
	ProfileImg  sql.NullString
	CreatedAt   sql.NullTime
	UpdatedAt   sql.NullTime
}

func (q *Queries) AddOrganization(ctx context.Context, arg AddOrganizationParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, addOrganization,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.HeaderImg,
		arg.ProfileImg,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
}

const deleteOrganization = `-- name: DeleteOrganization :exec
DELETE FROM organizations
WHERE id = ?
`

func (q *Queries) DeleteOrganization(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteOrganization, id)
	return err
}

const findOrganizationById = `-- name: FindOrganizationById :one
SELECT id, name, description, header_img, profile_img, created_at, updated_at FROM organizations WHERE id = ?
`

func (q *Queries) FindOrganizationById(ctx context.Context, id string) (Organization, error) {
	row := q.db.QueryRowContext(ctx, findOrganizationById, id)
	var i Organization
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.HeaderImg,
		&i.ProfileImg,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listOrganization = `-- name: ListOrganization :many
SELECT id, name, description, header_img, profile_img, created_at, updated_at FROM organizations
`

func (q *Queries) ListOrganization(ctx context.Context) ([]Organization, error) {
	rows, err := q.db.QueryContext(ctx, listOrganization)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Organization
	for rows.Next() {
		var i Organization
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.HeaderImg,
			&i.ProfileImg,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateOrganization = `-- name: UpdateOrganization :execresult
UPDATE organizations
SET name = ?, description = ?, header_img = ?, profile_img = ?
WHERE id = ?
`

type UpdateOrganizationParams struct {
	Name        string
	Description string
	HeaderImg   sql.NullString
	ProfileImg  sql.NullString
	ID          string
}

func (q *Queries) UpdateOrganization(ctx context.Context, arg UpdateOrganizationParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateOrganization,
		arg.Name,
		arg.Description,
		arg.HeaderImg,
		arg.ProfileImg,
		arg.ID,
	)
}
