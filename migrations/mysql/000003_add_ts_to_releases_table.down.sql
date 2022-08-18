BEGIN;

ALTER TABLE releases DROP COLUMN created_at;
ALTER TABLE releases DROP COLUMN updated_at;

COMMIT;