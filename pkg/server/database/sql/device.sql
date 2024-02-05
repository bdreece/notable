-- name: FindDeviceByID :one
SELECT *
FROM devices
WHERE id = $1;

-- name: FindDeviceByMAC :one
SELECT *
FROM devices
WHERE mac_address = $1;

-- name: FindDevicesByUser :many
SELECT *
FROM devices
WHERE user_id = $1;
