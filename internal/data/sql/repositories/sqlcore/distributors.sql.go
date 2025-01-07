// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: distributors.sql

package sqlcore

import (
	"context"

	"github.com/google/uuid"
)

const createDistributor = `-- name: CreateDistributor :one
INSERT INTO distributors (id, name, owner_id)
VALUES ($1, $2, $3)
    RETURNING id, name, owner_id, created_at, updated_at
`

type CreateDistributorParams struct {
	ID      uuid.UUID
	Name    string
	OwnerID uuid.UUID
}

func (q *Queries) CreateDistributor(ctx context.Context, arg CreateDistributorParams) (Distributor, error) {
	row := q.db.QueryRowContext(ctx, createDistributor, arg.ID, arg.Name, arg.OwnerID)
	var i Distributor
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.OwnerID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getDistributorByID = `-- name: GetDistributorByID :one
SELECT id, name, owner_id, created_at, updated_at FROM distributors
WHERE id = $1
`

func (q *Queries) GetDistributorByID(ctx context.Context, id uuid.UUID) (Distributor, error) {
	row := q.db.QueryRowContext(ctx, getDistributorByID, id)
	var i Distributor
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.OwnerID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listDistributors = `-- name: ListDistributors :many
SELECT id, name, owner_id, created_at, updated_at FROM distributors
ORDER BY created_at DESC
`

func (q *Queries) ListDistributors(ctx context.Context) ([]Distributor, error) {
	rows, err := q.db.QueryContext(ctx, listDistributors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Distributor
	for rows.Next() {
		var i Distributor
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.OwnerID,
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

const updateDistributorName = `-- name: UpdateDistributorName :one
UPDATE distributors
SET name = $2, updated_at = NOW()
WHERE id = $1
    RETURNING id, name, owner_id, created_at, updated_at
`

type UpdateDistributorNameParams struct {
	ID   uuid.UUID
	Name string
}

func (q *Queries) UpdateDistributorName(ctx context.Context, arg UpdateDistributorNameParams) (Distributor, error) {
	row := q.db.QueryRowContext(ctx, updateDistributorName, arg.ID, arg.Name)
	var i Distributor
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.OwnerID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}