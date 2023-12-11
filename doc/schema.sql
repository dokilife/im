-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2023-12-11T07:11:27.630Z

CREATE TABLE "messages" (
  "id" bigserial PRIMARY KEY,
  "from" varchar NOT NULL,
  "message" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
