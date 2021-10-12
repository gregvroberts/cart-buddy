-- name: CreateOrder :one
INSERT INTO orders (
                    order_user_id,
                    order_amount,
                    order_city,
                    order_state,
                    order_postal,
                    order_country,
                    order_addr_1,
                    order_addr_2,
                    order_phone,
                    order_shipping,
                    order_date,
                    order_shipped,
                    order_track_code
) VALUES (
             $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
         ) RETURNING *;

-- name: GetOrder :one
SELECT * FROM orders
WHERE order_id = $1 LIMIT 1;

-- name: ListOrders :many
SELECT * FROM orders
ORDER BY order_id
LIMIT $1
    OFFSET $2;

-- name: UpdateOrder :one
UPDATE orders
SET                     order_user_id = $2,
                        order_amount = $3,
                        order_city = $4,
                        order_state = $5,
                        order_postal = $6,
                        order_country = $7,
                        order_addr_1 = $8,
                        order_addr_2 = $9,
                        order_phone = $10,
                        order_shipping = $11,
                        order_date = $12,
                        order_shipped = $13,
                        order_track_code = $14
WHERE order_id = $1
RETURNING *;

-- name: DeleteOrder :exec
DELETE FROM orders
WHERE order_id = $1;
