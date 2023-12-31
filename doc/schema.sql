-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2023-12-13T09:22:49.165Z

CREATE TABLE "messages" (
  "id" bigserial PRIMARY KEY,
  "from" varchar NOT NULL,
  "message" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

-- Add a trigger to notify on insert
CREATE OR REPLACE FUNCTION notify_message_insert()
RETURNS TRIGGER AS $$
BEGIN
  PERFORM pg_notify('message_insert', NEW.id::text);
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER message_insert_trigger
AFTER INSERT ON messages
FOR EACH ROW EXECUTE FUNCTION notify_message_insert();