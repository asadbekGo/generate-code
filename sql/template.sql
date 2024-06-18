CREATE TABLE IF NOT EXISTS "supplier" (
    "id" UUID NOT NULL PRIMARY KEY,
    "first_name" VARCHAR(255) NOT NULL,
    "last_name" VARCHAR(255) NOT NULL,
    "birthday" DATE NOT NULL,
    "balance" DECIMAL NOT NULL,
    "currency" VARCHAR(255) NOT NULL DEFAULT 'UZS',
    "phone_number" VARCHAR(255) NOT NULL,
    "address" VARCHAR(255),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
