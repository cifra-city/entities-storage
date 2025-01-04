-- name: CreatePlaceStaff :one
INSERT INTO places_staff (id, place_id, user_id, role)
VALUES ($1, $2, $3, $4)
    RETURNING *;

-- name: GetPlaceStaffByID :one
SELECT * FROM places_staff
WHERE id = $1;

-- name: GetPlaceStaffByPlaceIDAndUserID :one
SELECT * FROM places_staff
WHERE place_id = $1 AND user_id = $2;

-- name: UpdatePlaceStaff :one
UPDATE places_staff
SET role = $2
WHERE id = $1
    RETURNING *;

-- name: DeletePlaceStaff :exec
DELETE FROM places_staff
WHERE id = $1;

-- name: ListPlaceStaff :many
SELECT * FROM places_staff
WHERE place_id = $1
ORDER BY created_at DESC;
