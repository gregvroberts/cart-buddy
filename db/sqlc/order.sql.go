// Code generated by sqlc. DO NOT EDIT.
// source: order.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createOrder = `-- name: CreateOrder :one
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
         ) RETURNING order_id, order_user_id, order_amount, order_paid, order_city, order_state, order_postal, order_country, order_addr_1, order_addr_2, order_phone, order_shipping, order_date, order_shipped, order_track_code, order_expire, created_at, updated_at
`

type CreateOrderParams struct {
	OrderUserID    int64           `json:"order_user_id"`
	OrderAmount    float64         `json:"order_amount"`
	OrderCity      sql.NullString  `json:"order_city"`
	OrderState     sql.NullString  `json:"order_state"`
	OrderPostal    sql.NullString  `json:"order_postal"`
	OrderCountry   sql.NullString  `json:"order_country"`
	OrderAddr1     sql.NullString  `json:"order_addr_1"`
	OrderAddr2     sql.NullString  `json:"order_addr_2"`
	OrderPhone     sql.NullString  `json:"order_phone"`
	OrderShipping  sql.NullFloat64 `json:"order_shipping"`
	OrderDate      time.Time       `json:"order_date"`
	OrderShipped   bool            `json:"order_shipped"`
	OrderTrackCode sql.NullString  `json:"order_track_code"`
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error) {
	row := q.db.QueryRowContext(ctx, createOrder,
		arg.OrderUserID,
		arg.OrderAmount,
		arg.OrderCity,
		arg.OrderState,
		arg.OrderPostal,
		arg.OrderCountry,
		arg.OrderAddr1,
		arg.OrderAddr2,
		arg.OrderPhone,
		arg.OrderShipping,
		arg.OrderDate,
		arg.OrderShipped,
		arg.OrderTrackCode,
	)
	var i Order
	err := row.Scan(
		&i.OrderID,
		&i.OrderUserID,
		&i.OrderAmount,
		&i.OrderPaid,
		&i.OrderCity,
		&i.OrderState,
		&i.OrderPostal,
		&i.OrderCountry,
		&i.OrderAddr1,
		&i.OrderAddr2,
		&i.OrderPhone,
		&i.OrderShipping,
		&i.OrderDate,
		&i.OrderShipped,
		&i.OrderTrackCode,
		&i.OrderExpire,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteOrder = `-- name: DeleteOrder :exec
DELETE FROM orders
WHERE order_id = $1
`

func (q *Queries) DeleteOrder(ctx context.Context, orderID int64) error {
	_, err := q.db.ExecContext(ctx, deleteOrder, orderID)
	return err
}

const getOrder = `-- name: GetOrder :one
SELECT order_id, order_user_id, order_amount, order_paid, order_city, order_state, order_postal, order_country, order_addr_1, order_addr_2, order_phone, order_shipping, order_date, order_shipped, order_track_code, order_expire, created_at, updated_at FROM orders
WHERE order_id = $1 LIMIT 1
`

func (q *Queries) GetOrder(ctx context.Context, orderID int64) (Order, error) {
	row := q.db.QueryRowContext(ctx, getOrder, orderID)
	var i Order
	err := row.Scan(
		&i.OrderID,
		&i.OrderUserID,
		&i.OrderAmount,
		&i.OrderPaid,
		&i.OrderCity,
		&i.OrderState,
		&i.OrderPostal,
		&i.OrderCountry,
		&i.OrderAddr1,
		&i.OrderAddr2,
		&i.OrderPhone,
		&i.OrderShipping,
		&i.OrderDate,
		&i.OrderShipped,
		&i.OrderTrackCode,
		&i.OrderExpire,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listOrders = `-- name: ListOrders :many
SELECT order_id, order_user_id, order_amount, order_paid, order_city, order_state, order_postal, order_country, order_addr_1, order_addr_2, order_phone, order_shipping, order_date, order_shipped, order_track_code, order_expire, created_at, updated_at FROM orders
ORDER BY order_id
LIMIT $1
    OFFSET $2
`

type ListOrdersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListOrders(ctx context.Context, arg ListOrdersParams) ([]Order, error) {
	rows, err := q.db.QueryContext(ctx, listOrders, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Order
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.OrderID,
			&i.OrderUserID,
			&i.OrderAmount,
			&i.OrderPaid,
			&i.OrderCity,
			&i.OrderState,
			&i.OrderPostal,
			&i.OrderCountry,
			&i.OrderAddr1,
			&i.OrderAddr2,
			&i.OrderPhone,
			&i.OrderShipping,
			&i.OrderDate,
			&i.OrderShipped,
			&i.OrderTrackCode,
			&i.OrderExpire,
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

const updateOrder = `-- name: UpdateOrder :one
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
RETURNING order_id, order_user_id, order_amount, order_paid, order_city, order_state, order_postal, order_country, order_addr_1, order_addr_2, order_phone, order_shipping, order_date, order_shipped, order_track_code, order_expire, created_at, updated_at
`

type UpdateOrderParams struct {
	OrderID        int64           `json:"order_id"`
	OrderUserID    int64           `json:"order_user_id"`
	OrderAmount    float64         `json:"order_amount"`
	OrderPaid      bool            `json:"order_paid"`
	OrderCity      sql.NullString  `json:"order_city"`
	OrderState     sql.NullString  `json:"order_state"`
	OrderPostal    sql.NullString  `json:"order_postal"`
	OrderCountry   sql.NullString  `json:"order_country"`
	OrderAddr1     sql.NullString  `json:"order_addr_1"`
	OrderAddr2     sql.NullString  `json:"order_addr_2"`
	OrderPhone     sql.NullString  `json:"order_phone"`
	OrderShipping  sql.NullFloat64 `json:"order_shipping"`
	OrderDate      time.Time       `json:"order_date"`
	OrderShipped   bool            `json:"order_shipped"`
	OrderTrackCode sql.NullString  `json:"order_track_code"`
}

func (q *Queries) UpdateOrder(ctx context.Context, arg UpdateOrderParams) (Order, error) {
	row := q.db.QueryRowContext(ctx, updateOrder,
		arg.OrderID,
		arg.OrderUserID,
		arg.OrderAmount,
		arg.OrderPaid,
		arg.OrderCity,
		arg.OrderState,
		arg.OrderPostal,
		arg.OrderCountry,
		arg.OrderAddr1,
		arg.OrderAddr2,
		arg.OrderPhone,
		arg.OrderShipping,
		arg.OrderDate,
		arg.OrderShipped,
		arg.OrderTrackCode,
	)
	var i Order
	err := row.Scan(
		&i.OrderID,
		&i.OrderUserID,
		&i.OrderAmount,
		&i.OrderPaid,
		&i.OrderCity,
		&i.OrderState,
		&i.OrderPostal,
		&i.OrderCountry,
		&i.OrderAddr1,
		&i.OrderAddr2,
		&i.OrderPhone,
		&i.OrderShipping,
		&i.OrderDate,
		&i.OrderShipped,
		&i.OrderTrackCode,
		&i.OrderExpire,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
