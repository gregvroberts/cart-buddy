// Code generated by sqlc. DO NOT EDIT.
// source: product.sql

package db

import (
	"context"
	"database/sql"
)

const createProduct = `-- name: CreateProduct :one
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
         ) RETURNING product_id, product_sku, product_name, product_price, product_weight, product_cart_desc, product_short_desc, product_long_desc, product_thumb, product_image, product_category_id, product_update_date, product_stock, product_live, product_unlimited, product_location, created_at, updated_at
`

type CreateProductParams struct {
	ProductSku        string         `json:"product_sku"`
	ProductName       string         `json:"product_name"`
	ProductPrice      float64        `json:"product_price"`
	ProductWeight     float64        `json:"product_weight"`
	ProductCartDesc   string         `json:"product_cart_desc"`
	ProductShortDesc  string         `json:"product_short_desc"`
	ProductLongDesc   string         `json:"product_long_desc"`
	ProductThumb      sql.NullString `json:"product_thumb"`
	ProductImage      sql.NullString `json:"product_image"`
	ProductCategoryID int64          `json:"product_category_id"`
	ProductStock      float64        `json:"product_stock"`
	ProductLive       sql.NullBool   `json:"product_live"`
	ProductUnlimited  sql.NullBool   `json:"product_unlimited"`
	ProductLocation   string         `json:"product_location"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, createProduct,
		arg.ProductSku,
		arg.ProductName,
		arg.ProductPrice,
		arg.ProductWeight,
		arg.ProductCartDesc,
		arg.ProductShortDesc,
		arg.ProductLongDesc,
		arg.ProductThumb,
		arg.ProductImage,
		arg.ProductCategoryID,
		arg.ProductStock,
		arg.ProductLive,
		arg.ProductUnlimited,
		arg.ProductLocation,
	)
	var i Product
	err := row.Scan(
		&i.ProductID,
		&i.ProductSku,
		&i.ProductName,
		&i.ProductPrice,
		&i.ProductWeight,
		&i.ProductCartDesc,
		&i.ProductShortDesc,
		&i.ProductLongDesc,
		&i.ProductThumb,
		&i.ProductImage,
		&i.ProductCategoryID,
		&i.ProductUpdateDate,
		&i.ProductStock,
		&i.ProductLive,
		&i.ProductUnlimited,
		&i.ProductLocation,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteProduct = `-- name: DeleteProduct :exec
DELETE FROM products
WHERE product_id = $1
`

func (q *Queries) DeleteProduct(ctx context.Context, productID int64) error {
	_, err := q.db.ExecContext(ctx, deleteProduct, productID)
	return err
}

const getProduct = `-- name: GetProduct :one
SELECT product_id, product_sku, product_name, product_price, product_weight, product_cart_desc, product_short_desc, product_long_desc, product_thumb, product_image, product_category_id, product_update_date, product_stock, product_live, product_unlimited, product_location, created_at, updated_at FROM products
WHERE product_id = $1 LIMIT 1
`

func (q *Queries) GetProduct(ctx context.Context, productID int64) (Product, error) {
	row := q.db.QueryRowContext(ctx, getProduct, productID)
	var i Product
	err := row.Scan(
		&i.ProductID,
		&i.ProductSku,
		&i.ProductName,
		&i.ProductPrice,
		&i.ProductWeight,
		&i.ProductCartDesc,
		&i.ProductShortDesc,
		&i.ProductLongDesc,
		&i.ProductThumb,
		&i.ProductImage,
		&i.ProductCategoryID,
		&i.ProductUpdateDate,
		&i.ProductStock,
		&i.ProductLive,
		&i.ProductUnlimited,
		&i.ProductLocation,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listProducts = `-- name: ListProducts :many
SELECT product_id, product_sku, product_name, product_price, product_weight, product_cart_desc, product_short_desc, product_long_desc, product_thumb, product_image, product_category_id, product_update_date, product_stock, product_live, product_unlimited, product_location, created_at, updated_at FROM products
ORDER BY product_id
LIMIT $1
    OFFSET $2
`

type ListProductsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListProducts(ctx context.Context, arg ListProductsParams) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, listProducts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ProductID,
			&i.ProductSku,
			&i.ProductName,
			&i.ProductPrice,
			&i.ProductWeight,
			&i.ProductCartDesc,
			&i.ProductShortDesc,
			&i.ProductLongDesc,
			&i.ProductThumb,
			&i.ProductImage,
			&i.ProductCategoryID,
			&i.ProductUpdateDate,
			&i.ProductStock,
			&i.ProductLive,
			&i.ProductUnlimited,
			&i.ProductLocation,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProduct = `-- name: UpdateProduct :one
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
RETURNING product_id, product_sku, product_name, product_price, product_weight, product_cart_desc, product_short_desc, product_long_desc, product_thumb, product_image, product_category_id, product_update_date, product_stock, product_live, product_unlimited, product_location, created_at, updated_at
`

type UpdateProductParams struct {
	ProductID         int64          `json:"product_id"`
	ProductSku        string         `json:"product_sku"`
	ProductName       string         `json:"product_name"`
	ProductPrice      float64        `json:"product_price"`
	ProductWeight     float64        `json:"product_weight"`
	ProductCartDesc   string         `json:"product_cart_desc"`
	ProductShortDesc  string         `json:"product_short_desc"`
	ProductLongDesc   string         `json:"product_long_desc"`
	ProductThumb      sql.NullString `json:"product_thumb"`
	ProductImage      sql.NullString `json:"product_image"`
	ProductCategoryID int64          `json:"product_category_id"`
	ProductStock      float64        `json:"product_stock"`
	ProductLive       sql.NullBool   `json:"product_live"`
	ProductUnlimited  sql.NullBool   `json:"product_unlimited"`
	ProductLocation   string         `json:"product_location"`
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, updateProduct,
		arg.ProductID,
		arg.ProductSku,
		arg.ProductName,
		arg.ProductPrice,
		arg.ProductWeight,
		arg.ProductCartDesc,
		arg.ProductShortDesc,
		arg.ProductLongDesc,
		arg.ProductThumb,
		arg.ProductImage,
		arg.ProductCategoryID,
		arg.ProductStock,
		arg.ProductLive,
		arg.ProductUnlimited,
		arg.ProductLocation,
	)
	var i Product
	err := row.Scan(
		&i.ProductID,
		&i.ProductSku,
		&i.ProductName,
		&i.ProductPrice,
		&i.ProductWeight,
		&i.ProductCartDesc,
		&i.ProductShortDesc,
		&i.ProductLongDesc,
		&i.ProductThumb,
		&i.ProductImage,
		&i.ProductCategoryID,
		&i.ProductUpdateDate,
		&i.ProductStock,
		&i.ProductLive,
		&i.ProductUnlimited,
		&i.ProductLocation,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}