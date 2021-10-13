// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
	"time"
)

type Order struct {
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
	OrderExpire    time.Time       `json:"order_expire"`
	// When the row was created
	CreatedAt time.Time `json:"created_at"`
	// When the row was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

type OrderDetail struct {
	DetailID          int64   `json:"detail_id"`
	DetailOrderID     int64   `json:"detail_order_id"`
	DetailProductID   int64   `json:"detail_product_id"`
	DetailProductName string  `json:"detail_product_name"`
	DetailUnitPrice   float64 `json:"detail_unit_price"`
	DetailSku         string  `json:"detail_sku"`
	DetailQuantity    int64   `json:"detail_quantity"`
	// When the row was created
	CreatedAt time.Time `json:"created_at"`
	// When the row was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

type Product struct {
	ProductID         int64           `json:"product_id"`
	ProductSku        string          `json:"product_sku"`
	ProductName       string          `json:"product_name"`
	ProductPrice      float64         `json:"product_price"`
	ProductWeight     float64         `json:"product_weight"`
	ProductCartDesc   string          `json:"product_cart_desc"`
	ProductShortDesc  string          `json:"product_short_desc"`
	ProductLongDesc   string          `json:"product_long_desc"`
	ProductThumb      string          `json:"product_thumb"`
	ProductImage      string          `json:"product_image"`
	ProductCategoryID int64           `json:"product_category_id"`
	ProductUpdateDate time.Time       `json:"product_update_date"`
	ProductStock      sql.NullFloat64 `json:"product_stock"`
	ProductLive       bool            `json:"product_live"`
	ProductUnlimited  bool            `json:"product_unlimited"`
	ProductLocation   string          `json:"product_location"`
	// When the row was created
	CreatedAt time.Time `json:"created_at"`
	// When the row was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

type ProductCategory struct {
	CategoryID   int64  `json:"category_id"`
	CategoryName string `json:"category_name"`
}

type User struct {
	UserID            int64          `json:"user_id"`
	UserFName         string         `json:"user_f_name"`
	UserLName         string         `json:"user_l_name"`
	UserEmail         string         `json:"user_email"`
	UserEmailVerified sql.NullBool   `json:"user_email_verified"`
	UserCity          string         `json:"user_city"`
	UserState         string         `json:"user_state"`
	UserPostal        string         `json:"user_postal"`
	UserCountry       string         `json:"user_country"`
	UserAddr1         string         `json:"user_addr_1"`
	UserAddr2         sql.NullString `json:"user_addr_2"`
	// When the row was created
	CreatedAt time.Time `json:"created_at"`
	// When the row was last updated
	UpdatedAt time.Time `json:"updated_at"`
}
