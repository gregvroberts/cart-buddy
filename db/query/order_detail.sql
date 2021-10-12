-- name: CreateOrderDetail :one
INSERT INTO order_details (
    detail_order_id,
    detail_product_id,
    detail_product_name,
    detail_unit_price,
    detail_unit_price,
    detail_sku,
    detail_quantity
) VALUES (
             $1, $2, $3, $4, $5, $6, $7
         ) RETURNING *;

-- name: GetOrderDetail :one
SELECT * FROM order_details
WHERE detail_id = $1 LIMIT 1;

-- name: ListOrderDetails :many
SELECT * FROM order_details
ORDER BY detail_id
LIMIT $1
    OFFSET $2;

-- name: UpdateOrderDetail :one
UPDATE order_details
SET detail_order_id = $2,
    detail_product_id = $3,
    detail_product_name = $4,
    detail_unit_price = $5,
    detail_unit_price = $6,
    detail_sku = $7,
    detail_quantity = $8
WHERE detail_id = $1
RETURNING *;

-- name: DeleteOrderDetail :exec
DELETE FROM order_details
WHERE detail_id = $1;
