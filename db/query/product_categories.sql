-- name: CreateProductCategory :one
INSERT INTO product_categories (
    category_name
) VALUES (
             $1
         ) RETURNING *;

-- name: GetProductCategory :one
SELECT * FROM product_categories
WHERE category_id = $1 LIMIT 1;

-- name: ListProductCategories :many
SELECT * FROM product_categories
ORDER BY category_id
LIMIT $1
    OFFSET $2;

-- name: UpdateProductCategory :one
UPDATE product_categories
SET category_name = $2
WHERE category_id = $1
RETURNING *;

-- name: DeleteProductCategory :exec
DELETE FROM product_categories
WHERE category_id = $1;
