-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2023-12-06T10:16:38.896Z

CREATE TABLE "messages" (
  "id" bigserial PRIMARY KEY,
  "from" varchar NOT NULL,
  "message_content" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
