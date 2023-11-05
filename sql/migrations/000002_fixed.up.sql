CREATE TABLE "orders" (
  "id" varchar NOT NULL PRIMARY KEY,
  "amount" int NOT NULL,
  "plan" varchar NOT NULL,
  "customer_id" varchar NOT NULL,
  "first_name" varchar NOT NULL,
  "last_name" varchar,
  "email" varchar NOT NULL,
  "status" varchar NOT NULL,
  "transaction_id" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transactions" (
  "id" varchar NOT NULL PRIMARY KEY,
  "amount" int NOT NULL,
  "currency" varchar NOT NULL,
  "payment_intent" varchar NOT NULL,
  "payment_method" varchar NOT NULL,
  "expire_month" int NOT NULL,
  "expire_year" int NOT NULL,
  "transaction_status" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz
);

ALTER TABLE "orders" ADD FOREIGN KEY ("transaction_id") REFERENCES "transactions" ("id");

CREATE INDEX ON "orders" ("email");

