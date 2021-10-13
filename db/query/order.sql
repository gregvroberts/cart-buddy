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
                        order_paid = $4,
                        order_city = $5,
                        order_state = $6,
                        order_postal = $7,
                        order_country = $8,
                        order_addr_1 = $9,
                        order_addr_2 = $10,
                        order_phone = $11,
                        order_shipping = $12,
                        order_date = $13,
                        order_shipped = $14,
                        order_track_code = $15
WHERE order_id = $1
RETURNING *;

-- name: DeleteOrder :exec
DELETE FROM orders
WHERE order_id = $1;


