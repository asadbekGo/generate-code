CREATE TABLE IF NOT EXISTS "tender" (
    "id" UUID NOT NULL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "tender_number" VARCHAR NOT NULL,
    "date_time" TIMESTAMP NOT NULL,
    "status" status_tender DEFAULT 'comes',
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);