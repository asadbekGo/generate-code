
CREATE TABLE IF NOT EXISTS "client" (
    "id" UUID NOT NULL PRIMARY KEY,
    "first_name" VARCHAR(255) NOT NULL,
    "last_name" VARCHAR(255) NOT NULL,
    "birthday" DATE NOT NULL,
    "balance" DECIMAL NOT NULL,
    "currency" VARCHAR(255) NOT NULL DEFAULT 'UZS',
    "phone_number" VARCHAR(255) NOT NULL,
    "address" VARCHAR(255),
    "status" status_type DEFAULT 'ACTIVE',
    "description" TEXT,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "client_contract" (
    "id" UUID NOT NULL PRIMARY KEY,
    "from_date" DATE NOT NULL,
    "to_date" DATE NOT NULL,
    "total_amount" DECIMAL NOT NULL,
    "file" VARCHAR(255),
    "description" TEXT,
    "client_id" UUID NOT NULL REFERENCES "client"("id"),
    "cashier_request_id" UUID NOT NULL REFERENCES "cashier_request"("id"),
    "status" contract_status DEFAULT 'no_completed',
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);
