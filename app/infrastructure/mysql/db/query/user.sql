-- name: UserFindByID :one
SELECT
    *
FROM
    users
WHERE
    id = ?;

-- name: UpsertUser :exec
INSERT INTO
    users
    (
        id,
        email,
        first_name,
        last_name,
        phone_number,
        postal_code,
        prefecture,
        city,
        address_extra,
        created_at,
        updated_at
    )
VALUES
    (
        sqlc.arg(id),
        sqlc.arg(email),
        sqlc.arg(first_name),
        sqlc.arg(last_name),
        sqlc.arg(phone_number),
        sqlc.arg(postal_code),
        sqlc.arg(prefecture),
        sqlc.arg(city),
        sqlc.arg(address_extra),
        NOW(),
        NOW()
    )
ON DUPLICATE KEY UPDATE
    email = sqlc.arg(email),
    first_name = sqlc.arg(first_name),
    last_name = sqlc.arg(last_name),
    phone_number = sqlc.arg(phone_number),
    postal_code = sqlc.arg(postal_code),
    prefecture = sqlc.arg(prefecture),
    city = sqlc.arg(city),
    address_extra = sqlc.arg(address_extra),
    updated_at = NOW();

-- name: UserFindAll :many
SELECT
    *
FROM
    users;
