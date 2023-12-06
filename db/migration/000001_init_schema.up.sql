CREATE TABLE "messages" (
  "id" bigserial PRIMARY KEY,
  "from" varchar NOT NULL,
  "message_content" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
