CREATE TABLE "products"
(
    "product_id"          bigserial PRIMARY KEY,
    "product_sku"         VARCHAR(50)   NOT NULL,
    "product_name"        VARCHAR(100)  NOT NULL,
    "product_price"       FLOAT         NOT NULL,
    "product_weight"      FLOAT         NOT NULL,
    "product_cart_desc"   VARCHAR(250)  NOT NULL,
    "product_short_desc"  VARCHAR(1000) NOT NULL,
    "product_long_desc"   text          NOT NULL,
    "product_thumb"       VARCHAR(100),
    "product_image"       VARCHAR(100),
    "product_category_id" bigint        NOT NULL,
    "product_update_date" timestamptz   NOT NULL DEFAULT (now()),
    "product_stock"       FLOAT         NOT NULL,
    "product_live"        boolean                DEFAULT false,
    "product_unlimited"   boolean                DEFAULT false,
    "product_location"    VARCHAR(250)  NOT NULL,
    "created_at"          timestamptz   NOT NULL DEFAULT (now()),
    "updated_at"          timestamptz   NOT NULL DEFAULT (now())
);

CREATE TABLE "product_categories"
(
    "category_id"   bigserial PRIMARY KEY,
    "category_name" VARCHAR(50) NOT NULL
);

CREATE TABLE "users"
(
    "user_id"             bigserial PRIMARY KEY,
    "user_f_name"         VARCHAR(50)  NOT NULL,
    "user_l_name"         VARCHAR(50)  NOT NULL,
    "user_email"          VARCHAR(128) NOT NULL,
    "user_email_verified" boolean               DEFAULT false,
    "user_city"           VARCHAR(50)  NOT NULL,
    "user_state"          VARCHAR(50)  NOT NULL,
    "user_postal"         VARCHAR(20)  NOT NULL,
    "user_country"        VARCHAR(50)  NOT NULL,
    "user_addr_1"         VARCHAR(100) NOT NULL,
    "user_addr_2"         VARCHAR(100),
    "created_at"          timestamptz  NOT NULL DEFAULT (now()),
    "updated_at"          timestamptz  NOT NULL DEFAULT (now())
);

CREATE TABLE "orders"
(
    "order_id"         bigserial PRIMARY KEY,
    "order_user_id"    bigint,
    "order_amount"     FLOAT        NOT NULL,
    "order_city"       VARCHAR(50)  NOT NULL,
    "order_state"      VARCHAR(50)  NOT NULL,
    "order_postal"     VARCHAR(20)  NOT NULL,
    "order_country"    VARCHAR(50)  NOT NULL,
    "order_addr_1"     VARCHAR(100) NOT NULL,
    "order_addr_2"     VARCHAR(100),
    "order_phone"      VARCHAR(20)  NOT NULL,
    "order_shipping"   FLOAT        NOT NULL,
    "order_date"       timestamptz           DEFAULT (now()),
    "order_shipped"    boolean               DEFAULT false,
    "order_track_code" VARCHAR(80),
    "created_at"       timestamptz  NOT NULL DEFAULT (now()),
    "updated_at"       timestamptz  NOT NULL DEFAULT (now())
);

CREATE TABLE "order_details"
(
    "detail_id"           bigserial PRIMARY KEY,
    "detail_order_id"     bigint       NOT NULL,
    "detail_product_id"   bigint       NOT NULL,
    "detail_product_name" VARCHAR(255) NOT NULL,
    "detail_unit_price"   FLOAT        NOT NULL,
    "detail_sku"          VARCHAR(50)  NOT NULL,
    "detail_quantity"     BIGINT       NOT NULL DEFAULT 1,
    "created_at"          timestamptz  NOT NULL DEFAULT (now()),
    "updated_at"          timestamptz  NOT NULL DEFAULT (now())
);

ALTER TABLE "products"
    ADD FOREIGN KEY ("product_category_id") REFERENCES "product_categories" ("category_id");

ALTER TABLE "orders"
    ADD FOREIGN KEY ("order_user_id") REFERENCES "users" ("user_id");

ALTER TABLE "order_details"
    ADD FOREIGN KEY ("detail_order_id") REFERENCES "orders" ("order_id");

ALTER TABLE "order_details"
    ADD FOREIGN KEY ("detail_product_id") REFERENCES "products" ("product_id");

COMMENT
ON COLUMN "products"."created_at" IS 'When the row was created';

COMMENT
ON COLUMN "products"."updated_at" IS 'When the row was last updated';

COMMENT
ON COLUMN "users"."created_at" IS 'When the row was created';

COMMENT
ON COLUMN "users"."updated_at" IS 'When the row was last updated';

COMMENT
ON COLUMN "orders"."created_at" IS 'When the row was created';

COMMENT
ON COLUMN "orders"."updated_at" IS 'When the row was last updated';

COMMENT
ON COLUMN "order_details"."created_at" IS 'When the row was created';

COMMENT
ON COLUMN "order_details"."updated_at" IS 'When the row was last updated';


CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
RETURN NEW;
END;
$$ language plpgsql;

CREATE OR REPLACE FUNCTION trigger_set_product_update_timestamp()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.product_update_date = NOW();
    RETURN NEW;
END;
$$ language plpgsql;


CREATE TRIGGER set_timestamp_products
    BEFORE UPDATE ON products
    FOR EACH ROW
    EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TRIGGER set_timestamp_users
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TRIGGER set_timestamp_orders
    BEFORE UPDATE ON orders
    FOR EACH ROW
    EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TRIGGER set_timestamp_order_details
    BEFORE UPDATE ON order_details
    FOR EACH ROW
    EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TRIGGER set_timestamp_product_update_date
    BEFORE UPDATE ON products
    FOR EACH ROW
    EXECUTE PROCEDURE trigger_set_product_update_timestamp();
