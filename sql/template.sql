CREATE TABLE IF NOT EXISTS "cashier_request" (
    "id" UUID NOT NULL PRIMARY KEY,
    "cashier_request_number" VARCHAR NOT NULL,
    "term_payment" term_payment_type NOT NULL DEFAULT 'one-time',
    "term_amount" VARCHAR(255) NOT NULL,
    "currency" VARCHAR(255) NOT NULL DEFAULT 'UZS',
    "description" TEXT,
    "file" VARCHAR(255),
    "supplier_id" UUID NOT NULL REFERENCES "supplier"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE IF NOT EXISTS "cashier_request_product" (
    "id" UUID NOT NULL PRIMARY KEY,
    "barcode" VARCHAR,
    "product_number" VARCHAR NOT NULL,
    "quantity" BIGINT NOT NULL DEFAULT 0,
    "quantity_type" VARCHAR,
    "size_type" VARCHAR,
    "size_value" DECIMAL,
    "weight_type" VARCHAR,
    "weight_value" DECIMAL,
    "price" DECIMAL,
    "metrics" VARCHAR(255) NOT NULL DEFAULT 'ШТ',
    "status" status_enum DEFAULT 'not_viewed',
    "product_id" UUID NOT NULL REFERENCES "product"("id"),
    "cashier_request" UUID NOT NULL REFERENCES "cashier_request"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);
