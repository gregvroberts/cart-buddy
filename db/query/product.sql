-- name: CreateProduct :one
INSERT INTO products (
                      product_sku,
                      product_name,
                      product_price,
                      product_weight,
                      product_cart_desc,
                      product_short_desc,
                      product_long_desc,
                      product_thumb,
                      product_image,
                      product_category_id,
                      product_stock,
                      product_live,
                      product_unlimited,
                      product_location
) VALUES (
             $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14
         ) RETURNING *;

-- name: GetProduct :one
SELECT * FROM products
WHERE product_id = $1 LIMIT 1;

-- name: ListProducts :many
SELECT * FROM products
ORDER BY product_id
LIMIT $1
    OFFSET $2;

-- name: UpdateProduct :one
UPDATE products
SET
    product_sku = $2,
    product_name = $3,
    product_price = $4,
    product_weight = $5,
    product_cart_desc = $6,
    product_short_desc = $7,
    product_long_desc = $8,
    product_thumb = $9,
    product_image = $10,
    product_category_id = $11,
    product_stock = $12,
    product_live = $13,
    product_unlimited = $14,
    product_location = $15
WHERE product_id = $1
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE product_id = $1;
