CREATE TABLE "messages" (
  "id" bigserial PRIMARY KEY,
  "from" varchar NOT NULL,
  "message" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
