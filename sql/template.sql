CREATE TABLE IF NOT EXISTS "branch" (
    "id" UUID NOT NULL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "phone_number" VARCHAR(255),
    "address" VARCHAR(255),
    "status" status_type DEFAULT "ACTIVE",
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "user" (
    "id" UUID NOT NULL PRIMARY KEY,
    "first_name" VARCHAR(255) NOT NULL,
    "last_name" VARCHAR(255) NOT NULL,
    "login" VARCHAR(255) NOT NULL,
    "phone_number" VARCHAR(255),
    "address" VARCHAR(255),
    "photo" VARCHAR(255),
    "type" role_type NOT NULL,
    "status" status_type DEFAULT "ACTIVE",
    "branch_id" UUID REFERENCES "branch"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "chapter_machine" (
    "id" UUID NOT NULL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "branch_id" UUID NOT NULL REFERENCES "branch"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "machine" (
    "id" UUID NOT NULL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "machine_id" NUMBER NOT NULL,
    "photo" VARCHAR(255),
    "description" TEXT,
    "branch_id" UUID NOT NULL REFERENCES "branch"("id"),
    "chapter_machine_id" UUID NOT NULL REFERENCES "chapter_machine"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "chapter_stock" (
    "id" UUID NOT NULL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "branch_id" UUID NOT NULL REFERENCES "branch"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "product" (
    "id" UUID NOT NULL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "product_id" NUMBER NOT NULL,
    "quantity" NUMBER NOT NULL,
    "photo" VARCHAR(255),
    "description" TEXT,
    "size_type" VARCHAR,
    "size_value" DECIMAL,
    "weight_type" VARCHAR,
    "weight_value" DECIMAL,
    "branch_id" UUID NOT NULL REFERENCES "branch"("id"),
    "chapter_stock_id" UUID NOT NULL REFERENCES "chapter_stock"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "write_off" (
    "id" UUID NOT NULL PRIMARY KEY,
    "write_off_id" NUMBER NOT NULL,
    "quantity" NUMBER NOT NULL,
    "size_type" VARCHAR NOT NULL,
    "size_value" DECIMAL NOT NULL DEFAULT 0,
    "weight_type" VARCHAR NOT NULL,
    "weight_value" DECIMAL NOT NULL DEFAULT 0,
    "date_time" TIMESTAMP NOT NULL,
    "product_id" UUID NOT NULL REFERENCES "product"("id"),
    "machine_id" UUID NOT NULL REFERENCES "machine"("id"),
    "branch_id" UUID NOT NULL REFERENCES "branch"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "request" (
    "id" UUID NOT NULL PRIMARY KEY,
    "date_time" TIMESTAMP NOT NULL,
    "description" TEXT,
    "comments" TEXT,
    "user_id" UUID NOT NULL REFERENCES "user"("id"),
    "user_id_2" UUID NOT NULL REFERENCES "user"("id"),
    "branch_id" UUID NOT NULL REFERENCES "branch"("id"),
    "status" status_enum DEFAULT 'not_viewed',
    "chapter_status" request_chapter DEFAULT 'all',
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "request_product" (
    "id" UUID NOT NULL PRIMARY KEY,
    "quantity" NUMBER NOT NULL,
    "size_type" VARCHAR,
    "size_value" DECIMAL,
    "weight_type" VARCHAR,
    "weight_value" DECIMAL,
    "status" status_enum DEFAULT 'review',
    "branch_id" UUID NOT NULL REFERENCES "branch"("id"),
    "product_id" UUID NOT NULL REFERENCES "product"("id"),
    "request_id" UUID NOT NULL REFERENCES "request"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "transfer_send" (
    "id" UUID NOT NULL PRIMARY KEY,
    "date_time" TIMESTAMP NOT NULL,
    "description" TEXT,
    "comments" TEXT,
    "user_id" UUID NOT NULL REFERENCES "user"("id"),
    "user_id_2" UUID NOT NULL REFERENCES "user"("id"),
    "branch_id" UUID NOT NULL REFERENCES "branch"("id"),
    "branch_id_2" UUID NOT NULL REFERENCES "branch"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "transfer_send_product" (
    "id" UUID NOT NULL PRIMARY KEY,
    "quantity" NUMBER NOT NULL,
    "size_type" VARCHAR,
    "size_value" DECIMAL,
    "weight_type" VARCHAR,
    "weight_value" DECIMAL,
    "status" status_enum DEFAULT 'review',
    "branch_id" UUID NOT NULL REFERENCES "branch"("id"),
    "branch_id_2" UUID NOT NULL REFERENCES "branch"("id"),
    "product_id" UUID NOT NULL REFERENCES "product"("id"),
    "transfer_send_id" UUID NOT NULL REFERENCES "transfer_send"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "transfer_receive" (
    "id" UUID NOT NULL PRIMARY KEY,
    "date_time" TIMESTAMP NOT NULL,
    "description" TEXT,
    "comments" TEXT,
    "user_id" UUID NOT NULL REFERENCES "user"("id"),
    "user_id_2" UUID NOT NULL REFERENCES "user"("id"),
    "branch_id" UUID NOT NULL REFERENCES "branch"("id"),
    "branch_id_2" UUID NOT NULL REFERENCES "branch"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "transfer_receive_product" (
    "id" UUID NOT NULL PRIMARY KEY,
    "quantity" NUMBER NOT NULL,
    "size_type" VARCHAR,
    "size_value" DECIMAL,
    "weight_type" VARCHAR,
    "weight_value" DECIMAL,
    "status" status_enum DEFAULT 'review',
    "branch_id" UUID NOT NULL REFERENCES "branch"("id"),
    "branch_id_2" UUID NOT NULL REFERENCES "branch"("id"),
    "product_id" UUID NOT NULL REFERENCES "product"("id"),
    "transfer_send_id" UUID NOT NULL REFERENCES "transfer_receive"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);
