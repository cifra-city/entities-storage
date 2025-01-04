-- name: CreatePlace :one
INSERT INTO places (
    name,
    type,
    description,
    street_id,
    house_number,
    location,
    schedule,
    total_score,
    distributor_id
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
    RETURNING *;

-- name: GetPlaceByID :one
SELECT * FROM places
WHERE id = $1;

-- name: UpdatePlace :one
UPDATE places
SET
    name = $2,
    type = $3,
    description = $4,
    street_id = $5,
    house_number = $6,
    location = $7,
    schedule = $8,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: UpdatePlaceLocation :one
UPDATE places
SET
    street_id = $1,
    house_number = $2,
    location = $3,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: UpdatePlaceHeadline :one
UPDATE places
SET
    name = $2,
    description = $3,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: UpdatePlaceType :one
UPDATE places
SET
    type = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: UpdateGrade :one
UPDATE places
SET
    total_score = total_score + $2,
    reviews_count = reviews_count + 1
WHERE id = $1
RETURNING *;

-- name: DeletePlace :exec
DELETE FROM places
WHERE id = $1;

-- name: ListPlacesByDistributor :many
SELECT * FROM places
WHERE distributor_id = $1;

-- name: ListPlacesByType :many
SELECT * FROM places
WHERE type = $1;

-- name: ListPlacesByStreet :many
SELECT * FROM places
WHERE street_id = $1;

-- name: ListPlacesByStreetAndType :many
SELECT * FROM places
WHERE street_id = $1 AND type = $2;