-- name: CreatePlacesType :one
INSERT INTO place_types (name)
VALUES ($1)
RETURNING *;

-- name: GetPlacesTypeByID :one
SELECT * FROM place_types
WHERE id = $1;

-- name: GetPlacesTypeByName :one
SELECT * FROM place_types
WHERE name = $1;

-- name: UpdatePlacesName :one
UPDATE place_types
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeletePlacesType :exec
DELETE FROM place_types
WHERE id = $1;
