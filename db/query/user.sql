-- name: Createuser :one
INSERT INTO users (
    user_f_name,
    user_l_name,
    user_email,
    user_city,
    user_state,
    user_postal,
    user_country,
    user_addr_1,
    user_addr_2
) VALUES (
    $1,$2,$3,$4,$5,$6,$7,$8,$9
         ) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE user_id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY user_id
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users
SET user_f_name = $2,
    user_l_name = $3,
    user_email = $4,
    user_city = $5,
    user_state = $6,
    user_postal = $7,
    user_country = $8,
    user_addr_1 = $9,
    user_addr_2 = $10
WHERE user_id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE user_id = $1;
