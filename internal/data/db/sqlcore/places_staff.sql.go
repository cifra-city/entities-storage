// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: places_staff.sql

package sqlcore

import (
	"context"

	"github.com/google/uuid"
)

const createPlaceStaff = `-- name: CreatePlaceStaff :one
INSERT INTO places_staff (id, place_id, users_id, role)
VALUES ($1, $2, $3, $4)
    RETURNING id, place_id, users_id, role, created_at
`

type CreatePlaceStaffParams struct {
	ID      uuid.UUID
	PlaceID uuid.UUID
	UsersID uuid.UUID
	Role    string
}

func (q *Queries) CreatePlaceStaff(ctx context.Context, arg CreatePlaceStaffParams) (PlacesStaff, error) {
	row := q.db.QueryRowContext(ctx, createPlaceStaff,
		arg.ID,
		arg.PlaceID,
		arg.UsersID,
		arg.Role,
	)
	var i PlacesStaff
	err := row.Scan(
		&i.ID,
		&i.PlaceID,
		&i.UsersID,
		&i.Role,
		&i.CreatedAt,
	)
	return i, err
}

const deletePlaceStaff = `-- name: DeletePlaceStaff :exec
DELETE FROM places_staff
WHERE id = $1
`

func (q *Queries) DeletePlaceStaff(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deletePlaceStaff, id)
	return err
}

const getPlaceStaffByID = `-- name: GetPlaceStaffByID :one
SELECT id, place_id, users_id, role, created_at FROM places_staff
WHERE id = $1
`

func (q *Queries) GetPlaceStaffByID(ctx context.Context, id uuid.UUID) (PlacesStaff, error) {
	row := q.db.QueryRowContext(ctx, getPlaceStaffByID, id)
	var i PlacesStaff
	err := row.Scan(
		&i.ID,
		&i.PlaceID,
		&i.UsersID,
		&i.Role,
		&i.CreatedAt,
	)
	return i, err
}

const listPlaceStaff = `-- name: ListPlaceStaff :many
SELECT id, place_id, users_id, role, created_at FROM places_staff
WHERE place_id = $1
ORDER BY created_at DESC
`

func (q *Queries) ListPlaceStaff(ctx context.Context, placeID uuid.UUID) ([]PlacesStaff, error) {
	rows, err := q.db.QueryContext(ctx, listPlaceStaff, placeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PlacesStaff
	for rows.Next() {
		var i PlacesStaff
		if err := rows.Scan(
			&i.ID,
			&i.PlaceID,
			&i.UsersID,
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

const updatePlaceStaff = `-- name: UpdatePlaceStaff :one
UPDATE places_staff
SET role = $2
WHERE id = $1
    RETURNING id, place_id, users_id, role, created_at
`

type UpdatePlaceStaffParams struct {
	ID   uuid.UUID
	Role string
}

func (q *Queries) UpdatePlaceStaff(ctx context.Context, arg UpdatePlaceStaffParams) (PlacesStaff, error) {
	row := q.db.QueryRowContext(ctx, updatePlaceStaff, arg.ID, arg.Role)
	var i PlacesStaff
	err := row.Scan(
		&i.ID,
		&i.PlaceID,
		&i.UsersID,
		&i.Role,
		&i.CreatedAt,
	)
	return i, err
}