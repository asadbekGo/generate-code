
CREATE TABLE IF NOT EXISTS "coming" (
    "id" UUID NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "quantity" BIGINT NOT NULL DEFAULT 0,
    "quantity_type" VARCHAR,
    "size_type" VARCHAR,
    "size_value" DECIMAL,
    "weight_type" VARCHAR,
    "weight_value" DECIMAL,
    "price" DECIMAL NOT NULL,
    "total_price" DECIMAL NOT NULL,
    "currency" VARCHAR(255) NOT NULL DEFAULT 'UZS',
    "date_time" TIMESTAMP NOT NULL,
    "client_id" UUID REFERENCES "client"("id"),
    "client_contract_id" UUID REFERENCES "client_contract"("id"),
    "product_id" UUID REFERENCES "product"("id"),
    "cashier_request_coming_id" UUID REFERENCES "cashier_request_coming"("id"),
    "user_id" UUID NOT NULL REFERENCES "user"("id"),
    "description" TEXT,
    "type" status_transaction NOT NULL DEFAULT 'other',
    "type_price" VARCHAR(255) NOT NULL DEFAULT '',
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);
