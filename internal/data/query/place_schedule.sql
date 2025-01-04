-- name: CreateSchedule :one
INSERT INTO place_schedule (
    place_id,
    day_of_week,
    open_time,
    close_time
)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetScheduleByID :one
SELECT * FROM place_schedule
WHERE id = $1;

-- name: GetScheduleByPlaceIDAndDay :one
SELECT * FROM place_schedule
WHERE place_id = $1 AND day_of_week = $2;

-- name: ListScheduleByPlaceID :many
SELECT * FROM place_schedule
WHERE place_id = $1;

-- name: ListScheduleByDay :many
SELECT * FROM place_schedule
WHERE day_of_week = $1;

-- name: UpdateSchedule :one
UPDATE place_schedule
SET
    day_of_week = $2,
    open_time = $3,
    close_time = $4
WHERE id = $1
RETURNING *;

-- name: UpdateScheduleByPlaceId :one
UPDATE place_schedule
SET
    day_of_week = $2,
    open_time = $3,
    close_time = $4
WHERE place_id = $1
RETURNING *;

-- name: DeleteSchedule :exec
DELETE FROM place_schedule
WHERE id = $1;