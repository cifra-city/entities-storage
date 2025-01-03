-- name: CreatePlace :one
INSERT INTO places (id, name, type, distributor_id, street_id, house_number, location, total_score, reviews_count)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
    RETURNING *;

-- name: GetPlaceByID :one
SELECT * FROM places
WHERE id = $1;

-- name: UpdatePlace :one
UPDATE places
SET name = $2, type = $3, distributor_id = $4, street_id = $5, house_number = $6, location = $7, updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: UpdatePlaceLocation :one
UPDATE places
SET location = $2, house_number = $3, updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: UpdatePlaceDistributor :one
UPDATE places
SET distributor_id = $2, updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: UpdatePlaceName :one
UPDATE places
SET name = $2, updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: UpdatePlaceType :one
UPDATE places
SET type = $2, updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: UpdatePlaceScore :one
UPDATE places
SET total_score = total_score + $2, reviews_count = reviews_count + 1
WHERE id = $1
RETURNING *;

-- name: DeletePlace :exec
DELETE FROM places
WHERE id = $1;

-- name: ListPlaces :many
SELECT * FROM places
ORDER BY created_at DESC;

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