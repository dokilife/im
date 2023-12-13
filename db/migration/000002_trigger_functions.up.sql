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