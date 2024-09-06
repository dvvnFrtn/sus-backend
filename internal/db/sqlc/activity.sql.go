// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: activity.sql

package sqlc

import (
	"context"
	"database/sql"
)

const createActivity = `-- name: CreateActivity :execresult
INSERT INTO activities (id, organization_id, title, note, start_time, end_time)
VALUES (?, ?, ?, ?, ?, ?)
`

type CreateActivityParams struct {
	ID             string
	OrganizationID string
	Title          sql.NullString
	Note           string
	StartTime      sql.NullTime
	EndTime        sql.NullTime
}

func (q *Queries) CreateActivity(ctx context.Context, arg CreateActivityParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createActivity,
		arg.ID,
		arg.OrganizationID,
		arg.Title,
		arg.Note,
		arg.StartTime,
		arg.EndTime,
	)
}

const deleteActivity = `-- name: DeleteActivity :exec
DELETE FROM activities WHERE id = ?
`

func (q *Queries) DeleteActivity(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteActivity, id)
	return err
}

const getActivitiesByOrganizationID = `-- name: GetActivitiesByOrganizationID :many
SELECT id, organization_id, title, note, start_time, end_time, created_at, updated_at FROM activities WHERE organization_id = ?
`

func (q *Queries) GetActivitiesByOrganizationID(ctx context.Context, organizationID string) ([]Activity, error) {
	rows, err := q.db.QueryContext(ctx, getActivitiesByOrganizationID, organizationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Activity
	for rows.Next() {
		var i Activity
		if err := rows.Scan(
			&i.ID,
			&i.OrganizationID,
			&i.Title,
			&i.Note,
			&i.StartTime,
			&i.EndTime,
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

const getActivityByID = `-- name: GetActivityByID :one
SELECT id, organization_id, title, note, start_time, end_time, created_at, updated_at FROM activities WHERE id = ?
`

func (q *Queries) GetActivityByID(ctx context.Context, id string) (Activity, error) {
	row := q.db.QueryRowContext(ctx, getActivityByID, id)
	var i Activity
	err := row.Scan(
		&i.ID,
		&i.OrganizationID,
		&i.Title,
		&i.Note,
		&i.StartTime,
		&i.EndTime,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}