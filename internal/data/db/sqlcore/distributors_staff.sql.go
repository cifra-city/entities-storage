// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: distributors_staff.sql

package sqlcore

import (
	"context"

	"github.com/google/uuid"
)

const createDistributorStaff = `-- name: CreateDistributorStaff :one
INSERT INTO distributors_staff (id, distributors_id, user_id, role)
VALUES ($1, $2, $3, $4)
    RETURNING id, distributors_id, user_id, role, created_at
`

type CreateDistributorStaffParams struct {
	ID             uuid.UUID
	DistributorsID uuid.UUID
	UserID         uuid.UUID
	Role           string
}

func (q *Queries) CreateDistributorStaff(ctx context.Context, arg CreateDistributorStaffParams) (DistributorsStaff, error) {
	row := q.db.QueryRowContext(ctx, createDistributorStaff,
		arg.ID,
		arg.DistributorsID,
		arg.UserID,
		arg.Role,
	)
	var i DistributorsStaff
	err := row.Scan(
		&i.ID,
		&i.DistributorsID,
		&i.UserID,
		&i.Role,
		&i.CreatedAt,
	)
	return i, err
}

const deleteDistributorStaff = `-- name: DeleteDistributorStaff :exec
DELETE FROM distributors_staff
WHERE id = $1
`

func (q *Queries) DeleteDistributorStaff(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteDistributorStaff, id)
	return err
}

const deleteDistributorStaffByDistributorIDAndUserId = `-- name: DeleteDistributorStaffByDistributorIDAndUserId :exec
DELETE FROM distributors_staff
WHERE distributors_id = $1 AND user_id = $2
`

type DeleteDistributorStaffByDistributorIDAndUserIdParams struct {
	DistributorsID uuid.UUID
	UserID         uuid.UUID
}

func (q *Queries) DeleteDistributorStaffByDistributorIDAndUserId(ctx context.Context, arg DeleteDistributorStaffByDistributorIDAndUserIdParams) error {
	_, err := q.db.ExecContext(ctx, deleteDistributorStaffByDistributorIDAndUserId, arg.DistributorsID, arg.UserID)
	return err
}

const getDistributorStaffByDistributorID = `-- name: GetDistributorStaffByDistributorID :many
SELECT id, distributors_id, user_id, role, created_at FROM distributors_staff
WHERE distributors_id = $1
`

func (q *Queries) GetDistributorStaffByDistributorID(ctx context.Context, distributorsID uuid.UUID) ([]DistributorsStaff, error) {
	rows, err := q.db.QueryContext(ctx, getDistributorStaffByDistributorID, distributorsID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []DistributorsStaff
	for rows.Next() {
		var i DistributorsStaff
		if err := rows.Scan(
			&i.ID,
			&i.DistributorsID,
			&i.UserID,
			&i.Role,
			&i.CreatedAt,
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

const getDistributorStaffByDistributorIDAndUserID = `-- name: GetDistributorStaffByDistributorIDAndUserID :one
SELECT id, distributors_id, user_id, role, created_at FROM distributors_staff
WHERE distributors_id = $1 AND user_id = $2
`

type GetDistributorStaffByDistributorIDAndUserIDParams struct {
	DistributorsID uuid.UUID
	UserID         uuid.UUID
}

func (q *Queries) GetDistributorStaffByDistributorIDAndUserID(ctx context.Context, arg GetDistributorStaffByDistributorIDAndUserIDParams) (DistributorsStaff, error) {
	row := q.db.QueryRowContext(ctx, getDistributorStaffByDistributorIDAndUserID, arg.DistributorsID, arg.UserID)
	var i DistributorsStaff
	err := row.Scan(
		&i.ID,
		&i.DistributorsID,
		&i.UserID,
		&i.Role,
		&i.CreatedAt,
	)
	return i, err
}

const getDistributorStaffByID = `-- name: GetDistributorStaffByID :one
SELECT id, distributors_id, user_id, role, created_at FROM distributors_staff
WHERE id = $1
`

func (q *Queries) GetDistributorStaffByID(ctx context.Context, id uuid.UUID) (DistributorsStaff, error) {
	row := q.db.QueryRowContext(ctx, getDistributorStaffByID, id)
	var i DistributorsStaff
	err := row.Scan(
		&i.ID,
		&i.DistributorsID,
		&i.UserID,
		&i.Role,
		&i.CreatedAt,
	)
	return i, err
}

const getDistributorStaffByUserID = `-- name: GetDistributorStaffByUserID :many
SELECT id, distributors_id, user_id, role, created_at FROM distributors_staff
WHERE user_id = $1
`

func (q *Queries) GetDistributorStaffByUserID(ctx context.Context, userID uuid.UUID) ([]DistributorsStaff, error) {
	rows, err := q.db.QueryContext(ctx, getDistributorStaffByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []DistributorsStaff
	for rows.Next() {
		var i DistributorsStaff
		if err := rows.Scan(
			&i.ID,
			&i.DistributorsID,
			&i.UserID,
			&i.Role,
			&i.CreatedAt,
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

const listDistributorStaff = `-- name: ListDistributorStaff :many
SELECT id, distributors_id, user_id, role, created_at FROM distributors_staff
WHERE distributors_id = $1
ORDER BY created_at DESC
`

func (q *Queries) ListDistributorStaff(ctx context.Context, distributorsID uuid.UUID) ([]DistributorsStaff, error) {
	rows, err := q.db.QueryContext(ctx, listDistributorStaff, distributorsID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []DistributorsStaff
	for rows.Next() {
		var i DistributorsStaff
		if err := rows.Scan(
			&i.ID,
			&i.DistributorsID,
			&i.UserID,
			&i.Role,
			&i.CreatedAt,
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

const updateDistributorStaff = `-- name: UpdateDistributorStaff :one
UPDATE distributors_staff
SET role = $2
WHERE id = $1
RETURNING id, distributors_id, user_id, role, created_at
`

type UpdateDistributorStaffParams struct {
	ID   uuid.UUID
	Role string
}

func (q *Queries) UpdateDistributorStaff(ctx context.Context, arg UpdateDistributorStaffParams) (DistributorsStaff, error) {
	row := q.db.QueryRowContext(ctx, updateDistributorStaff, arg.ID, arg.Role)
	var i DistributorsStaff
	err := row.Scan(
		&i.ID,
		&i.DistributorsID,
		&i.UserID,
		&i.Role,
		&i.CreatedAt,
	)
	return i, err
}
